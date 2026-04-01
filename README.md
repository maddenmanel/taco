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

## Installation

One command. No admin rights. No package manager required.

### Linux / macOS

```bash
curl -sSL https://raw.githubusercontent.com/maddenmanel/taco/main/install.sh | sh
```

Installs to `~/.local/bin/taco`. That's it.

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/maddenmanel/taco/main/install.ps1 | iex
```

Installs to `%USERPROFILE%\.taco\bin\taco.exe`, adds it to your PATH, and automatically removes the "downloaded from internet" flag — **no SmartScreen warning**.

Open a new terminal window after installation and run `taco --help`.

### With Go (any platform)

```bash
go install github.com/maddenmanel/taco@latest
```

## Uninstall

```bash
taco uninstall
```

That's it — restores Claude to official config, removes all TACO data, and deletes the binary. One command, completely clean.

## Quick Start

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
| `taco use [name]` | Switch provider (interactive picker if no name given) |
| `taco test [name]` | Verify API key and connectivity |
| `taco current` | Show the active provider |
| `taco list` | List configured providers and all available presets |
| `taco restore` | Restore official Anthropic config |
| `taco remove <name>` | Remove a provider |
| `taco uninstall` | Remove TACO completely |
| `taco version` | Print version |

## Built-in Presets

15 providers preconfigured — just supply your API key.

### International

| Name | Provider |
|------|----------|
| `deepseek` | DeepSeek |
| `openrouter` | OpenRouter |
| `groq` | Groq (fastest inference) |
| `together` | Together AI |
| `mistral` | Mistral AI |

### China / 国内

| Name | Provider |
|------|----------|
| `siliconflow` | SiliconFlow (硅基流动) |
| `zhipu` | Zhipu AI (智谱 GLM) |
| `volcengine` | Volcengine (火山引擎·豆包) |
| `moonshot` | Moonshot AI (月之暗面·Kimi) |
| `qwen` | Alibaba Qwen (通义千问) |
| `yi` | 01.AI (零一万物·Yi) |
| `baichuan` | Baichuan AI (百川智能) |
| `minimax` | MiniMax (海螺 AI) |
| `stepfun` | Stepfun (阶跃星辰) |
| `infini` | Infini-AI (无问芯穹) |

### Custom Providers

```bash
taco add my-proxy \
  --key="sk-xxx" \
  --url="https://my-proxy.com/v1" \
  --sonnet="my-model" \
  --opus="my-model-pro"
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

| Linux / macOS | Windows | Purpose |
|---------------|---------|---------|
| `~/.taco/config.json` | `%USERPROFILE%\.taco\config.json` | Your provider configurations |
| `~/.claude/settings.json` | `%USERPROFILE%\.claude\settings.json` | Claude Code settings (modified by TACO) |
| `~/.claude/.settings.taco-backup.json` | `%USERPROFILE%\.claude\.settings.taco-backup.json` | Auto-backup before each switch |

## License

MIT
