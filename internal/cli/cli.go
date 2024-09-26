package cli

import (
	"bytes"
	"errors"
	"fmt"
	"go-medium2markdown/pkg/md2"
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

		buf := new(bytes.Buffer)
		// w := zip.NewWriter(buf)

		mco := md2.Options{}
		mco.IsDownloadAssets = true
		mc := md2.NewConverter(buf, mco)

		err := mc.Convert(args[0])
		if err != nil {
			log.Println(err)
		}

		// log.Println("Closing zip writer...")
		// err = w.Close()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		log.Println("Writing zip file to disk...")
		err = os.WriteFile("archive.zip", buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

		log.Printf("Zip file written successfully. Size: %d bytes", buf.Len())

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
