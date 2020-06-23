package kubectl

import (
	"fmt"
	"os/exec"
	"strings"
)

func findFreePort(from int, to int) int {
	port := from
	for port < to {
		cmd := fmt.Sprintf("lsof -i -P -n | grep LISTEN | grep :%d", port)
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic(err)
		}
		if strings.Trim(string(out), ` \n`) == "" {
			return port
		}
		port++
	}
	return 0
}
