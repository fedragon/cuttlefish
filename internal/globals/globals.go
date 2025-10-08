package globals

import (
	"fmt"

	"github.com/fedragon/cuttlefish/internal/validate"
)

func Set(additions map[string]string, deletions []string) []string {
	cmds := make([]string, 0)

	for _, key := range deletions {
		if err := validate.Validate(key); err == nil {
			cmds = append(cmds, fmt.Sprintf("set --universal --erase %v", key))
		}
	}

	for key, value := range additions {
		if err := validate.Validate(key); err == nil {
			if err := validate.Validate(value); err == nil {
				cmds = append(cmds, fmt.Sprintf("set --universal --export %v %v", key, value))
			}
		}
	}

	return cmds
}
