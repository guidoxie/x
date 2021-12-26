package cmd

import (
	"github.com/guidoxie/x/internal/ssh"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "OpenSSH SSH client (remote login program)",
	Run: func(cmd *cobra.Command, args []string) {
		ssh.Exec(args...)
	},
}
