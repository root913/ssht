package client

import (
	// "bytes"
	// "errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"time"

	"ssht/config"
	"ssht/util"

	"golang.org/x/crypto/ssh"
)

// A Client implements an SSH client that supports running commands and scripts remotely.
type Client struct {
	client *ssh.Client
}

// DialWithPasswd starts a client connection to the given SSH server with passwd authmethod.
func DialWithPasswd(addr, user, passwd string) (*Client, error) {
	util.Logger.Debug().
		Str("addr", addr).
		Str("user", user).
		Str("passwd", passwd).
		Msg("DialWithPasswd")

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
			ssh.KeyboardInteractive(PasswordKeyboardInteractive(passwd)),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
		Timeout:         5 * time.Second,
	}

	return Dial("tcp", addr, config)
}

// DialWithKey starts a client connection to the given SSH server with key authmethod.
func DialWithKey(addr, user, keyfile string) (*Client, error) {
	util.Logger.Debug().
		Str("addr", addr).
		Str("user", user).
		Str("keyfile", keyfile).
		Msg("DialWithKey")

	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
		Timeout:         5 * time.Second,
	}

	return Dial("tcp", addr, config)
}

// DialWithKeyWithPassphrase same as DialWithKey but with a passphrase to decrypt the private key
func DialWithKeyWithPassphrase(addr, user, keyfile string, passphrase string) (*Client, error) {
	util.Logger.Debug().
		Str("addr", addr).
		Str("user", user).
		Str("keyfile", keyfile).
		Str("passphrase", passphrase).
		Msg("DialWithKeyWithPassphrase")

	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(passphrase))
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
		Timeout:         5 * time.Second,
	}

	return Dial("tcp", addr, config)
}

// Dial starts a client connection to the given SSH server.
// This wraps ssh.Dial.
func Dial(network, addr string, config *ssh.ClientConfig) (*Client, error) {
	client, err := ssh.Dial(network, addr, config)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

func PasswordKeyboardInteractive(password string) ssh.KeyboardInteractiveChallenge {
	return func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		// Just send the password back for all questions
		answers := make([]string, len(questions))
		for i, question := range answers {
			util.Logger.Debug().
				Str("question", question).
				Msg("KeyboardInteractive")

			answers[i] = string(password)
		}

		return answers, nil
	}
}

// Close closes the underlying client network connection.
func (c *Client) Close() error {
	util.Logger.Debug().Msg("Closing connection")
	return c.client.Close()
}

// UnderlyingClient get the underlying client.
func (c *Client) UnderlyingClient() *ssh.Client {
	return c.client
}

// A RemoteShell represents a login shell on the client.
type RemoteShell struct {
	client         *ssh.Client
	requestPty     bool
	terminalConfig *TerminalConfig

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// A TerminalConfig represents the configuration for an interactive shell session.
type TerminalConfig struct {
	Term   string
	Height int
	Weight int
	Modes  ssh.TerminalModes
}

// Terminal create a interactive shell on client.
func (c *Client) Terminal(config *TerminalConfig) *RemoteShell {
	return &RemoteShell{
		client:         c.client,
		terminalConfig: config,
		requestPty:     true,
	}
}

// Shell create a noninteractive shell on client.
func (c *Client) Shell() *RemoteShell {
	return &RemoteShell{
		client:     c.client,
		requestPty: false,
	}
}

// SetStdio specifies where the its standard output and error data will be written.
func (rs *RemoteShell) SetStdio(stdin io.Reader, stdout, stderr io.Writer) *RemoteShell {
	rs.stdin = stdin
	rs.stdout = stdout
	rs.stderr = stderr
	return rs
}

// Start starts a remote shell on client.
func (rs *RemoteShell) Start() error {
	session, err := rs.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if rs.stdin == nil {
		session.Stdin = os.Stdin
	} else {
		session.Stdin = rs.stdin
	}
	if rs.stdout == nil {
		session.Stdout = os.Stdout
	} else {
		session.Stdout = rs.stdout
	}
	if rs.stderr == nil {
		session.Stderr = os.Stderr
	} else {
		session.Stderr = rs.stderr
	}

	if rs.requestPty {
		tc := rs.terminalConfig
		if tc == nil {
			tc = &TerminalConfig{
				Term:   "xterm-256color",
				Height: 40,
				Weight: 80,
			}
		}
		if err := session.RequestPty(tc.Term, tc.Height, tc.Weight, tc.Modes); err != nil {
			return err
		}
	}

	if err := session.Shell(); err != nil {
		return err
	}

	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}

func newClient(connection config.Connection) (*Client, error) {
	util.Logger.Debug().
		Str("addr", connection.Host+":"+fmt.Sprint(connection.Port)).
		Str("user", connection.Username).
		Str("passwd", connection.Password).
		Str("type", connection.Type.String()).
		Str("keyPath", connection.KeyPath).
		Msg("newClient")

	switch connection.Type {
	case config.PasswordConnection:
		return DialWithPasswd(connection.Host+":"+fmt.Sprint(connection.Port), connection.Username, connection.Password)
	case config.KeyConnection:
		return DialWithKey(connection.Host+":"+fmt.Sprint(connection.Port), connection.Username, connection.KeyPath)
	case config.KeyPassphraseConnection:
		return DialWithKeyWithPassphrase(connection.Host+":"+fmt.Sprint(connection.Port), connection.Username, connection.KeyPath, connection.KeyPass)
	default:
		return DialWithPasswd(connection.Host+":"+fmt.Sprint(connection.Port), connection.Username, connection.Password)
	}
}

func Connect(connection *config.Connection) {
	util.Logger.Info().
		Str("host", connection.Host).
		Int("port", connection.Port).
		Msg("Conneting")

	client, err := newClient(*connection)

	if err != nil {
		util.Logger.Fatal().Err(err).Msg("failed to connect")

		return
	}
	if client == nil {
		util.Logger.Fatal().Msg("timeout")

		return
	}

	defer client.Close()

	config := &TerminalConfig{
		Term:   "xterm-256color",
		Height: 40,
		Weight: 80,
		Modes: ssh.TerminalModes{
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}
	if err := client.Terminal(config).Start(); err != nil {
		util.Logger.Fatal().Err(err)
	}
}
