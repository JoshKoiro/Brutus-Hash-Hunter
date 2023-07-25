package configFile

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/ui"
	"fmt"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

const configPath string = "./config/config.yaml"

type Config struct {
	Filepath  string              `yaml:"-"`
	Wordlists map[string]Wordlist `yaml:"-"`
}

func (config *Config) Initialize() {
	config.Filepath = configPath
	configFile, err := os.ReadFile(config.Filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var wordLists = make(map[string]Wordlist)
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
			color.Cyan("(%v/%v) %v - Already downloaded", index, len(config.Wordlists), value.Name)
		} else {
			color.HiBlue("(%v/%v) %v", index, len(config.Wordlists), value.Name)
			ui.Spinner("Downloading...")
			appio.DownloadURL(value.Name, value.Link)
			fmt.Printf(" Downloaded\n")
		}
		index++
	}
}

type Wordlist struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}
