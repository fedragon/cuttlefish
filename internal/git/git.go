package git

import (
	"fmt"

	"github.com/fedragon/cuttlefish/internal/validate"
)

func SetName(name string) (string, bool) {
	if err := validate.Validate(name); err != nil {
		return "", false
	}
	return fmt.Sprintf(`git config --global user.name "%v"`, name), true
}

func SetEmail(email string) (string, bool) {
	if err := validate.Validate(email); err != nil {
		return "", false
	}

	return fmt.Sprintf(`git config --global user.email "%v"`, email), true
}
