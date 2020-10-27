package ssh

import (
	"fmt"

	"github.com/fedragon/cuttlefish/common"
)

func Set(additions []string, deletions []string) []string {
	cmds := make([]string, 0)
	for _, identity := range deletions {
		if err := common.Validate(identity); err == nil {
			cmds = append(cmds, fmt.Sprintf("ssh-add -d %v 2> /dev/null", identity))
		}
	}
	for _, identity := range additions {
		if err := common.Validate(identity); err == nil {
			cmds = append(cmds, fmt.Sprintf("ssh-add %v", identity))
		}
	}

	return cmds
}
