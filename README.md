# layoutfix

**EN** · [English](#english) | **RU** · [Русский](#русский)

Fix text typed with the wrong keyboard layout (Russian **ЙЦУКЕН** ↔ English **QWERTY**).

```
ыгвщ ыныеуьсед   →   sudo systemctl
ды               →   ls
```

---

## English

### Two modes

| Mode | Hotkey | Needs X11? | Use case |
|------|--------|------------|----------|
| **Bash (default)** | **Alt+T** or **Alt+E** in the shell | **No** | Current command line in bash — SSH, TTY, VM without graphics |
| **Selection** | System/terminal shortcut → `layoutfix -selection` | **Yes** (or Wayland tools) | Any highlighted text in a graphical terminal or app |

**Recommended for servers and terminals without X:** install + bash integration (readline).  
**Optional:** `-selection` only if you have `xdotool`/`xclip` or `wtype`/`wl-clipboard`.

### Requirements

| Mode | Required |
|------|----------|
| Bash **Alt+T** / **Alt+E** | `bash`, built `layoutfix` binary, `~/.bashrc` snippet |
| `-selection` | X11: `xdotool`, `xclip` — or Wayland: `wtype`, `wl-clipboard` |

**Go 1.22+** — only to build from source.

### Install (bash / no root)

```bash
git clone git@github.com:Kizerfifas/layoutfix.git
cd layoutfix
chmod +x install.sh
./install.sh
source ~/.bashrc    # or open a new terminal
```

Installs:

- `~/.local/bin/layoutfix`
- `~/.config/layoutfix/layoutfix.bash` — binds **Alt+T** and **Alt+E** via readline
- adds `source` to `~/.bashrc` if missing

System-wide (optional):

```bash
sudo install -m 755 layoutfix /usr/local/bin/
```

### Usage in bash (no X11)

1. Type with wrong layout on the **current line**: `ыгвщ ыныеуьсед`
2. Press **Alt+T** or **Alt+E** (on Russian keyboard **Alt+е** is usually the same as Alt+T)
3. The line becomes: `sudo systemctl`

Works in:

- local terminal (TTY),
- SSH,
- `screen` / `tmux` (bash readline in the pane).

Does **not** need GNOME, KDE, or a global desktop shortcut.

Manual convert (pipe / CLI):

```bash
echo 'ыгвщ' | layoutfix -print
layoutfix -print -text 'руддщ'
layoutfix -print <<<"$READLINE_LINE"   # same as the hotkey handler
```

### Optional: highlighted text (`-selection`)

Needs a graphical stack (X11 or Wayland):

```bash
sudo apt install xdotool xclip    # X11
# or: sudo apt install wtype wl-clipboard   # Wayland
```

1. Select text in the terminal or editor.
2. Run `layoutfix -selection` (bind to a key in **your terminal emulator** or `xbindkeys` on X11).

Example **xbindkeys** (`config/xbindkeys.rc`):

```bash
xbindkeys -f /path/to/layoutfix/config/xbindkeys.rc
```

Command in that file must call `layoutfix -selection`, not plain `layoutfix`.

### What it does *not* do

- Does not fix typos on the same layout (`ыныеуьмед` ≠ `systemctl` if you pressed the wrong key, not the wrong layout).
- Bash mode converts the **whole current line**, not an arbitrary mouse selection (use `-selection` for that).

### Troubleshooting (bash)

| Problem | Fix |
|---------|-----|
| Alt+T / Alt+E does nothing | `source ~/.bashrc`; check `which layoutfix`; use **bash** (not plain `sh`). |
| Works in bash, not in `sh` | Binding is readline-specific; start `bash`. |
| `command not found: layoutfix` | Ensure `~/.local/bin` is in `PATH`. |

### Development

```bash
go test ./...
go build -ldflags="-s -w" -o layoutfix .
```

### Disclaimer

No license is provided. Use at **your own risk**. Authors assume **no liability** for damage or data loss.

---

## Русский

### Два режима

| Режим | Клавиши | Нужен X11? | Когда |
|-------|---------|------------|--------|
| **Bash (по умолчанию)** | **Alt+T** или **Alt+E** в shell | **Нет** | Текущая строка в bash — SSH, TTY, VM без графики |
| **Выделение** | Системная/терминальная команда → `layoutfix -selection` | **Да** (или Wayland) | Выделенный текст в графическом терминале или приложении |

**Для серверов и терминала без X:** установка + интеграция с bash (readline).  
**Опционально:** `-selection`, если установлены `xdotool`/`xclip` или `wtype`/`wl-clipboard`.

### Требования

| Режим | Нужно |
|-------|--------|
| Bash **Alt+T** / **Alt+E** | `bash`, бинарник `layoutfix`, фрагмент в `~/.bashrc` |
| `-selection` | X11: `xdotool`, `xclip` — или Wayland: `wtype`, `wl-clipboard` |

**Go 1.22+** — только для сборки.

### Установка (bash, без root)

```bash
git clone git@github.com:Kizerfifas/layoutfix.git
cd layoutfix
chmod +x install.sh
./install.sh
source ~/.bashrc    # или новый терминал
```

Ставится:

- `~/.local/bin/layoutfix`
- `~/.config/layoutfix/layoutfix.bash` — привязка **Alt+T** и **Alt+E** через readline
- `source` в `~/.bashrc`, если ещё не было

В систему (опционально):

```bash
sudo install -m 755 layoutfix /usr/local/bin/
```

### Использование в bash (без X11)

1. Наберите на **текущей строке**: `ыгвщ ыныеуьсед`
2. **Alt+T** или **Alt+E** (на русской раскладке **Alt+е** обычно то же, что Alt+T)
3. Строка станет: `sudo systemctl`

Работает в локальном TTY, по SSH, в `screen` / `tmux`.

**GNOME/KDE не нужны** — это не «системная клавиша рабочего стола», а **горячая клавиша bash (readline)** внутри терминала.

Вручную из CLI:

```bash
echo 'ыгвщ' | layoutfix -print
layoutfix -print -text 'руддщ'
```

### Опционально: выделенный текст (`-selection`)

Нужен графический стек (X11 или Wayland):

```bash
sudo apt install xdotool xclip
```

1. Выделите текст.
2. Запустите `layoutfix -selection` (привяжите в **настройках эмулятора терминала** или через `xbindkeys` на X11).

В `config/xbindkeys.rc` должна быть команда `layoutfix -selection`.

### Чего не делает

- Не исправляет опечатки на той же раскладке.
- В bash меняется **вся текущая строка**, не произвольное выделение мышью (для выделения — `-selection`).

### Решение проблем (bash)

| Симптом | Решение |
|---------|---------|
| Alt+T / Alt+E не работает | `source ~/.bashrc`; `which layoutfix`; запускайте **bash**, не `sh`. |
| Нет `layoutfix` | Добавьте `~/.local/bin` в `PATH`. |

### Разработка

```bash
go test ./...
go build -ldflags="-s -w" -o layoutfix .
```

### Отказ от ответственности

Лицензия **не предоставляется**. Использование **на свой страх и риск**. Авторы **не несут ответственности** за ущерб и потерю данных.
