package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/halfbakedio/saas/constants"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env     string        `mapstructure:"env"`
	DB      DBConfig      `mapstructure:"db"`
	Log     LogConfig     `mapstructure:"log"`
	Service ServiceConfig `mapstructure:"service"`
}

const envPrefix = "SAAS"

var (
	lock     = &sync.Mutex{}
	instance *Config
)

var defaults = map[string]interface{}{
	"env":           constants.Development.String(),
	"db.host":       "127.0.0.1",
	"db.port":       5432,
	"db.name":       "saas",
	"db.user":       "postgres",
	"db.password":   "postgres",
	"db.sslmode":    "disable",
	"log.formatter": "text",
	"log.level":     "debug",
	"service.bind":  "0.0.0.0",
	"service.port":  9000,
}

func prepare(name string) (*viper.Viper, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	envConfig := fmt.Sprintf("%s_CONFIG", envPrefix)

	config := viper.New()

	file := os.Getenv(envConfig)
	if file == "" {
		config.SetConfigName(name)
		config.AddConfigPath(".")
		config.AddConfigPath(fmt.Sprintf("%s/.config/saas", home))
		config.AddConfigPath("/etc/saas")
	} else {
		var extension string
		regex := regexp.MustCompile("((y(a)?ml)|json|toml)$")
		base := filepath.Base(file)
		if regex.Match([]byte(base)) {
			// strip the file type for viper
			parts := strings.Split(filepath.Base(file), ".")
			base = strings.Join(parts[:len(parts)-1], ".")
			extension = parts[len(parts)-1]
		} else {
			return nil, errors.New("configuration does not support that extension type")
		}
		config.SetConfigName(base)
		config.SetConfigType(extension)
		config.SetConfigFile(file)
		config.AddConfigPath(filepath.Dir(file))
	}

	return config, nil
}

// LoadConfigWithDefaults reads in a configuration file from a set of
// locations, applies any defaults that have been provided, and deserializes it
// into a Config instance.
func LoadConfigWithDefaults(name string, c interface{}, defaults map[string]interface{}) error {
	config, err := prepare(name)
	if err != nil {
		return err
	}

	err = config.ReadInConfig()
	if err != nil {
		return err
	}

	config.SetEnvPrefix(envPrefix)
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	for key, value := range defaults {
		log.Tracef("Setting default: %s = %s\n", key, value)
		config.SetDefault(key, value)
	}

	err = config.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig returns the application configuration singleton.
func GetConfig() *Config {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			if err := LoadConfigWithDefaults("saas", &instance, defaults); err != nil {
				log.Fatalf("error reading config file: %s\n", err)
			}
		}
	}

	log.Tracef("config: %+v", instance)

	return instance
}
