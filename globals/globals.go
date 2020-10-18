package globals

import (
	"fmt"
)

func Set(additions map[string]string, deletions []string) []string {
	cmds := make([]string, 0)

	for _, key := range deletions {
		cmds = append(cmds, fmt.Sprintf("set --global --erase %v;", key))
	}

	for key, value := range additions {
		cmds = append(cmds, fmt.Sprintf("set --global --export %v %v;", key, value))
	}

	return cmds
}
