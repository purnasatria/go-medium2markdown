package core

import (
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

func (m *MediumConverter) Convert(url string) (string, error) {
	err := isValidURL(url)
	if err != nil {
		return "", err
	}

	mp, err := fetchMediumPost(url)
	if err != nil {
		return "", err
	}

	media, err := fetchMediaResource(url)
	if err != nil {
		return "", err
	}

	res := mp.Parse(m.Options, media)

	return res, nil
}

func isValidURL(input string) error {
	_, err := url.Parse(input)
	if err != nil {
		return err
	}
	return nil
}
