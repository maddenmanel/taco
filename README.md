# TACO 🌮

**Terminal AI Configuration Organizer**

A lightweight CLI tool that lets you seamlessly switch between AI providers (DeepSeek, OpenRouter, SiliconFlow, etc.) while continuing to use the `claude` command as usual.

No GUI. No database. No background processes. Just pure text config — the Unix way.

## Why TACO?

Heavy tools like cc-switch use SQLite databases, GUI frameworks, and background daemons just to change a few config values. TACO takes the opposite approach:

- **Single binary** — download and run, zero dependencies
- **Pure text** — reads/writes JSON config files, fully transparent
- **Instant** — executes in milliseconds, then exits
- **Reversible** — `taco restore` cleanly removes all changes
- **Non-destructive** — only modifies the `env` field in Claude's settings, preserving everything else

## Quick Start

### Install

```bash
go install github.com/maddenmanel/taco@latest
```

Or build from source:

```bash
git clone https://github.com/maddenmanel/taco.git
cd taco
go build -o taco .
sudo mv taco /usr/local/bin/
```

### Usage

**1. Add a provider (one-time setup):**

```bash
taco add deepseek --key="sk-your-deepseek-key"
```

**2. Switch to it:**

```bash
taco use deepseek
```

**3. Use Claude as usual:**

```bash
claude  # Now powered by DeepSeek behind the scenes!
```

**4. Switch back to official Claude:**

```bash
taco restore
```

## Commands

| Command | Description |
|---------|-------------|
| `taco add <name> --key=KEY` | Add/update a provider |
| `taco use <name>` | Switch Claude Code to a provider |
| `taco current` | Show the active provider |
| `taco list` | List configured & available providers |
| `taco restore` | Restore official Anthropic config |
| `taco remove <name>` | Remove a provider |

## Built-in Presets

These providers are preconfigured — just supply your API key:

| Name | Provider | Base URL |
|------|----------|----------|
| `deepseek` | DeepSeek | `https://api.deepseek.com/anthropic` |
| `openrouter` | OpenRouter | `https://openrouter.ai/api/v1` |
| `siliconflow` | SiliconFlow (硅基流动) | `https://api.siliconflow.cn/anthropic` |
| `zhipu` | Zhipu AI (智谱) | `https://open.bigmodel.cn/api/paas/v4` |
| `volcengine` | Volcengine (火山引擎/豆包) | `https://ark.cn-beijing.volces.com/api/v3` |

### Custom Providers

```bash
taco add my-proxy \
  --key="sk-xxx" \
  --url="https://my-proxy.com/v1" \
  --sonnet="gpt-4o" \
  --opus="o1-preview"
```

## How It Works

TACO modifies the `env` field in `~/.claude/settings.json` to redirect Claude Code's API calls:

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "https://api.deepseek.com/anthropic",
    "ANTHROPIC_AUTH_TOKEN": "sk-your-key",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "deepseek-chat",
    "ANTHROPIC_DEFAULT_OPUS_MODEL": "deepseek-reasoner"
  }
}
```

- Only the `env` field is touched — your other Claude settings (theme, shortcuts, permissions) are preserved
- A backup is saved to `~/.claude/.settings.taco-backup.json` before each switch
- Provider configs are stored in `~/.taco/config.json` as plain JSON

## File Locations

| Path | Purpose |
|------|---------|
| `~/.taco/config.json` | Your provider configurations |
| `~/.claude/settings.json` | Claude Code settings (modified by TACO) |
| `~/.claude/.settings.taco-backup.json` | Auto-backup before each switch |

## License

MIT
