package cmd

import (
	"github.com/guidoxie/x/internal/ip"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Configure network interface parameters",
	Run: func(cmd *cobra.Command, args []string) {
		ip.Exec()
	},
}
