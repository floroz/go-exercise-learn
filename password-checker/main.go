package main

import (
	"fmt"
	"os"
)

const (
	// Mock usr and pwd to compare against in a db call
	UserName = "jackie"
	Password = "1234"
	// Messages
	UsageMessage         = "Usage: [username] [password]"
	AccessDeniedMessage  = "Access denied"
	AccessGrantedMessage = "Access Granted!"
)

func getUserAndPassword() (string, string) {
	return UserName, Password
}

func main() {
	username, password := getUserAndPassword()

	var usr string
	var pwd string

	args := os.Args[1:] // remove the go invocation from first argument

	if len(args) < 2 {
		fmt.Println(UsageMessage)
		return
	}

	usr, pwd = args[0], args[1]

	if username != usr {
		fmt.Println(AccessDeniedMessage)
		return
	}

	if password != pwd {
		fmt.Println(AccessDeniedMessage)
		return
	}

	fmt.Println(AccessGrantedMessage)
}
