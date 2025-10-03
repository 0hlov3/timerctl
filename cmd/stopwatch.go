/*
Copyright © 2025 0hlov3
*/
package cmd

import (
	"context"
	"time"

	"github.com/0hlov3/timerctl/internal/models"
	"github.com/0hlov3/timerctl/internal/timer"
	"github.com/spf13/cobra"
)

var (
	swFlagTick     time.Duration
	swFlagNoMillis bool
	swFlagBeep     bool

	// Optional: auto-stop at a max duration (otherwise runs until CTRL+C)
	swFlagMax time.Duration

	// Optional progress bar (only visible if --max > 0)
	swFlagBar      bool
	swFlagBarWidth int
)

// stopwatchCmd represents the stopwatch command
var stopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "Count up from zero until CTRL+C (or until --max)",
	Long: `Count up from zero until CTRL+C (or until --max).

Examples:
  timerctl stopwatch
  timerctl stopwatch --no-ms
  timerctl stopwatch --max 30s --beep --bar --bar-width 40`,
	Example: `  timerctl stopwatch
  timerctl stopwatch --no-ms
  timerctl stopwatch --max 30s --beep --bar --bar-width 40`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := models.Options{
			Tick:       swFlagTick,
			ShowMillis: !swFlagNoMillis,
			BellOnDone: swFlagBeep,
			Max:        swFlagMax,
			ShowBar:    swFlagBar,
			BarWidth:   swFlagBarWidth,
			// (Output still goes to stdout; we print with fmt like the countdown)
		}
		return timer.RunStopwatch(context.Background(), opt)
	},
}

func init() {
	rootCmd.AddCommand(stopwatchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopwatchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopwatchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	stopwatchCmd.Flags().DurationVar(&swFlagTick, "tick", 100*time.Millisecond, "update interval (e.g. 100ms, 1s)")
	stopwatchCmd.Flags().BoolVar(&swFlagNoMillis, "no-ms", false, "Do not display milliseconds")
	stopwatchCmd.Flags().BoolVar(&swFlagBeep, "beep", false, "Terminal bell when stopped")

	stopwatchCmd.Flags().DurationVar(&swFlagMax, "max", 0, "Stop automatically at this duration (0 = infinite)")

	stopwatchCmd.Flags().BoolVar(&swFlagBar, "bar", false, "Show a progress bar (only meaningful if --max is set)")
	stopwatchCmd.Flags().IntVar(&swFlagBarWidth, "bar-width", 30, "Progress bar width (characters)")
}
