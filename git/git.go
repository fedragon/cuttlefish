package git

import (
	"fmt"

	"github.com/fedragon/cuttlefish/common"
)

func SetIdentity(email string) string {
	err := common.Validate(email)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(`git config --global user.email "%v"`, email)
}
