package error

import "errors"

func CustomErrors(username string) (bool, error) {
	return true, errors.New(username + " is exist")
}
