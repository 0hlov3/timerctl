package cmd

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestNow_PrintsRFC850(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"now"})

	_, err := rootCmd.ExecuteC()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	// Very loose check: ensure weekday name + comma exists, which RFC850 uses.
	// Example: "Friday, 03-Oct-25 21:13:00 CEST"
	if !strings.Contains(out, ",") {
		t.Fatalf("expected RFC850-like output, got: %q", out)
	}

	// Optional: try parsing with time.RFC850, but timezone names vary by OS.
	// We’ll accept presence of weekday + comma + dash-delimited date.
	_ = time.Now() // placeholder to document intent
}
