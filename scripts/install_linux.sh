#!/bin/sh
TARGET="$HOME/.local/bin"
mkdir -p "$TARGET"
cp gitzy "$TARGET/gitzy"
chmod +x "$TARGET/gitzy"

# Append to shell config
SHELL_RC=""
if [ -n "$ZSH_VERSION" ]; then
  SHELL_RC="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
  SHELL_RC="$HOME/.bashrc"
else
  SHELL_RC="$HOME/.profile"
fi

if ! grep -q "$TARGET" "$SHELL_RC"; then
  echo "export PATH=\"\$PATH:$TARGET\"" >> "$SHELL_RC"
fi

# Source the file to take effect immediately
. "$SHELL_RC"

echo "Installed gitzy to $TARGET. You can now use 'gitzy' from the terminal."
