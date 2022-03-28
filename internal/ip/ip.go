package ip

import (
	"fmt"
	"log"
	"os/exec"
)

func Exec() {
	var cmds = map[string][]string{
		"ip":       {"route"},
		"ifconfig": {},
		"ipconfig": {},
	}
	for cmd, args := range cmds {
		if _, err := exec.LookPath(cmd); err == nil {
			cmd := exec.Command(cmd, args...)
			res, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(res))
			break
		}
	}
}
