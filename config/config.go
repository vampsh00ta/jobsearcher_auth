package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		Telegram `yaml:"telegram"`
		HTTP     `yaml:"http"`
		PG       `yaml:"postgres"`
	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Name string `yaml:"name" env:"HTTP_NAME"`
	}
	Telegram struct {
		APIToken string `env-required:"true" yaml:"name"    env:"API_TG_TOKEN"`
	}
	Log struct {
		Level string ` yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax  int    ` yaml:"pool_max" env:"PG_POOL_MAX"`
		Username string `env-required:"true" yaml:"username" env:"POSTGRES_USER"`
		Password string `env-required:"true" yaml:"password" env:"POSTGRES_PASSWORD"`
		Host     string `env-required:"true" yaml:"host" env:"POSTGRES_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"POSTGRES_PORT"`
		Name     string `env-required:"true" yaml:"name" env:"POSTGRES_DB"`
	}
)

func New(configPath string) (*Config, error) {
	cfg := &Config{}
	var err error

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err = godotenv.Load(filepath.Join(pwd, ".env")); err != nil {
		return nil, err
	}
	err = cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
