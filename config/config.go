package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-logr/logr"
	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

var Cfg DatabaseConfig

// LoadConfig loads configs from config file.
func LoadConfig(configFile string, log logr.Logger) error {
	log.Info("Load configs", "configFile", configFile)

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error(err, "Read config file")
		return err
	}

	if len(b) <= 0 {
		return fmt.Errorf("empty config file: %s", configFile)
	}
	err = yaml.Unmarshal(b, &Cfg)
	if err != nil {
		log.Error(err, "Unmarshal config file", "configFile", configFile)
		return err
	}
	log.Info("Configs loaded")
	return nil
}
