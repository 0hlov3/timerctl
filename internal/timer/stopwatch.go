package timer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/0hlov3/timerctl/internal/models"
	"github.com/0hlov3/timerctl/internal/notification"
)

func RunStopwatch(ctx context.Context, opt models.Options) error {
	if opt.Tick <= 0 {
		opt.Tick = 100 * time.Millisecond
	}
	if opt.BarWidth <= 0 {
		opt.BarWidth = 30
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	start := time.Now()
	ticker := time.NewTicker(opt.Tick)
	defer ticker.Stop()

	lastLen := 0
	printLine := func(s string) {
		if len(s) < lastLen {
			fmt.Printf("\r%s%s", s, strings.Repeat(" ", lastLen-len(s)))
		} else {
			fmt.Printf("\r%s", s)
		}
		lastLen = len(s)
	}

	formatElapsed := func(d time.Duration) string {
		hrs := int(d / time.Hour)
		mins := int((d % time.Hour) / time.Minute)
		secs := int((d % time.Minute) / time.Second)
		ms := int((d % time.Second) / time.Millisecond)

		if hrs > 0 {
			if opt.ShowMillis {
				return fmt.Sprintf("%02d:%02d:%02d.%03d", hrs, mins, secs, ms)
			}
			return fmt.Sprintf("%02d:%02d:%02d", hrs, mins, secs)
		}
		if opt.ShowMillis {
			return fmt.Sprintf("%02d:%02d.%03d", mins, secs, ms)
		}
		return fmt.Sprintf("%02d:%02d", mins, secs)
	}

	for {
		select {
		case <-ctx.Done():
			elapsed := time.Since(start)
			fmt.Printf("\rStopped at %s.\n", formatElapsed(elapsed))
			return context.Canceled

		case <-ticker.C:
			elapsed := time.Since(start)

			// auto-stop if Max set
			if opt.Max > 0 && elapsed >= opt.Max {
				if opt.BellOnDone {
					notification.Notify(fmt.Sprintf("Stopwatch reached %s", opt.Max))
				}
				fmt.Printf("\rDone! %s\n", formatElapsed(opt.Max))
				return nil
			}

			line := formatElapsed(elapsed)

			// Optional progress bar if Max is known
			if opt.ShowBar && opt.Max > 0 {
				frac := float64(elapsed) / float64(opt.Max)
				if frac < 0 {
					frac = 0
				}
				if frac > 1 {
					frac = 1
				}
				filled := int(frac * float64(opt.BarWidth))
				if filled > opt.BarWidth {
					filled = opt.BarWidth
				}
				bar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", opt.BarWidth-filled) + "]"
				line = fmt.Sprintf("%s %s", bar, line)
			}

			printLine(line)
		}
	}
}
