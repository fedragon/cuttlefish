package ssh

import (
	"fmt"
)

func Set(additions []string, deletions []string) []string {
	cmds := make([]string, 0)
	for _, identity := range deletions {
		cmds = append(cmds, fmt.Sprintf("ssh-add -d %v 2> /dev/null", identity))
	}
	for _, identity := range additions {
		cmds = append(cmds, fmt.Sprintf("ssh-add %v", identity))
	}

	return cmds
}
