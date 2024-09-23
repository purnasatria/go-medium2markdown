package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MarkupSymbol struct {
	Italic  string `yaml:"italic"`
	Section string `yaml:"section"`
}

type Config struct {
	MarkupSymbol     MarkupSymbol `yaml:"markupSymbol"`
	IsFrontMatter    bool         `yaml:"isFrontMatter"`
	IsDownloadAssets bool         `yaml:"isDownloadAssets"`
}

/* Load the configuration, the highest priority loaded last
 * First: Initialise to default config
 * Second: Replace with environment variables
 * Third: Replace with configuration file
 */
func getDefaultConfig() Config {
	// set default config
	conf := Config{
		MarkupSymbol{
			Italic:  "*",
			Section: "***",
		},
		false,
		false,
	}

	return conf
}

func readConfigFromFile(filePath string, config *Config) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
	}
}
