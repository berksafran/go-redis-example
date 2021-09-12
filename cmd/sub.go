package cmd

import (
	"go-redis-example/subscribe"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Run with Subscription Mode",
	Run: func(cmd *cobra.Command, args []string) {
		subscribe.Sub()

		/*
			To get value of "channel_name" flag as a string:

			chanName := cmd.Flags().Lookup("channel_name").Value.String()
			 							OR
			chanName2, _ := cmd.Flags().GetString("channel_name")
		*/
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
