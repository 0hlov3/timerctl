/*
Copyright © 2025 0hlov3
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Shows the actual time in RFC850",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), time.Now().Format(time.RFC850))
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
