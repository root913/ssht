package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/root913/ssht/util"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

var appConfigInstance *Config = nil

// appConfig represents App configuration dir env var.
const appConfig = "appCONFIG"

var (
	// appConfigFile represents app config file location.
	AppConfigFile = filepath.Join(AppHome(), "config.yml")
	PassPath      = filepath.Join(AppHome(), "pass")
)

type (
	// Config tracks app configuration options.
	Config struct {
		App *App `yaml:"app"`
	}
)

func AppHome() string {
	if env := os.Getenv(appConfig); env != "" {
		return env
	}

	xdgAppHome, err := xdg.ConfigFile("ssht")
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("Unable to create configuration directory for app")
	}

	return xdgAppHome
}

func GetConfig() *Config {
	if appConfigInstance == nil {
		appConfigInstance = &Config{App: NewApp()}
		util.Logger.Debug().
			Str("config_path", AppHome()).
			Msg("Loading configuration.")

		if err := appConfigInstance.Load(AppConfigFile); err != nil {
			util.Logger.Debug().
				Str("config_file_path", AppConfigFile).
				Msg("Unable to locate App config. Generating new configuration.")
		}
	}

	return appConfigInstance
}

func (c *Config) Json() string {
	val, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

// Load App configuration from file.
func (c *Config) Load(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	c.App = NewApp()

	var cfg Config
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return err
	}
	if cfg.App != nil {
		c.App = cfg.App
	}
	return nil
}

// Save configuration to disk.
func (c *Config) Save() error {
	//TODO c.Validate()

	return c.SaveFile(AppConfigFile)
}

// SaveFile App configuration to disk.
func (c *Config) SaveFile(path string) error {
	EnsurePath(path, DefaultDirMod)
	cfg, err := yaml.Marshal(c)
	if err != nil {
		util.Logger.Error().Msgf("[Config] Unable to save app config file: %v", err)
		return err
	}
	return os.WriteFile(path, cfg, 0644)
}
