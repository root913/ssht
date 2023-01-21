package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOkValidateKeyConnection(t *testing.T) {
	errs := NewKeyConnection("host.localhost", int16(22), "username", "/some/path", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidateKeyConnection(t *testing.T) {
	// empty host
	errs := NewKeyConnection("", int16(22), "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewKeyConnection("host.localhost", int16(0), "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// empty username
	errs = NewKeyConnection("host.localhost", int16(22), "", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
	// empty keyPath
	errs = NewKeyConnection("host.localhost", int16(22), "username", "", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "KeyPath: cannot be blank")
}

func TestOkValidateKeyPassphraseConnection(t *testing.T) {
	errs := NewKeyPassphraseConnection("host.localhost", int16(22), "username", "/some/path", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidateKeyPassphraseConnection(t *testing.T) {
	// empty host
	errs := NewKeyPassphraseConnection("", int16(22), "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewKeyPassphraseConnection("host.localhost", int16(0), "username", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// empty username
	errs = NewKeyPassphraseConnection("host.localhost", int16(22), "", "/some/path", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
	// empty keyPath
	errs = NewKeyPassphraseConnection("host.localhost", int16(22), "username", "", "").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "KeyPath: cannot be blank")
}

func TestOkValidatePasswordConnection(t *testing.T) {
	errs := NewPasswordConnection("host.localhost", int16(22), "username", "alias").Validate()
	assert.Nil(t, errs)
}

func TestFailValidatePasswordConnection(t *testing.T) {
	// empty host
	errs := NewPasswordConnection("", int16(22), "username", "alias").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Host: cannot be blank")
	// empty port
	errs = NewPasswordConnection("host.localhost", int16(0), "username", "alias").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Port: cannot be blank")
	// empty username
	errs = NewPasswordConnection("host.localhost", int16(22), "", "alias").Validate()
	assert.NotNil(t, errs)
	assert.Contains(t, errs.Error(), "Username: cannot be blank")
}
