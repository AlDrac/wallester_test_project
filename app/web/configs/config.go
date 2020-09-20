package configs

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/yaml/web.yaml", "path to web config file")
}

type (
	Http struct {
		Port   string `yaml:"port"`
		ApiUrl string `yaml:"api_url"`
	}

	Template struct {
		Layout  string `yaml:"layout"`
		Include string `yaml:"include"`
	}

	Session struct {
		Key string `yaml:"key"`
	}

	Config struct {
		Http     `yaml:"http"`
		Template `yaml:"template"`
		Session  `yaml:"session"`
	}
)

func New() *Config {
	flag.Parse()

	configContent, err := ioutil.ReadFile(configPath)
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
