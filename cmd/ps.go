package cmd

import (
	"github.com/guidoxie/x/internal/ps"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps COMMAND",
	Short: "Filter process status",
	Args:  cobra.ExactArgs(1), // 必须有一个参数
	Run: func(cmd *cobra.Command, args []string) {
		ps.Exec(args[0])
	},
}
