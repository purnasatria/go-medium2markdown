package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var (
	mediumUrl  string
	cfgFile    string
	outputFile string
)

var cmd = &cobra.Command{
	Use:   "md2 [medium_url]",
	Short: "md2 is program to covert your Medium post to markdown format",
	Long:  `md2 is program built with Go to covert your Medium post to markdown format`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please input Medium url")
		}

		//load default config
		cfg := getDefaultConfig()

		// load config from file
		if cfgFile != "" {
			readConfigFromFile(cfgFile, &cfg)
		}

		mediumUrl = args[0]
		if err := isValidMediumURL(mediumUrl); err != nil {
			return err
		}

		post, err := fetchMediumPost(mediumUrl)
		if err != nil {
			return err
		}

		media, err := fetchMediaResource(mediumUrl)
		if err != nil {
			return err
		}

		result := post.process(&cfg, media)
		fmt.Println(result)

		return nil
	},
}

func init() {
	cmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Config file path")
	cmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "Output File")
}

func isValidMediumURL(input string) error {
	_, err := url.Parse(input)
	if err != nil {
		return err
	}

	/* qp := u.Query()
	format := qp.Get("format")
	if format == "" {
		return errors.New("url must contain query param 'format'")
	}

	if format != "json" {
		return errors.New("format must be json")
	} */

	return nil
}
