package ps

import (
	"fmt"
	"log"
	"os/exec"
)

func Exec(process string) {
	var ps *exec.Cmd
	if len(process) > 0 {
		ps = exec.Command("ps", "-ef", "|", "grep", process)
	} else {
		ps = exec.Command("ps", "-ef")
	}

	out, err := ps.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
