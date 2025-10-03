package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func execCmd(args ...string) (string, error) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	_, err := rootCmd.ExecuteC()
	return buf.String(), err
}

func TestSet_NoArgs_ShowsHelp(t *testing.T) {
	out, err := execCmd("set")
	if err != nil {
		// cobra.Help() path returns nil, so any error means parsing failed
		t.Fatalf("unexpected error: %v (out=%q)", err, out)
	}
	if !(strings.Contains(out, "Usage:") || strings.Contains(out, "timerctl set 25m")) {
		t.Fatalf("expected help text to contain usage or example, got: %q", out)
	}
}

func TestSet_InvalidDuration_ReturnsError(t *testing.T) {
	_, err := execCmd("set", "not-a-duration")
	if err == nil {
		t.Fatalf("expected error for invalid duration")
	}
}

func TestSet_ValidDuration_ParsesFlags(t *testing.T) {
	// We only assert that command executes without cobra parsing errors.
	// The actual countdown is tested in internal/timer tests.
	_, err := execCmd("set", "100ms", "--tick", "10ms", "--no-ms", "--beep=false")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}
