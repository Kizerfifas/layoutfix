# layoutfix — readline hotkeys for the current command line (no X11).
# Alt+T / Alt+E (and Alt+е on RU keyboard — same key as T on many terminals).

_layoutfix_cmd() {
  local bin="${LAYOUTFIX_BIN:-layoutfix}"
  command -v "$bin" &>/dev/null || return 0

  local converted
  converted=$("$bin" -print <<<"$READLINE_LINE") || return 0
  READLINE_LINE=${converted%$'\n'}
  READLINE_POINT=${#READLINE_LINE}
}

_layoutfix_bind() {
  bind -x "$1": _layoutfix_cmd 2>/dev/null || true
}

if [[ -n ${BASH_VERSION-} ]]; then
  _layoutfix_bind '"\et"'   # Alt+T / Alt+е (клавиша «е» в ЙЦУКЕН)
  _layoutfix_bind '"\ee"'   # Alt+E
fi
