package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mranderson/layoutfix/internal/layout"
	"github.com/mranderson/layoutfix/internal/platform"
)

func main() {
	printMode := flag.Bool("print", false, "convert text to stdout (no X11/Wayland, for bash/readline)")
	text := flag.String("text", "", "text to convert (with -print); empty = read stdin")
	selection := flag.Bool("selection", false, "convert highlighted text via clipboard (needs xdotool/xclip or Wayland tools)")
	flag.Parse()

	switch {
	case *printMode || *text != "":
		runPrint(*text)
	case *selection:
		runSelection()
	default:
		// Default: bash/TTY-friendly convert mode (no X11 required).
		runPrint(*text)
	}
}

func runPrint(arg string) {
	var in string
	switch {
	case arg != "":
		in = arg
	default:
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		in = string(b)
	}
	// Preserve trailing newline from stdin (readline often sends a line).
	hasNL := strings.HasSuffix(in, "\n")
	in = strings.TrimSuffix(in, "\n")
	out := layout.Convert(in)
	if hasNL {
		out += "\n"
	}
	fmt.Print(out)
}

func runSelection() {
	if err := platform.CheckDependencies(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := platform.FixSelection(layout.Convert); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
