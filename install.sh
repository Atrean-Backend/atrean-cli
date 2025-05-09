#!/bin/bash

set -e

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Atrean CLI Installer${NC}"
echo "This script will install the Atrean CLI tool on your system."
echo ""

# Detect OS
echo "Detecting operating system..."
OS="$(uname)"
case $OS in
  Darwin)
    PLATFORM="darwin"
    ;;
  Linux)
    PLATFORM="linux"
    ;;
  *)
    echo -e "${RED}Unsupported operating system: $OS${NC}"
    exit 1
    ;;
esac

# Detect architecture
echo "Detecting architecture..."
ARCH="$(uname -m)"
case $ARCH in
  x86_64)
    ARCHITECTURE="amd64"
    ;;
  arm64|aarch64)
    ARCHITECTURE="arm64"
    ;;
  *)
    echo -e "${RED}Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

echo -e "Detected: ${GREEN}$PLATFORM/$ARCHITECTURE${NC}"

# Set install directory
INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
  echo -e "${YELLOW}Notice: $INSTALL_DIR is not writable. Using $HOME/bin instead.${NC}"
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"
fi

# Get the latest release URL
GITHUB_REPO="Atrean-Backend/atrean-cli"
echo "Fetching latest release information..."
DOWNLOAD_FILENAME="atrean-$PLATFORM-$ARCHITECTURE"
if [ "$PLATFORM" = "windows" ]; then
  DOWNLOAD_FILENAME="$DOWNLOAD_FILENAME.exe"
fi

DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/latest/download/$DOWNLOAD_FILENAME"

echo "Downloading Atrean CLI from $DOWNLOAD_URL..."
if command -v curl &> /dev/null; then
  curl -L -o "$INSTALL_DIR/atrean" "$DOWNLOAD_URL"
elif command -v wget &> /dev/null; then
  wget -O "$INSTALL_DIR/atrean" "$DOWNLOAD_URL"
else
  echo -e "${RED}Error: Neither curl nor wget found. Please install one of them and try again.${NC}"
  exit 1
fi

# Make executable
chmod +x "$INSTALL_DIR/atrean"

# Check if installation was successful
if [ -x "$INSTALL_DIR/atrean" ]; then
  echo -e "${GREEN}Atrean CLI installed successfully to $INSTALL_DIR/atrean${NC}"
  
  # Add to PATH if needed
  if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${YELLOW}Notice: $INSTALL_DIR is not in your PATH.${NC}"
    echo "Add it to your shell profile with:"
    echo -e "${GREEN}  export PATH=\$PATH:$INSTALL_DIR${NC}"
  fi
  
  echo ""
  echo "To get started, run:"
  echo -e "${GREEN}  atrean setup${NC}"
  echo -e "Then start your services with: ${GREEN}atrean setup start${NC} or ${GREEN}atrean up${NC}"
else
  echo -e "${RED}Installation failed.${NC}"
  exit 1
fi 