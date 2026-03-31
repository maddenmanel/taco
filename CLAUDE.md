# CLAUDE.md

This file provides guidance to AI assistants working in this repository.

## Repository Overview

**Project**: TACO 🌮 (Terminal AI Configuration Organizer)  
**Owner**: maddenmanel  
**Language**: Go 1.24+  
**Purpose**: Lightweight CLI tool for seamlessly switching Claude Code between AI providers (DeepSeek, OpenRouter, SiliconFlow, etc.)

---

## Codebase Structure

```
taco/
├── main.go                  # Entry point — calls cmd.Execute()
├── go.mod / go.sum           # Go module dependencies
├── cmd/                      # CLI command definitions (cobra)
│   ├── root.go               # Root command & app description
│   ├── add.go                # `taco add` — register a provider
│   ├── use.go                # `taco use` — switch active provider
│   ├── list.go               # `taco list` — show all providers
│   ├── current.go            # `taco current` — show active provider
│   ├── restore.go            # `taco restore` — revert to official Claude
│   └── remove.go             # `taco remove` — delete a provider
├── pkg/
│   ├── config/
│   │   └── config.go         # TacoConfig model, ~/.taco/config.json R/W
│   ├── claude/
│   │   └── settings.go       # ~/.claude/settings.json injection & restore
│   └── provider/
│       └── presets.go         # Built-in provider presets (DeepSeek, etc.)
├── CLAUDE.md                 # This file
├── README.md                 # User-facing documentation
└── .gitignore
```

---

## Technology Stack

- **Language**: Go 1.24+
- **CLI Framework**: [spf13/cobra](https://github.com/spf13/cobra)
- **Package Manager**: Go modules (`go mod`)
- **Database**: None — all config is plain JSON files
- **External Dependencies**: cobra, pflag (transitive)

---

## Development Workflow

### Build

```bash
go build -o taco .
```

### Run locally

```bash
./taco --help
./taco add deepseek --key="sk-test"
./taco use deepseek
./taco current
./taco restore
```

### Install globally

```bash
go install github.com/maddenmanel/taco@latest
```

### Cross-compile

```bash
GOOS=darwin  GOARCH=arm64 go build -o taco-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o taco.exe .
GOOS=linux   GOARCH=amd64 go build -o taco-linux-amd64 .
```

---

## Architecture & Key Concepts

### How switching works

1. User runs `taco use <provider>`
2. TACO reads `~/.claude/settings.json`
3. Backs up the file to `~/.claude/.settings.taco-backup.json`
4. Injects/updates only the `env` field with provider-specific vars:
   - `ANTHROPIC_BASE_URL`, `ANTHROPIC_AUTH_TOKEN`
   - `ANTHROPIC_DEFAULT_OPUS_MODEL`, `ANTHROPIC_DEFAULT_SONNET_MODEL`
5. User runs `claude` normally — requests route through the new provider

### Key design principles

- **Non-destructive**: Only the `env` field in settings.json is modified; all other user settings are preserved
- **Reversible**: `taco restore` cleanly removes all injected vars
- **Transparent**: All config is plain JSON — users can inspect and hand-edit
- **No daemon**: TACO runs, modifies config, exits immediately

### File locations at runtime

| Path | Purpose |
|------|---------|
| `~/.taco/config.json` | Provider configurations (API keys, URLs, models) |
| `~/.claude/settings.json` | Claude Code settings (TACO modifies `env` field) |
| `~/.claude/.settings.taco-backup.json` | Auto-backup before each switch |

---

## Code Conventions

### Go style
- Follow standard `gofmt` formatting
- Error handling: return `fmt.Errorf("context: %w", err)` for wrapping
- Package naming: short, lowercase, single-word (`config`, `claude`, `provider`)
- Exported functions have doc comments; unexported helpers do not need them

### Package responsibilities
- `cmd/` — CLI wiring only; delegates to `pkg/` for logic
- `pkg/config/` — TACO's own config (provider list, active provider)
- `pkg/claude/` — reading/writing Claude Code's `settings.json`
- `pkg/provider/` — built-in preset definitions (no I/O)

### Adding a new provider preset
Edit `pkg/provider/presets.go` and add an entry to `BuiltinPresets`.

### Adding a new CLI command
1. Create `cmd/<command>.go`
2. Define the cobra command
3. Register it with `rootCmd.AddCommand()` in `init()`

---

## Git Conventions

### Commit Messages
- Imperative mood: "Add feature" not "Added feature"
- Subject line under 72 characters
- Reference issues when applicable: `Fix login bug (#42)`

### Branch Naming
- Features: `feature/<short-description>`
- Bug fixes: `fix/<short-description>`
- Claude-initiated: `claude/<short-description>-<id>`

### Commit Signing
SSH commit signing is enabled. Do not bypass with `--no-gpg-sign`.

---

## AI Assistant Behavior

- Read files before modifying them
- Do not add features beyond what is asked
- Do not create a PR unless explicitly requested
- Do not commit secrets or API keys
- Confirm before destructive actions (deleting files, force-pushing, etc.)
- Keep the binary name as `taco` — do not rename
- When modifying provider presets, verify the base URL format is correct for that provider's Anthropic-compatible endpoint
