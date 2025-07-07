#!/bin/sh
TARGET="$HOME/.local/bin"
mkdir -p "$TARGET"
cp gitzy "$TARGET/gitzy"
chmod +x "$TARGET/gitzy"
echo 'export PATH="$PATH:$TARGET"' >> "$HOME/.profile"
echo "Installed gitzy to $TARGET. Restart your terminal or run: source ~/.profile"
