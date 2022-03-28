package ps

import (
	"log"
	"os"
	"os/exec"
)

func Exec(process string) {
	var err error
	ps := exec.Command("bash", "-c", "ps -ef | grep -v grep | grep -v 'x ps'")
	grep := exec.Command("grep", process)

	grep.Stdin, err = ps.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return
	}
	grep.Stdout = os.Stdout
	grep.Stderr = os.Stderr

	for _, f := range []func() error{
		grep.Start,
		ps.Run,
	} {
		if err := f(); err != nil {
			log.Fatal(err)
			return
		}
	}
}
