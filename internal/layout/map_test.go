package layout

import "testing"

func TestConvert_RuToEn(t *testing.T) {
	cases := map[string]string{
		"ды":                  "ls",
		"ыгвщ ыныеуьсед":      "sudo systemctl",
		"руддщ":               "hello",
		"GHBDTN":              "ПРИВЕТ",
	}
	for in, want := range cases {
		if got := Convert(in); got != want {
			t.Errorf("Convert(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestConvert_RoundTrip(t *testing.T) {
	in := "sudo systemctl --status nginx"
	got := Convert(Convert(in))
	if got != in {
		t.Errorf("round trip: got %q, want %q", got, in)
	}
}
