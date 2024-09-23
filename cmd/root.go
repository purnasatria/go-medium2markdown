package cmd

import (
	"errors"
	"fmt"
	"go-medium2markdown/internal/core"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	mediumUrl  string
	cfgFile    string
	outputFile string
)

var rootCmd = &cobra.Command{
	Use:   "md2 [medium_url]",
	Short: "md2 is program to covert your Medium post to markdown format",
	Long:  `md2 is program built with Go to covert your Medium post to markdown format`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please input Medium url")
		}

		mco := core.NewMediumConverterOption()
		mc := core.NewMediumConverter(mco)

		res, err := mc.Convert(args[0])
		if err != nil {
			log.Println(err)
		}

		fmt.Println(res)

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Config file path")
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "Output File")
}

func Exceute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error occured: %e", err)
		os.Exit(1)
	}
}
