package md2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (p *Paragraph) isEmbedType() bool {
	return p.Type == Embed
}

func (p *Paragraph) fetcMediaResource() (*MediumMedia, error) {
	if !p.isEmbedType() {
		return nil, newError("Paragraph %s invalid type %d", p.Name, p.Type)
	}

	var mm MediumMedia

	res, err := callMediumAPI(mediumMediaUrl + p.Iframe.MediaResourceId + mediumJsonPrefix)
	if err != nil {
		return nil, err
	}

	if err = toMediumMedia(&mm, res); err != nil {
		return nil, err
	}

	return &mm, nil
}

func (p *Paragraph) embedDefault(mm *MediumMedia) string {
	return fmt.Sprintf("<iframe src=\"%s\" width=\"%d\" height=\"%d\"></iframe>\n[Original URL](%s)",
		fmt.Sprintf("%s%s", mediumMediaUrl, mm.Payload.Value.MediaResourceID),
		mm.Payload.Value.IframeWidth,
		mm.Payload.Value.IframeHeight,
		mm.Payload.Value.Href,
	)
}

type GistFile struct {
	Filename string `json:"filename"`
	Language string `json:"language"`
	RawURL   string `json:"raw_url"`
}

type GistResponse struct {
	Files map[string]GistFile `json:"files"`
}

func (p *Paragraph) embedGithubGist(mm *MediumMedia) (string, error) {
	url := fmt.Sprintf("%s%s", githubGistUrl, mm.Payload.Value.Gist.GistID)

	resp, err := http.Get(url)
	if err != nil {
		return "", newError("error fetching gist: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", newError("error reading response body: %w", err)
	}

	var gistJson GistResponse
	if err := json.Unmarshal(body, &gistJson); err != nil {
		return "", newError("error unmarshaling JSON: %w", err)
	}

	var res strings.Builder

	for _, file := range gistJson.Files {
		language := strings.ToLower(file.Language)
		fmt.Println("gistCode", file.RawURL)

		gistCodeResp, err := http.Get(file.RawURL)
		if err != nil {
			return "", newError("error fetching gist code: %w", err)
		}
		defer gistCodeResp.Body.Close()

		gistCode, err := io.ReadAll(gistCodeResp.Body)
		if err != nil {
			return "", newError("error reading gist code: %w", err)
		}

		res.WriteString(fmt.Sprintf("\n```%s\n", language))
		res.WriteString(strings.ReplaceAll(string(gistCode), "\t", "  "))
		res.WriteString(fmt.Sprintf("\n```\n[Original URL](%s)\n"))
	}

	return res.String(), nil
}
