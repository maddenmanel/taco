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
	// ── International ──────────────────────────────────────────────────────────
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
		SonnetModel: "anthropic/claude-sonnet-4-5",
		HaikuModel:  "anthropic/claude-haiku-3-5",
	},
	"groq": {
		Name:        "groq",
		DisplayName: "Groq (fastest inference)",
		BaseURL:     "https://api.groq.com/openai/v1",
		OpusModel:   "llama-3.3-70b-versatile",
		SonnetModel: "llama-3.3-70b-versatile",
		HaikuModel:  "llama-3.1-8b-instant",
	},
	"together": {
		Name:        "together",
		DisplayName: "Together AI",
		BaseURL:     "https://api.together.xyz/v1",
		OpusModel:   "meta-llama/Llama-3.3-70B-Instruct-Turbo",
		SonnetModel: "meta-llama/Llama-3.3-70B-Instruct-Turbo",
		HaikuModel:  "meta-llama/Llama-3.2-11B-Vision-Instruct-Turbo",
	},
	"mistral": {
		Name:        "mistral",
		DisplayName: "Mistral AI",
		BaseURL:     "https://api.mistral.ai/v1",
		OpusModel:   "mistral-large-latest",
		SonnetModel: "mistral-small-latest",
		HaikuModel:  "mistral-small-latest",
	},
	// ── China ──────────────────────────────────────────────────────────────────
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
		DisplayName: "Zhipu AI (智谱 GLM)",
		BaseURL:     "https://open.bigmodel.cn/api/paas/v4",
		OpusModel:   "glm-4-plus",
		SonnetModel: "glm-4-flash",
		HaikuModel:  "glm-4-flash",
	},
	"volcengine": {
		Name:        "volcengine",
		DisplayName: "Volcengine (火山引擎·豆包)",
		BaseURL:     "https://ark.cn-beijing.volces.com/api/v3",
		OpusModel:   "doubao-pro-256k",
		SonnetModel: "doubao-pro-32k",
		HaikuModel:  "doubao-lite-32k",
	},
	"moonshot": {
		Name:        "moonshot",
		DisplayName: "Moonshot AI (月之暗面·Kimi)",
		BaseURL:     "https://api.moonshot.cn/v1",
		OpusModel:   "moonshot-v1-128k",
		SonnetModel: "moonshot-v1-32k",
		HaikuModel:  "moonshot-v1-8k",
	},
	"qwen": {
		Name:        "qwen",
		DisplayName: "Alibaba Qwen (通义千问)",
		BaseURL:     "https://dashscope.aliyuncs.com/compatible-mode/v1",
		OpusModel:   "qwen-max",
		SonnetModel: "qwen-plus",
		HaikuModel:  "qwen-turbo",
	},
	"yi": {
		Name:        "yi",
		DisplayName: "01.AI (零一万物·Yi)",
		BaseURL:     "https://api.lingyiwanwu.com/v1",
		OpusModel:   "yi-large",
		SonnetModel: "yi-medium",
		HaikuModel:  "yi-spark",
	},
	"baichuan": {
		Name:        "baichuan",
		DisplayName: "Baichuan AI (百川智能)",
		BaseURL:     "https://api.baichuan-ai.com/v1",
		OpusModel:   "Baichuan4",
		SonnetModel: "Baichuan3-Turbo",
		HaikuModel:  "Baichuan3-Turbo-128k",
	},
	"minimax": {
		Name:        "minimax",
		DisplayName: "MiniMax (海螺 AI)",
		BaseURL:     "https://api.minimax.chat/v1",
		OpusModel:   "abab6.5s-chat",
		SonnetModel: "abab5.5-chat",
		HaikuModel:  "abab5.5s-chat",
	},
	"stepfun": {
		Name:        "stepfun",
		DisplayName: "Stepfun (阶跃星辰)",
		BaseURL:     "https://api.stepfun.com/v1",
		OpusModel:   "step-2-16k",
		SonnetModel: "step-1-8k",
		HaikuModel:  "step-1-flash",
	},
	"infini": {
		Name:        "infini",
		DisplayName: "Infini-AI (无问芯穹)",
		BaseURL:     "https://cloud.infini-ai.com/maas/v1",
		OpusModel:   "deepseek-r1",
		SonnetModel: "deepseek-v3",
		HaikuModel:  "deepseek-v3",
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
