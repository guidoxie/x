package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "x",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.AddCommand(sshCmd, psCmd, ipCmd)
}
func Execute() {
	rootCmd.Execute()
}
