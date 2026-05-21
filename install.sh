#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

go build -ldflags="-s -w" -o layoutfix .

BIN_DIR="${LAYOUTFIX_BIN_DIR:-$HOME/.local/bin}"
CONF_DIR="${LAYOUTFIX_CONF_DIR:-$HOME/.config/layoutfix}"

mkdir -p "$BIN_DIR" "$CONF_DIR"
install -m 755 layoutfix "$BIN_DIR/layoutfix"
install -m 644 scripts/layoutfix.bash "$CONF_DIR/layoutfix.bash"

# PATH for this session
case ":$PATH:" in
  *":$BIN_DIR:"*) ;;
  *) export PATH="$BIN_DIR:$PATH" ;;
esac

SNIPPET="source $CONF_DIR/layoutfix.bash"
BASHRC="${HOME}/.bashrc"
if [[ -f $BASHRC ]] && grep -qF 'layoutfix.bash' "$BASHRC"; then
  echo "bash: уже подключён $CONF_DIR/layoutfix.bash в ~/.bashrc"
else
  {
    echo ''
    echo '# layoutfix: Alt+T — исправить раскладку текущей строки в bash'
    echo "export LAYOUTFIX_BIN=$BIN_DIR/layoutfix"
    echo "$SNIPPET"
  } >>"$BASHRC"
  echo "bash: добавлено в ~/.bashrc"
fi

echo ""
echo "Установлено:"
echo "  бинарник:  $BIN_DIR/layoutfix"
echo "  bash:      $CONF_DIR/layoutfix.bash"
echo ""
echo "Использование в терминале (без X11):"
echo "  1) new bash   или   source ~/.bashrc"
echo "  2) наберите:  ыгвщ ыныеуьсед"
echo "  3) Alt+T / Alt+E  (или Ctrl+] в Cursor)  →  sudo systemctl"
echo ""
echo "Режим выделения (нужны xdotool+xclip или wtype+wl-clipboard):"
echo "  layoutfix -selection"
echo ""
echo "Убедитесь, что $BIN_DIR в PATH (обычно уже есть для ~/.local/bin)."
