package git

import (
	"fmt"
)

func SetIdentity(email string) string {
	return fmt.Sprintf(`git config --global user.email "%v"`, email)
}
