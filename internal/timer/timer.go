package timer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gen2brain/beeep"
)

type Options struct {
	Tick       time.Duration
	ShowMillis bool
	BellOnDone bool
}

func Run(ctx context.Context, d time.Duration, opt Options) error {
	if d <= 0 {
		return fmt.Errorf("duration must be > 0")
	}
	if opt.Tick <= 0 {
		opt.Tick = 100 * time.Millisecond
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	ticker := time.NewTicker(opt.Tick)
	defer ticker.Stop()

	deadline := time.Now().Add(d)

	for {
		select {
		case <-ctx.Done():
			fmt.Print("\rCanceled.\n")
			return context.Canceled

		case <-ticker.C:
			remaining := time.Until(deadline)
			if remaining <= 0 {
				if opt.BellOnDone {
					fmt.Print("\a")
					err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
					var icon []byte
					err = beeep.Notify("timerctl", fmt.Sprintf("The %s Timer ended", d.String()), icon)
					if err != nil {
						fmt.Println("\rcould not send Notification")
					}
				}
				fmt.Print("\rDone!\n")
				return nil
			}

			hrs := int(remaining / time.Hour)
			mins := int((remaining % time.Hour) / time.Minute)
			secs := int((remaining % time.Minute) / time.Second)
			ms := int((remaining % time.Second) / time.Millisecond)

			if hrs > 0 {
				if opt.ShowMillis {
					fmt.Printf("\r%02d:%02d:%02d.%03d", hrs, mins, secs, ms)
				} else {
					fmt.Printf("\r%02d:%02d:%02d   ", hrs, mins, secs)
				}
			} else {
				if opt.ShowMillis {
					fmt.Printf("\r%02d:%02d.%03d", mins, secs, ms)
				} else {
					fmt.Printf("\r%02d:%02d   ", mins, secs)
				}
			}
		}
	}
}
