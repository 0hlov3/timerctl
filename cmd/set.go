/*
Copyright © 2025 0hlov3
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/0hlov3/timerctl/internal/timer"
)

var (
	flagTick     time.Duration
	flagNoMillis bool
	flagBeep     bool
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Start a countdown (e.g. 25m, 10s, 1h5m)",
	Long: `timerctl set 25m
  timerctl set 90s --no-ms
  timerctl set 10s --tick 200ms --beep`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			_ = cmd.Help()
			return nil
		}
		d, err := time.ParseDuration(args[0])
		if err != nil || d <= 0 {
			return fmt.Errorf("invalid duration: %q", args[0])
		}

		opt := timer.Options{
			Tick:       flagTick,
			ShowMillis: !flagNoMillis,
			BellOnDone: flagBeep,
		}
		return timer.Run(context.Background(), d, opt)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().DurationVar(&flagTick, "tick", 100*time.Millisecond, "update interval (e.g. 100ms, 1s)")
	setCmd.Flags().BoolVar(&flagNoMillis, "no-ms", false, "Do not display milliseconds")
	setCmd.Flags().BoolVar(&flagBeep, "beep", false, "Terminal bell when expired")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
