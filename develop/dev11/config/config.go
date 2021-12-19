package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

// Config of service
type Config struct {
	DBcon   string
	WebPort string
}

// ParseConfig of service
func ParseConfig(configPath string) (*Config, error) {
	fileBody, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBody, &cfg)
	if err != nil {
		return nil, err
	}

	sign := getCredential("secrets/secret")
	if sign == "" {
		logrus.Warn("invalid signature")
	}
	cfg.DBcon = sign

	return &cfg, nil
}

func getCredential(path string) string {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Warn(err)
	}
	return strings.TrimSpace(string(c))
}
