package main

import (
	"fmt"
	"ssht/config"
	"strings"
	"syscall"

	"github.com/jedib0t/go-pretty/table"
	"golang.org/x/term"
)

var storePassword string = ""

func getStorePassword() string {
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

func askPassword() string {
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

func askKeyPath() string {
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

func getHostAndUsername(hostAndUsername string) (string, string) {
	lastIndex := strings.LastIndex(hostAndUsername, "@")
	username := hostAndUsername[:lastIndex]
	host := hostAndUsername[lastIndex+1:]

	return host, username
}

func connectionTable(config *config.Config) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Alias", "Host", "Port", "Username", "Key", "Type"})
	for _, connection := range config.App.Connections {
		fmt.Print(connection.Host)
		fmt.Println()
		tw.AppendRow(table.Row{connection.Uuid, connection.Alias, connection.Host, connection.Port, connection.Username, connection.KeyPath, connection.Type})
	}
	tw.SetIndexColumn(1)

	fmt.Println(tw.Render())
}
