package timer_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/0hlov3/timerctl/internal/timer"
)

// captureStdout captures stdout during fn, then restores it.
func captureStdout(fn func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = orig
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestRun_ErrOnNonPositiveDuration(t *testing.T) {
	opt := timer.Options{Tick: 10 * time.Millisecond, ShowMillis: true}
	if err := timer.Run(context.Background(), 0, opt); err == nil {
		t.Fatalf("expected error for 0 duration, got nil")
	}
	if err := timer.Run(context.Background(), -1*time.Second, opt); err == nil {
		t.Fatalf("expected error for negative duration, got nil")
	}
}

func TestRun_DefaultTickApplied(t *testing.T) {
	// Tick <= 0 should default to 100ms (we just assert it doesn't panic or block strangely)
	opt := timer.Options{Tick: 0, ShowMillis: true}
	out := captureStdout(func() {
		_ = timer.Run(context.Background(), 10*time.Millisecond, opt)
	})
	if !strings.Contains(out, "Done!") {
		t.Fatalf("expected output to contain 'Done!', got: %q", out)
	}
}

func TestRun_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	opt := timer.Options{Tick: 5 * time.Millisecond, ShowMillis: true}

	// Cancel immediately
	cancel()
	out := captureStdout(func() {
		err := timer.Run(ctx, 200*time.Millisecond, opt)
		if err == nil {
			t.Fatalf("expected context.Canceled, got nil")
		}
		if err != context.Canceled {
			t.Fatalf("expected context.Canceled, got %v", err)
		}
	})

	if !strings.Contains(out, "Canceled.") {
		t.Fatalf("expected output to contain 'Canceled.', got: %q", out)
	}
}

func TestRun_CompletesAndPrintsDone(t *testing.T) {
	opt := timer.Options{
		Tick:       5 * time.Millisecond,
		ShowMillis: true,
		BellOnDone: false, // avoid desktop notification / bell in tests
	}

	out := captureStdout(func() {
		err := timer.Run(context.Background(), 30*time.Millisecond, opt)
		if err != nil {
			t.Fatalf("Run returned error: %v", err)
		}
	})

	if !strings.Contains(out, "Done!") {
		t.Fatalf("expected output to contain 'Done!', got: %q", out)
	}
}
