package timer_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/0hlov3/timerctl/internal/models"
	"github.com/0hlov3/timerctl/internal/timer"
)

func capStdout(fn func()) string {
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

func TestStopwatch_AutoStopsAtMax(t *testing.T) {
	out := capStdout(func() {
		err := timer.RunStopwatch(context.Background(), models.Options{
			Tick:       5 * time.Millisecond,
			ShowMillis: true,
			BellOnDone: false,
			Max:        30 * time.Millisecond,
			ShowBar:    true,
			BarWidth:   10,
		})
		if err != nil {
			t.Fatalf("RunStopwatch returned error: %v", err)
		}
	})
	if !strings.Contains(out, "Done!") {
		t.Fatalf("expected 'Done!' in output, got: %q", out)
	}
}

func TestStopwatch_CancelPrintsStopped(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// cancel immediately
	cancel()
	out := capStdout(func() {
		err := timer.RunStopwatch(ctx, models.Options{
			Tick:       5 * time.Millisecond,
			ShowMillis: false,
		})
		if err == nil {
			t.Fatalf("expected context.Canceled, got nil")
		}
	})
	if !strings.Contains(out, "Stopped at") {
		t.Fatalf("expected 'Stopped at' in output, got: %q", out)
	}
}
