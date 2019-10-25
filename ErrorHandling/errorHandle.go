package main

import (
	"errors"
	"fmt"
)

type errUserNameExist struct {
	Username string
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("Username %s already existed", e.Username)
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}

func checkUserNameExisted(username string) (bool, error) {
	if username == "foo" {
		return true, errUserNameExist{Username: username}
	}

	if username == "bar" {
		return true, errors.New("Username bar already existed")
	}

	return false, nil
}

func main() {
	if _, err := checkUserNameExisted("foo"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		}
	}

	if _, err := checkUserNameExisted("bar"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
	}
}
