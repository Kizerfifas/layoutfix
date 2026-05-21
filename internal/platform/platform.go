package platform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	copyDelay  = 120 * time.Millisecond
	pasteDelay = 80 * time.Millisecond
)

func SessionType() string {
	if v := os.Getenv("XDG_SESSION_TYPE"); v != "" {
		return strings.ToLower(v)
	}
	return "x11"
}

func ClipboardGet() (string, error) {
	switch SessionType() {
	case "wayland":
		out, err := exec.Command("wl-paste", "-n").CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("wl-paste: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return string(out), nil
	default:
		out, err := exec.Command("xclip", "-o", "-selection", "clipboard").CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("xclip: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return string(out), nil
	}
}

func ClipboardSet(text string) error {
	switch SessionType() {
	case "wayland":
		cmd := exec.Command("wl-copy", "-n")
		cmd.Stdin = strings.NewReader(text)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("wl-copy: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return nil
	default:
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(text)
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("xclip: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return nil
	}
}

func SimulateCopy() error {
	return simulateKey("ctrl+c")
}

func SimulatePaste() error {
	return simulateKey("ctrl+v")
}

func simulateKey(keys string) error {
	switch SessionType() {
	case "wayland":
		parts := strings.Split(keys, "+")
		args := make([]string, 0, len(parts)*2)
		for i := 0; i < len(parts)-1; i++ {
			args = append(args, "-M", strings.ToLower(parts[i]))
		}
		args = append(args, strings.ToLower(parts[len(parts)-1]))
		out, err := exec.Command("wtype", args...).CombinedOutput()
		if err != nil {
			return fmt.Errorf("wtype: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return nil
	default:
		out, err := exec.Command("xdotool", "key", "--clearmodifiers", keys).CombinedOutput()
		if err != nil {
			return fmt.Errorf("xdotool: %w (%s)", err, strings.TrimSpace(string(out)))
		}
		return nil
	}
}

// FixSelection copies the current selection, converts layout, and pastes back.
func FixSelection(convert func(string) string) error {
	old, _ := ClipboardGet()
	_ = ClipboardSet("")

	time.Sleep(pasteDelay)
	if err := SimulateCopy(); err != nil {
		_ = ClipboardSet(old)
		return err
	}
	time.Sleep(copyDelay)

	selected, err := ClipboardGet()
	if err != nil {
		_ = ClipboardSet(old)
		return err
	}
	if strings.TrimSpace(selected) == "" {
		_ = ClipboardSet(old)
		return fmt.Errorf("ничего не выделено: выделите текст и нажмите Alt+T")
	}

	converted := convert(selected)
	if err := ClipboardSet(converted); err != nil {
		_ = ClipboardSet(old)
		return err
	}

	time.Sleep(pasteDelay)
	if err := SimulatePaste(); err != nil {
		_ = ClipboardSet(old)
		return err
	}

	return nil
}

func CheckDependencies() error {
	var missing []string
	session := SessionType()
	if session == "wayland" {
		for _, bin := range []string{"wl-copy", "wl-paste", "wtype"} {
			if _, err := exec.LookPath(bin); err != nil {
				missing = append(missing, bin)
			}
		}
	} else {
		for _, bin := range []string{"xdotool", "xclip"} {
			if _, err := exec.LookPath(bin); err != nil {
				missing = append(missing, bin)
			}
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("не найдены утилиты: %s (sudo apt install xdotool xclip  # X11, или wl-clipboard wtype  # Wayland)", strings.Join(missing, ", "))
	}
	return nil
}
