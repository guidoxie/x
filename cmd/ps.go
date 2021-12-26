package cmd

import (
	"github.com/guidoxie/x/internal/ps"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Filter process status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			ps.Exec(args[0])
		} else {
			ps.Exec("")
		}
	},
}
