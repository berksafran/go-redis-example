package cmd

import (
	"go-redis-example/publisher"
	"strings"

	"github.com/spf13/cobra"
)

var Message string

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Run with Publisher Mode",
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")

		builder := strings.Builder{}
		builder.WriteString(message)

		for _, value := range args {
			builder.WriteString(" " + value)
		}

		totalMessage := builder.String()
		publisher.Publish(totalMessage)
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)

	pubCmd.Flags().StringVarP(
		&Message,
		"message",
		"m",
		"",
		"Publish a message (required)",
	)

	pubCmd.MarkFlagRequired("message")
}
