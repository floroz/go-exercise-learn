package main

import (
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
}

var Users = [2]User{
	{Username: "jackie", Password: "1234"},
	{Username: "donald", Password: "4567"},
}

const (
	// Messages
	UsageMessage         = "Usage: [username] [password]"
	AccessDeniedMessage  = "Access denied"
	AccessGrantedMessage = "Access Granted!"
)

func findUser(username string) (*User, error) {
	for _, user := range Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func checkPassword(usr *User, password string) error {
	if usr.Password != password {
		return fmt.Errorf("invalid Password")
	}

	return nil
}

func main() {
	args := os.Args[1:] // remove the go invocation from first argument

	if len(args) < 2 {
		fmt.Println(UsageMessage)
		return
	}
	usrname, pwd := args[0], args[1]

	user, err := findUser(usrname)

	if err != nil {
		fmt.Println(AccessDeniedMessage)
		return
	}

	pwdErr := checkPassword(user, pwd)

	if pwdErr != nil {
		fmt.Println(AccessDeniedMessage)
		return
	}

	fmt.Println(AccessGrantedMessage)
}
