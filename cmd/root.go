package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "",
	Long: "",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.AddCommand(sshCmd, psCmd, ipCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
