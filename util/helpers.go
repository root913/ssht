package util

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

var storePassword string = ""

func GetStorePassword() string {
	if len(storePassword) != 0 {
		return storePassword
	}
	fmt.Print("Enter store password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println()
		return ""
	}
	fmt.Println()

	storePassword = string(bytePassword)
	return strings.TrimSpace(storePassword)
}

func AskPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println()
		return ""
	}
	fmt.Println()

	password := string(bytePassword)
	return strings.TrimSpace(password)
}

func AskKeyPath() string {
	fmt.Print("Enter key path: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println()
		return ""
	}
	fmt.Println()

	password := string(bytePassword)
	return strings.TrimSpace(password)
}

func GetHostAndUsername(hostAndUsername string) (string, string) {
	lastIndex := strings.LastIndex(hostAndUsername, "@")
	username := hostAndUsername[:lastIndex]
	host := hostAndUsername[lastIndex+1:]

	return host, username
}
