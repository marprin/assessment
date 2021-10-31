package config

import (
	"time"

	pkgConfig "github.com/marprin/assessment/fetchapp/pkg/config"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		App      AppConfig      `yaml:"app"`
		Token    TokenConfig    `yaml:"token"`
		Gateway  GatewayConfig  `yaml:"gateway"`
		Currency CurrencyConfig `yaml:"currency"`
	}

	AppConfig struct {
		Host            string        `yaml:"host"`
		Port            int           `yaml:"port"`
		GracefulTimeout time.Duration `yaml:"graceful_timeout"`
		ReadTimeout     time.Duration `yaml:"read_timeout"`
		WriteTimeout    time.Duration `yaml:"write_timeout"`
		Stage           string        `yaml:"stage"`
	}

	TokenConfig struct {
		Secret string `yaml:"secret"`
		Issuer string `yaml:"issuer"`
	}

	GatewayConfig struct {
		URL string `yaml:"url"`
	}

	CurrencyConfig struct {
		URL    string `yaml:"url"`
		APIKey string `yaml:"api_key"`
	}
)

func ReadConfig(cfg interface{}, filename string) error {
	configPath := "./config"
	err := pkgConfig.ReadYMLConfig(configPath, filename, cfg)
	if err != nil {
		logrus.Fatalln("Failed to read config")
	}
	return nil
}
