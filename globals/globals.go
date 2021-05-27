package globals

import (
	"fmt"

	"github.com/fedragon/cuttlefish/common"
)

func Set(additions map[string]string, deletions []string) []string {
	cmds := make([]string, 0)

	for _, key := range deletions {
		if err := common.Validate(key); err == nil {
			cmds = append(cmds, fmt.Sprintf("set --universal --erase %v", key))
		}
	}

	for key, value := range additions {
		if err := common.Validate(key); err == nil {
			if err := common.Validate(value); err == nil {
				cmds = append(cmds, fmt.Sprintf("set --universal --export %v %v", key, value))
			}
		}
	}

	return cmds
}
