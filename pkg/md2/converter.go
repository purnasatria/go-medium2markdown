package md2

import (
	"archive/zip"
	"bytes"
	"log"
)

type Converter struct {
	Buffer  *bytes.Buffer
	Options Options
}

type CoverterWithZip struct {
	Converter
	zip.Writer
}

func NewConverter(buf *bytes.Buffer, opt ...Options) *Converter {
	c := &Converter{
		Options: getDefaultOptions(),
		Buffer:  buf,
	}

	if len(opt) > 0 {
		usedOpt := opt[0]
		c.Options = usedOpt
	}

	return c
}

func (c *Converter) Convert(mediumPostURL string) error {
	w := zip.NewWriter(c.Buffer)
	defer w.Close()

	log.Printf("IsDownloadAssets: %v", c.Options.IsDownloadAssets)
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

	title := mp.Payload.Value.Slug

	log.Printf("Fetched Medium post. Title: %s", title)
	log.Printf("Number of paragraphs: %d", len(mp.Payload.Value.Content.BodyModel.Paragraphs))

	mpContent := mp.Parse(w, c.Options)

	stringWriter, err := w.Create(title + ".md")
	if err != nil {
		log.Printf("Failed to create markdown entry in zip: %v", err)
		return err
	}

	_, err = stringWriter.Write([]byte(mpContent))
	if err != nil {
		return err
	}

	log.Println("Markdown content added to zip successfully")

	return nil
}

// type CloseableWriter struct {
// 	Writer  io.Writer
// 	Buffer  *bytes.Buffer
// 	cleanup func() error
// }
//
// func (cw *CloseableWriter) Close() error {
// 	if cw.cleanup != nil {
// 		return cw.cleanup()
// 	}
// 	return nil
// }
