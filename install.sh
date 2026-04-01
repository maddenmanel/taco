#!/bin/sh
# TACO Installer — Linux & macOS
# Usage: curl -sSL https://raw.githubusercontent.com/maddenmanel/taco/main/install.sh | sh
set -e

REPO="maddenmanel/taco"
BIN_NAME="taco"
INSTALL_DIR="$HOME/.local/bin"

# Colors
GREEN='\033[0;32m'
BOLD='\033[1m'
RESET='\033[0m'

say() { printf "${BOLD}%s${RESET}\n" "$1"; }
ok()  { printf "${GREEN}✓${RESET} %s\n" "$1"; }

say "🌮 TACO Installer"
echo ""

# Detect platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64|amd64)   ARCH="amd64" ;;
    arm64|aarch64)  ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    linux)  ;;
    darwin) ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

FILENAME="${BIN_NAME}-${OS}-${ARCH}"

# Get latest version
LATEST=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST" ]; then
    echo "Could not determine latest version."
    echo "Check: https://github.com/${REPO}/releases"
    exit 1
fi

ok "Version: $LATEST"
ok "Platform: ${OS}/${ARCH}"

# Ensure install dir exists and is in PATH
mkdir -p "$INSTALL_DIR"

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    SHELL_RC=""
    case "$SHELL" in
        */zsh)  SHELL_RC="$HOME/.zshrc" ;;
        */fish) SHELL_RC="$HOME/.config/fish/config.fish" ;;
        *)      SHELL_RC="$HOME/.bashrc" ;;
    esac
    echo "export PATH=\"\$HOME/.local/bin:\$PATH\"" >> "$SHELL_RC"
    ok "Added ~/.local/bin to PATH in $SHELL_RC"
    export PATH="$HOME/.local/bin:$PATH"
fi

# Download
URL="https://github.com/${REPO}/releases/download/${LATEST}/${FILENAME}"
TMPFILE=$(mktemp)
echo ""
echo "Downloading from GitHub..."
curl -sSL -o "$TMPFILE" "$URL"

if [ ! -s "$TMPFILE" ]; then
    echo "Download failed. URL: $URL"
    rm -f "$TMPFILE"
    exit 1
fi

chmod +x "$TMPFILE"
mv "$TMPFILE" "${INSTALL_DIR}/${BIN_NAME}"

echo ""
ok "Installed to ${INSTALL_DIR}/${BIN_NAME}"
echo ""
printf "${BOLD}Get started:${RESET}\n"
echo "  taco add deepseek --key=\"sk-your-key\""
echo "  taco use deepseek"
echo "  taco --help"
echo ""
echo "To uninstall at any time: taco uninstall"
