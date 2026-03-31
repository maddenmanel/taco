#!/bin/sh
set -e

# TACO Installer — downloads the latest pre-built binary for your platform.
# Usage: curl -sSL https://raw.githubusercontent.com/maddenmanel/taco/main/install.sh | sh

REPO="maddenmanel/taco"
INSTALL_DIR="/usr/local/bin"
APP_NAME="taco"

get_os() {
  case "$(uname -s)" in
    Linux*)  echo "linux" ;;
    Darwin*) echo "darwin" ;;
    MINGW*|MSYS*|CYGWIN*) echo "windows" ;;
    *) echo "unsupported"; exit 1 ;;
  esac
}

get_arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    *) echo "unsupported"; exit 1 ;;
  esac
}

OS=$(get_os)
ARCH=$(get_arch)

echo "🌮 TACO Installer"
echo "   OS:   ${OS}"
echo "   Arch: ${ARCH}"
echo ""

# Get latest release tag
LATEST=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST" ]; then
  echo "Error: Could not determine latest release."
  echo "Please download manually from: https://github.com/${REPO}/releases"
  exit 1
fi

echo "   Version: ${LATEST}"

# Build download URL
SUFFIX=""
if [ "$OS" = "windows" ]; then
  SUFFIX=".exe"
fi
FILENAME="${APP_NAME}-${OS}-${ARCH}${SUFFIX}"
URL="https://github.com/${REPO}/releases/download/${LATEST}/${FILENAME}"

echo "   Downloading: ${URL}"
echo ""

# Download binary
TMPFILE=$(mktemp)
curl -sSL -o "$TMPFILE" "$URL"

if [ ! -s "$TMPFILE" ]; then
  echo "Error: Download failed."
  rm -f "$TMPFILE"
  exit 1
fi

chmod +x "$TMPFILE"

# Install
if [ -w "$INSTALL_DIR" ]; then
  mv "$TMPFILE" "${INSTALL_DIR}/${APP_NAME}"
else
  echo "   Installing to ${INSTALL_DIR} (requires sudo)..."
  sudo mv "$TMPFILE" "${INSTALL_DIR}/${APP_NAME}"
fi

echo "🌮 TACO installed successfully!"
echo "   Run: taco --help"
