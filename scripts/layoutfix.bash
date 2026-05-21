# layoutfix — readline hotkeys for the current command line (no X11).

_layoutfix_cmd() {
  local bin="${LAYOUTFIX_BIN:-layoutfix}"
  command -v "$bin" &>/dev/null || return 0

  local converted
  converted=$("$bin" -print <<<"$READLINE_LINE") || return 0
  READLINE_LINE=${converted%$'\n'}
  READLINE_POINT=${#READLINE_LINE}
}

_layoutfix_bind() {
  bind -x "$1" 2>/dev/null || true
}

if [[ -n ${BASH_VERSION-} ]]; then
  _layoutfix_bind '"\et": _layoutfix_cmd'    # Alt+T / Alt+е
  _layoutfix_bind '"\ee": _layoutfix_cmd'    # Alt+E
  _layoutfix_bind '"\C-]": _layoutfix_cmd'       # Ctrl+] — fallback в Cursor/VS Code
  _layoutfix_bind '"\C-x\C-f": _layoutfix_cmd' # Ctrl+X, затем Ctrl+F — fallback
fi
