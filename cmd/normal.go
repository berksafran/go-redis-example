package cmd

import (
	"go-redis-example/normal"

	"github.com/spf13/cobra"
)

// normalCmd represents the normal command
var normalCmd = &cobra.Command{
	Use:   "normal",
	Short: "Run with Normal Mode including GET, SET, EXISTS examples",
	Run: func(cmd *cobra.Command, args []string) {
		normal.RunNormalMode()
	},
}

func init() {
	rootCmd.AddCommand(normalCmd)
}
