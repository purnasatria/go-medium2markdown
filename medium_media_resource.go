package main

import (
	"fmt"
	"net/url"
)

type MediaResource struct {
	Typename     string `json:"__typename"`
	ID           string `json:"id"`
	IframeSrc    string `json:"iframeSrc"`
	IframeHeight int    `json:"iframeHeight"`
	IframeWidth  int    `json:"iframeWidth"`
	Title        string `json:"title"`
}

type MediaResources map[string]MediaResource

func (m MediaResources) getMediaResource(id string) MediaResource {
	return m[fmt.Sprintf("MediaResource:%s", id)]
}

func (m *MediaResource) getOriginalURL() string {
	u, err := url.Parse(m.IframeSrc)
	if err != nil {
		return ""
	}

	queryParams := u.Query()
	urlValue := queryParams.Get("url")

	return urlValue
}
