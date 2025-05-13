package git

import (
	"fmt"

	"github.com/fedragon/cuttlefish/common"
)

func SetName(name string) string {
	err := common.Validate(name)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(`git config --global user.name "%v"`, name)
}

func SetEmail(email string) string {
	err := common.Validate(email)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(`git config --global user.email "%v"`, email)
}
