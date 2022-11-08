package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOkValidateKeyConnection(t *testing.T) {
	errs := NewKeyConnection("host.localhost", 22, "username", "/some/path", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidateKeyConnection(t *testing.T) {
	// empty host
	errs := NewKeyConnection("", 22, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewKeyConnection("host.localhost", 0, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// wrong port
	errs = NewKeyConnection("host.localhost", 99999999, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: must be valid port")
	// empty username
	errs = NewKeyConnection("host.localhost", 22, "", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
	// empty keyPath
	errs = NewKeyConnection("host.localhost", 22, "username", "", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "KeyPath: cannot be blank")
}

func TestOkValidateKeyPassphraseConnection(t *testing.T) {
	errs := NewKeyPassphraseConnection("host.localhost", 22, "username", "/some/path", "pass", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidateKeyPassphraseConnection(t *testing.T) {
	// empty host
	errs := NewKeyPassphraseConnection("", 22, "username", "/some/path", "pass", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewKeyPassphraseConnection("host.localhost", 0, "username", "/some/path", "pass", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// wrong port
	errs = NewKeyPassphraseConnection("host.localhost", 99999999, "username", "/some/path", "pass", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: must be valid port")
	// empty username
	errs = NewKeyPassphraseConnection("host.localhost", 22, "", "/some/path", "pass", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
	// empty keyPath
	errs = NewKeyPassphraseConnection("host.localhost", 22, "username", "", "pass", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "KeyPath: cannot be blank")
	// empty keyPass
	errs = NewKeyPassphraseConnection("host.localhost", 22, "username", "/some/path", "", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "KeyPass: cannot be blank")
}

func TestOkValidatePasswordConnection(t *testing.T) {
	errs := NewPasswordConnection("host.localhost", 22, "username", "/some/path", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidatePasswordConnection(t *testing.T) {
	// empty host
	errs := NewPasswordConnection("", 22, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewPasswordConnection("host.localhost", 0, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// wrong port
	errs = NewPasswordConnection("host.localhost", 99999999, "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: must be valid port")
	// empty username
	errs = NewPasswordConnection("host.localhost", 22, "", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
	// empty password
	errs = NewPasswordConnection("host.localhost", 22, "username", "", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Password: cannot be blank")
}
