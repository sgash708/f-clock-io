package cmd

import (
	"github.com/sgash708/f-clock-io/internal/application"
	"github.com/spf13/cobra"
)

var clockinCmd = &cobra.Command{
	Use:   "clockin",
	Short: "clock in a certain site",
	Run: func(cmd *cobra.Command, args []string) {
		application.RunClockIn()
	},
}

func init() {
	rootCmd.AddCommand(clockinCmd)
}
