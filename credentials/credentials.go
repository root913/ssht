package credentials

import (
	"fmt"
)

type Credentials struct{}

func NewCredentials() *Credentials {
	return &Credentials{}
}

func (c *Credentials) Get(service string, username string) (string, error) {
	return "", nil
	// secret, err := keyring.Get(c.getServiceName(service), username)
	// if err != nil {
	// 	log.Fatal().Err(err)

	// 	return "", err
	// }

	// return secret, nil
}

func (c *Credentials) Set(service string, username string, password string) error {
	// err := keyring.Set(c.getServiceName(service), username, password)
	// if err != nil {
	// 	log.Fatal().Err(err)

	// 	return err
	// }

	return nil
}

func (c *Credentials) Destroy(service string, username string) error {
	// err := keyring.Delete(service, username)
	// if err != nil {
	// 	log.Fatal().Err(err)

	// 	return err
	// }

	return nil
}

func (c *Credentials) getServiceName(service string) string {
	return fmt.Sprintf("%s|%s", "ssht", service)
}
