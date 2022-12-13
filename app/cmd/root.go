/*
Copyright © 2022 Tahir BULBROOK <t@hir.me.uk>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "metric",
	Short: "Data analysis for listening data",
	Long: `metic imports and analyses listening data to produce chats and insights
to discover forgotten music.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
