package core

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MarkupSymbol struct {
	Italic  string `yaml:"italic"`
	Section string `yaml:"section"`
}

type MediumConverterOptions struct {
	MarkupSymbol     MarkupSymbol `yaml:"markupSymbol"`
	IsFrontMatter    bool         `yaml:"isFrontMatter"`
	IsDownloadAssets bool         `yaml:"isDownloadAssets"`
}

func NewMediumConverterOption() *MediumConverterOptions {
	opt := getDefaultOption()
	return &opt
}

func getDefaultOption() MediumConverterOptions {
	// set default config
	conf := MediumConverterOptions{
		MarkupSymbol{
			Italic:  "*",
			Section: "***",
		},
		false,
		false,
	}

	return conf
}

func readConfigFromFile(filePath string, config *MediumConverterOptions) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
	}
}
