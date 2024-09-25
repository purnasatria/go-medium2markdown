package cmd

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"go-medium2markdown/internal/core"
	"log"
	"os"
	"time"

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
		mco.IsDownloadAssets = true
		mc := core.NewMediumConverter(mco)

		buf := new(bytes.Buffer)
		w := zip.NewWriter(buf)

		// Before adding any files to the assets folder
		assetsFolder := &zip.FileHeader{
			Name:     "assets/",
			Method:   zip.Store, // directories should use Store method
			Modified: time.Now(),
		}
		assetsFolder.SetMode(0755) // This sets permissions and ensures it's not hidden
		_, err := w.CreateHeader(assetsFolder)
		if err != nil {
			log.Printf("Failed to create assets folder in zip: %v", err)
			return err
		}

		err = mc.Convert(args[0], w)
		if err != nil {
			log.Println(err)
		}

		log.Println("Closing zip writer...")
		err = w.Close()
		if err != nil {
			log.Fatal(err)
		}

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
