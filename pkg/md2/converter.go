package md2

import (
	"archive/zip"
	"bytes"
	"fmt"
	"time"
)

type Converter struct {
	Metadata MediumMetadata
	Buffer   *bytes.Buffer
	Options  Options
}

type MediumMetadata struct {
	Title    string
	Subtitle string
	Slug     string
}

func NewConverter(buf *bytes.Buffer, opt ...Options) *Converter {
	c := &Converter{
		Options: Options{},
		Buffer:  buf,
	}

	if len(opt) > 0 {
		custOpt := opt[0]
		c.Options = custOpt
	}

	defOpt := getDefaultOptions()

	if c.Options.MarkupSymbol.Section == "" {
		c.Options.MarkupSymbol.Section = defOpt.MarkupSymbol.Section
	}

	if c.Options.MarkupSymbol.Italic == "" {
		c.Options.MarkupSymbol.Italic = defOpt.MarkupSymbol.Italic
	}

	return c
}

func (c *Converter) Convert(mediumPostURL string) error {
	err := isValidURL(mediumPostURL)
	if err != nil {
		return err
	}

	var mp MediumPost
	res, err := callMediumAPI(mediumPostURL + mediumFormatJson)
	if err != nil {
		return err
	}

	toMediumPost(&mp, res)

	c.Metadata.Title = mp.Payload.Value.Title
	c.Metadata.Subtitle = mp.Payload.Value.Content.Subtitle
	c.Metadata.Slug = mp.Payload.Value.Slug

	if c.Options.IsDownloadAssets {
		w := zip.NewWriter(c.Buffer)
		// defer w.Close()

		_, err := w.Create("assets/")

		mpContent := mp.Parse(w, c.Options)

		filename := c.Metadata.Slug + ".md"
		// stringWriter, err := w.Create(c.Metadata.Slug + ".md")
		stringWriter, err := w.CreateHeader(&zip.FileHeader{
			Name:     filename,
			Modified: time.Now(),
		})
		if err != nil {
			return err
		}

		// _, err = io.Copy(stringWriter, bytes.NewBuffer([]byte(mpContent)))
		_, err = stringWriter.Write([]byte(mpContent))
		if err != nil {
			return err
		}
		fmt.Println("test1")

		if err = w.Close(); err != nil {
			fmt.Println("test")
			return err
		}
	} else {
		mpContent := mp.Parse(nil, c.Options)
		c.Buffer.Write([]byte(mpContent))
	}
	return nil
}
