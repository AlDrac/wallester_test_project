package configs

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	configPath     string
	configTestPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/yaml/api.yaml", "path to api config file")
	flag.StringVar(&configTestPath, "config-test-path", "../../../configs/yaml/api.yaml", "path to test api config file")
}

type (
	Database struct {
		Url     string `yaml:"url"`
		UrlTest string `yaml:"url_test"`
	}

	Http struct {
		Port string `yaml:"port"`
	}

	Logger struct {
		Level string `yaml:"level"`
	}

	Config struct {
		Database `yaml:"database"`
		Http     `yaml:"http"`
		Logger   `yaml:"logger"`
	}
)

func New(isTest bool) *Config {
	flag.Parse()

	path := configPath
	if isTest == true {
		path = configTestPath
	}

	configContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	configContent = []byte(os.ExpandEnv(string(configContent)))
	config := &Config{}
	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
