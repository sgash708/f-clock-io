package cmd

import (
	"github.com/sgash708/f-clock-io/internal/application"
	"github.com/spf13/cobra"
)

var startbreakCmd = &cobra.Command{
	Use:   "startbreak",
	Short: "start break",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := application.RunStartBreak(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startbreakCmd)
}
