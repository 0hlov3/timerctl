package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestStopwatch_HelpMentioned(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stopwatch", "--help"})
	_, err := rootCmd.ExecuteC()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "Count up from zero") {
		t.Fatalf("expected help to mention 'Count up from zero', got: %q", out)
	}
}
