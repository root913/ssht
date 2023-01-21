package config

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type ConnectionType string

const (
	PasswordConnection      ConnectionType = "password"
	KeyConnection           ConnectionType = "key"
	KeyPassphraseConnection ConnectionType = "key_passphrase"
)

func (c ConnectionType) String() string {
	switch c {
	case PasswordConnection:
		return "password"
	case KeyConnection:
		return "key"
	case KeyPassphraseConnection:
		return "key_passphrase"
	}
	return "unknown"
}

type Connection struct {
	Uuid      string `yaml:"uuid"`
	Alias     string `yaml:"alias"`
	Host      string `yaml:"host"`
	Port      int16  `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string
	KeyPath   string         `yaml:"keyPath"`
	KeyPass   string         `yaml:"keyPass"`
	Type      ConnectionType `yaml:"type"`
	CreatedAt time.Time      `yaml:"createdAt"`
	UpdatedAt time.Time      `yaml:"updatedAt"`
}

func getUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")[:12]
}

func newConnection(host string, port int16, username string, keyPath string, alias string, connectionType ConnectionType) *Connection {
	return &Connection{
		Uuid:      getUuid(),
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  "",
		KeyPath:   keyPath,
		KeyPass:   "",
		Type:      connectionType,
		Alias:     alias,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewPasswordConnection(host string, port int16, username string, alias string) *Connection {
	return newConnection(host, port, username, "", alias, PasswordConnection)
}

func NewKeyConnection(host string, port int16, username string, keyPath string, alias string) *Connection {
	return newConnection(host, port, username, keyPath, alias, KeyConnection)
}

func NewKeyPassphraseConnection(host string, port int16, username string, keyPath string, alias string) *Connection {
	return newConnection(host, port, username, keyPath, alias, KeyPassphraseConnection)
}

func isPort(value interface{}) error {
	s, _ := value.(int16)
	isPort := govalidator.IsPort(strconv.FormatInt(int64(s), 10))

	if isPort {
		return nil
	}
	return errors.New("must be valid port")
}

func (c *Connection) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.Host,
			validation.Required,
			is.Host,
		),
		validation.Field(
			&c.Port,
			validation.Required,
			validation.By(isPort),
		),
		validation.Field(
			&c.Username,
			validation.Required,
		),
		validation.Field(
			&c.Type,
			validation.Required,
		),
		validation.Field(
			&c.KeyPath,
			validation.Required.When(c.Type == KeyConnection || c.Type == KeyPassphraseConnection),
		),
	)
}

func (c *Connection) Json() string {
	val, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

func (c *Connection) String() string {
	return c.Json()
}
