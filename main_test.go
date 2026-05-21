package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunPrintCLI(t *testing.T) {
	bin := buildTestBinary(t)
	cmd := exec.Command(bin, "-print", "-text", "ыгвщ ыныеуьсед")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	if got := string(out); got != "sudo systemctl" {
		t.Fatalf("got %q", got)
	}
}

func TestRunPrintStdin(t *testing.T) {
	bin := buildTestBinary(t)
	cmd := exec.Command(bin, "-print")
	cmd.Stdin = bytes.NewBufferString("ды\n")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	if got := string(out); got != "ls\n" {
		t.Fatalf("got %q", got)
	}
}

func buildTestBinary(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	bin := filepath.Join(dir, "layoutfix")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	return bin
}

func TestMain(m *testing.M) {
	// Ensure tests run from module root.
	if wd, err := os.Getwd(); err == nil {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err != nil {
			_ = os.Chdir(filepath.Dir(wd))
		}
	}
	os.Exit(m.Run())
}
