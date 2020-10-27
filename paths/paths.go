package paths

import (
	"fmt"

	"github.com/fedragon/cuttlefish/common"
)

func Set(additions []string, deletions []string) []string {
	cmds := make([]string, 0)
	for _, path := range deletions {
		if err := common.Validate(path); err == nil {
			cmds = append(cmds, fmt.Sprintf(`
      if set -l index (contains -i %v $fish_user_paths)
        set --erase fish_user_paths[$index]
			end
		`, path))
		}
	}

	for _, path := range additions {
		if err := common.Validate(path); err == nil {
			cmds = append(cmds, fmt.Sprintf("contains %v $fish_user_paths; or set --global --export --append fish_user_paths %v", path, path))
		}
	}

	return cmds
}
