package cfg

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"path/filepath"
)

type Cfg struct {
	Key string `yaml:"API_KEYS"`
}

func New() Cfg {
	data, err := readFile()
	var result Cfg

	if err != nil {
		result = getApiStringFromTerminal()
	} else {
		result = parseYmlFile(data)
	}

	return result
}

func readFile() ([]byte, error) {
	fileName, _ := filepath.Abs("./cfg.yml")
	return os.ReadFile(fileName)
}

func parseYmlFile(data []byte) Cfg {
	var cfg Cfg
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		cfg = getApiStringFromTerminal()
	}

	return cfg
}

func getApiStringFromTerminal() Cfg {
	var cfg Cfg
	for {
		fmt.Printf("yml file is invalid. Please enter a weather reports' api key's :\n")
		if _, err := fmt.Scanf("%s\n", &cfg.Key); err == nil && cfg.Key != "" {
			break
		}
	}

	return cfg
}
