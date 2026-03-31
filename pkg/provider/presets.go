package provider

// Preset represents a built-in provider configuration.
type Preset struct {
	Name        string
	DisplayName string
	BaseURL     string
	OpusModel   string
	SonnetModel string
	HaikuModel  string
}

// BuiltinPresets contains all built-in provider presets.
// Users only need to supply their API key.
var BuiltinPresets = map[string]Preset{
	"deepseek": {
		Name:        "deepseek",
		DisplayName: "DeepSeek",
		BaseURL:     "https://api.deepseek.com/anthropic",
		OpusModel:   "deepseek-reasoner",
		SonnetModel: "deepseek-chat",
		HaikuModel:  "deepseek-chat",
	},
	"openrouter": {
		Name:        "openrouter",
		DisplayName: "OpenRouter",
		BaseURL:     "https://openrouter.ai/api/v1",
		OpusModel:   "anthropic/claude-opus-4",
		SonnetModel: "anthropic/claude-sonnet-4",
		HaikuModel:  "anthropic/claude-haiku-3.5",
	},
	"siliconflow": {
		Name:        "siliconflow",
		DisplayName: "SiliconFlow (硅基流动)",
		BaseURL:     "https://api.siliconflow.cn/anthropic",
		OpusModel:   "deepseek-ai/DeepSeek-R1",
		SonnetModel: "deepseek-ai/DeepSeek-V3",
		HaikuModel:  "deepseek-ai/DeepSeek-V3",
	},
	"zhipu": {
		Name:        "zhipu",
		DisplayName: "Zhipu AI (智谱)",
		BaseURL:     "https://open.bigmodel.cn/api/paas/v4",
		OpusModel:   "glm-4-plus",
		SonnetModel: "glm-4-flash",
		HaikuModel:  "glm-4-flash",
	},
	"volcengine": {
		Name:        "volcengine",
		DisplayName: "Volcengine (火山引擎/豆包)",
		BaseURL:     "https://ark.cn-beijing.volces.com/api/v3",
		OpusModel:   "doubao-pro-256k",
		SonnetModel: "doubao-pro-32k",
		HaikuModel:  "doubao-lite-32k",
	},
}

// GetPreset returns a preset by name, if it exists.
func GetPreset(name string) (Preset, bool) {
	p, ok := BuiltinPresets[name]
	return p, ok
}

// ListPresets returns all preset names.
func ListPresets() []string {
	names := make([]string, 0, len(BuiltinPresets))
	for name := range BuiltinPresets {
		names = append(names, name)
	}
	return names
}
