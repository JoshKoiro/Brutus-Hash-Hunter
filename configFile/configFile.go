package configFile

import (
	"brutus-hash-hunter/appio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

const configPath string = "./config.yaml"

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
		color.Red("\nFailed to unmarshal YAML config file: %v", err)
	}

	config.Wordlists = wordLists

}

func (config *Config) DownloadFiles() {
	var index int = 1
	for _, value := range config.Wordlists {
		if _, err := os.Stat("./wordlists/" + value.Name + ".txt"); err == nil {
			color.Cyan("(%v/%v) Already downloaded %v\n", index, len(config.Wordlists), value.Name)
		} else {
			color.HiBlue("(%v/%v) Downloading %v...\n", index, len(config.Wordlists), value.Name)
			appio.DownloadURL(value.Name, value.Link)
		}
		index++
	}
}
