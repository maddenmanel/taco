package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [provider-name]",
	Short: "Test whether a provider's API key is valid and reachable",
	Long: `Send a minimal API request to verify that your provider is reachable
and your API key is accepted. If no provider is specified, tests the
currently active one.`,
	Example: `  taco test              # test the active provider
  taco test deepseek     # test a specific provider`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		var providerName string
		if len(args) == 1 {
			providerName = args[0]
		} else {
			providerName = cfg.ActiveProvider
		}

		if providerName == "" {
			return fmt.Errorf("no active provider. Run: taco use <name>")
		}

		p, ok := cfg.GetProvider(providerName)
		if !ok {
			return fmt.Errorf("provider %q not found. Run: taco list", providerName)
		}

		fmt.Printf("Testing %s ...\n", p.DisplayName)
		fmt.Printf("  Endpoint: %s\n", p.BaseURL)
		fmt.Println()

		result, latency, err := testProvider(p)
		if err != nil {
			fmt.Printf("  ✗ Failed: %v\n", err)
			return nil
		}

		fmt.Printf("  ✓ Connected  (%dms)\n", latency.Milliseconds())
		fmt.Printf("  ✓ Auth valid\n")
		if result != "" {
			fmt.Printf("  ✓ Response:  %s\n", result)
		}
		return nil
	},
}

// testProvider sends a minimal 1-token request to verify connectivity and auth.
func testProvider(p config.Provider) (reply string, latency time.Duration, err error) {
	model := p.SonnetModel
	if model == "" {
		model = p.HaikuModel
	}
	if model == "" {
		model = p.OpusModel
	}

	// Build the messages endpoint URL.
	// Handles both base URLs ending in /anthropic (DeepSeek, SiliconFlow)
	// and plain base URLs (OpenRouter, Groq, etc.).
	baseURL := strings.TrimRight(p.BaseURL, "/")
	var endpoint string
	if strings.HasSuffix(baseURL, "/anthropic") {
		endpoint = baseURL + "/v1/messages"
	} else {
		endpoint = baseURL + "/v1/messages"
	}

	payload := map[string]interface{}{
		"model":      model,
		"max_tokens": 8,
		"messages": []map[string]string{
			{"role": "user", "content": "Reply with just: ok"},
		},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return "", 0, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.APIKey)
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	// Some providers (e.g. OpenRouter) require an origin header
	req.Header.Set("HTTP-Referer", "https://github.com/maddenmanel/taco")

	client := &http.Client{Timeout: 20 * time.Second}
	start := time.Now()
	resp, err := client.Do(req)
	latency = time.Since(start)

	if err != nil {
		return "", latency, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

	switch resp.StatusCode {
	case http.StatusOK:
		// Try to extract the reply text for a nice confirmation
		var result struct {
			Content []struct {
				Text string `json:"text"`
			} `json:"content"`
			// OpenAI-compat fallback
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}
		if json.Unmarshal(respBody, &result) == nil {
			if len(result.Content) > 0 {
				return strings.TrimSpace(result.Content[0].Text), latency, nil
			}
			if len(result.Choices) > 0 {
				return strings.TrimSpace(result.Choices[0].Message.Content), latency, nil
			}
		}
		return "", latency, nil

	case http.StatusUnauthorized:
		return "", latency, fmt.Errorf("invalid API key (401 Unauthorized)")

	case http.StatusForbidden:
		return "", latency, fmt.Errorf("access denied (403 Forbidden) — check your key permissions")

	case http.StatusNotFound:
		return "", latency, fmt.Errorf("endpoint not found (404) — base URL may be wrong: %s", endpoint)

	case http.StatusTooManyRequests:
		return "", latency, fmt.Errorf("rate limited (429) — your key works but is being throttled")

	default:
		// Try to surface the error message from the response body
		var errResp struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(respBody, &errResp) == nil && errResp.Error.Message != "" {
			return "", latency, fmt.Errorf("HTTP %d: %s", resp.StatusCode, errResp.Error.Message)
		}
		return "", latency, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
}

func init() {
	rootCmd.AddCommand(testCmd)
}
