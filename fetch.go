package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	splitParam = ";</x>"
	queryParam = "?format=json"
)

func fetchMediumPost(mediumUrl string) (MediumPost, error) {
	resp, err := http.Get(mediumUrl + queryParam)
	if err != nil {
		return MediumPost{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return MediumPost{}, err
	}

	post, err := toMediumPost(string(respBody))
	if err != nil {
		return MediumPost{}, err
	}

	return post, nil
}

func toMediumPost(raw string) (MediumPost, error) {
	var post MediumPost
	jsonStr, err := cleanPostResponse(raw)
	if err != nil {
		return post, err
	}

	if err = json.Unmarshal([]byte(jsonStr), &post); err != nil {
		log.Println(err)
		return MediumPost{}, errors.New("invalid json string")
	}

	return post, nil
}

// TODO: Potential issue if using split
// Handle it later
func cleanPostResponse(raw string) (string, error) {
	res := strings.Split(raw, splitParam)
	if len(res) != 2 {
		return "", errors.New("invalid Medium response")
	}
	return res[1], nil
}

// Embedded content response is only mediaResourceId
// so we need to fetch it manually from HTML
func fetchMediaResource(url string) (MediaResources, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML: %v", err)
		return nil, err
	}

	// Find the script tags and extract the one containing appolo state
	var scriptContent string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		if strings.Contains(content, "window.__APOLLO_STATE__") {
			scriptContent = content
			return
		}
	})

	if scriptContent == "" {
		return nil, errors.New("can't find apollo state script")
	}

	// Extract the value of the "MediaResource" object using regular expressions
	re := regexp.MustCompile(`"MediaResource:[^"]+":{[^}]+}`)
	matches := re.FindAllString(scriptContent, -1)
	if len(matches) == 0 {
		return nil, errors.New("failed to find the 'MediaResource' object")
	}

	var mrs map[string]MediaResource
	for _, mr := range matches {
		err = json.Unmarshal([]byte("{"+mr+"}"), &mrs)
		if err != nil {
			log.Printf("Error parsing JSON: %v", err)
			return nil, err
		}
	}
	log.Println(mrs)
	return mrs, nil
}
