package credentials

import (
	"fmt"

	"github.com/root913/ssht/util"

	"github.com/99designs/keyring"
)

type Credentials struct {
	keyring keyring.Keyring
}

func NewCredentials(passPath string) *Credentials {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: "ssht",
		//AllowedBackends:  []keyring.BackendType{keyring.FileBackend, keyring.},
		KeyCtlScope:      "user",
		KeyCtlPerm:       0x3f3f0000,
		FileDir:          passPath,
		FilePasswordFunc: keyring.TerminalPrompt,
	})
	if err != nil {
		util.Logger.Err(err).Msg("")
	}
	return &Credentials{keyring: ring}
}

func (c *Credentials) Get(service string, connectionType string, username string) (string, error) {
	serviceName := getServiceName(service, connectionType, username)
	secret, err := c.keyring.Get(serviceName)
	if err != nil {
		util.Logger.
			Err(err).
			Str("service name", serviceName).
			Msg("")

		return "", err
	}

	return string(secret.Data), nil
}

func (c *Credentials) Set(service string, connectionType string, username string, password string) error {
	serviceName := getServiceName(service, connectionType, username)
	err := c.keyring.Set(keyring.Item{
		Key:  serviceName,
		Data: []byte(password),
	})
	if err != nil {
		util.Logger.
			Err(err).
			Str("service name", serviceName).
			Msg("")

		return err
	}

	return nil
}

func (c *Credentials) Destroy(service string, connectionType string, username string) error {
	serviceName := getServiceName(service, connectionType, username)
	err := c.keyring.Remove(serviceName)
	if err != nil {
		util.Logger.
			Err(err).
			Str("service name", serviceName).
			Msg("")

		return err
	}

	return nil
}

func getServiceName(service string, connectionType string, username string) string {
	return fmt.Sprintf("%s|%s|%s", service, connectionType, username)
}
