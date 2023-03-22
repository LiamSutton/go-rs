package utils

import (
	"errors"
	"fmt"
	"strings"
)

func ReadUsername() (string, error) {

	var username string

	fmt.Print("Enter your username: ")
	fmt.Scanf("%s", &username)

	username = strings.Trim(username, "")

	if len(username) == 0 {
		return username, errors.New("No username provided.")
	}

	return username, nil
}
