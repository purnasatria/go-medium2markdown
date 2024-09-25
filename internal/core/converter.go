package core

import (
	"archive/zip"
	"log"
	"net/url"
)

type MediumConverter struct {
	Options *MediumConverterOptions
}

func NewMediumConverter(opt *MediumConverterOptions) *MediumConverter {
	return &MediumConverter{
		Options: opt,
	}
}

func (m *MediumConverter) Convert(url string, w *zip.Writer) error {
	log.Printf("Starting conversion for URL: %s", url)
	log.Printf("IsDownloadAssets: %v", m.Options.IsDownloadAssets)
	err := isValidURL(url)
	if err != nil {
		return err
	}

	mp, err := fetchMediumPost(url)
	if err != nil {
		return err
	}

	title := mp.Payload.Value.Slug

	media, err := fetchMediaResource(url)
	if err != nil {
		return err
	}

	log.Printf("Fetched Medium post. Title: %s", title)
	log.Printf("Number of paragraphs: %d", len(mp.Payload.Value.Content.BodyModel.Paragraphs))

	res := mp.Parse(w, m.Options, media)

	log.Printf("Parsing completed. Result length: %d", len(res))

	stringWriter, err := w.Create(title + ".md")
	if err != nil {
		log.Printf("Failed to create markdown entry in zip: %v", err)
		return err
	}

	_, err = stringWriter.Write([]byte(res))
	if err != nil {
		return err
	}

	log.Println("Markdown content added to zip successfully")

	return nil
}

func isValidURL(input string) error {
	_, err := url.Parse(input)
	if err != nil {
		return err
	}
	return nil
}
