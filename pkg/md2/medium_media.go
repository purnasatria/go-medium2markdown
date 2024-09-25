package md2

import (
	"encoding/json"
	"errors"
)

type MediumMedia struct {
	B       string       `json:"b"`
	Payload MediaPayload `json:"payload"`
	V       int          `json:"v"`
	Success bool         `json:"success"`
}

type MediaPayload struct {
	References struct{}          `json:"references"`
	Value      MediaPayloadValue `json:"value"`
}

type MediaPayloadValue struct {
	Tweet             Tweet  `json:"tweet"`
	Gist              Gist   `json:"gist"`
	MediaResourceID   string `json:"mediaResourceId"`
	MediaResourceType string `json:"mediaResourceType"`
	Href              string `json:"href"`
	Domain            string `json:"domain"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	IframeSrc         string `json:"iframeSrc"`
	ThumbnailURL      string `json:"thumbnailUrl"`
	ThumbnailImageID  string `json:"thumbnailImageId"`
	AuthorName        string `json:"authorName"`
	Type              string `json:"type"`
	IframeWidth       int    `json:"iframeWidth"`
	IframeHeight      int    `json:"iframeHeight"`
	ThumbnailWidth    int    `json:"thumbnailWidth"`
	ThumbnailHeight   int    `json:"thumbnailHeight"`
	Display           int    `json:"display"`
}

type Gist struct {
	GistID         string `json:"gistId"`
	GithubUsername string `json:"githubUsername"`
	GistScriptURL  string `json:"gistScriptUrl"`
}

type Tweet struct {
	TweetID           string `json:"tweetId"`
	Tweet             string `json:"tweet"`
	TwitterScreenName string `json:"twitterScreenName"`
	TwitterName       string `json:"twitterName"`
}

func toMediumMedia(mm *MediumMedia, response []byte) error {
	if err := json.Unmarshal(response, mm); err != nil {
		return errors.New("invalid media resource resonse")
	}
	return nil
}
