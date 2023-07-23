package configFile

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath string = "./config.yml"

type Config struct {
	Filepath  string              `yaml:"-"`
	Wordlists map[string]WordList `yaml:"-"`
}

type WordList struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

func (config *Config) Initialize() {
	config.Filepath = configPath
	configFile, err := os.ReadFile(config.Filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var wordLists = make(map[string]WordList)
	err = yaml.Unmarshal(configFile, &wordLists)
	if err != nil {
		fmt.Printf("\nFailed to unmarshal YAML config file: %v", err)
	}

	config.Wordlists = wordLists

}
