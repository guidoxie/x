package ssh

import (
	"log"
	"os"
	"os/exec"
)

func Exec(args ...string) {
	sshClient := exec.Command("ssh", args...)
	sshClient.Stdin = os.Stdin
	sshClient.Stdout = os.Stdout
	sshClient.Stderr = os.Stderr
	if err := sshClient.Run(); err != nil {
		log.Fatal(err)
	}
}
