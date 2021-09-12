package cmd

import (
	"go-redis-example/config"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-redis-example",
	Short: "A minimal example of using Redis by Go",
}

// Execute for CLI.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVarP(
		&config.Config.ChannelName,
		"channel_name",
		"c",
		"default_chan",
		"Channel Name (required)",
	)
}
