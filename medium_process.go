package main

import (
	"strings"
)

type OrderListCounter int

func (p *MediumPost) process(cfg *Config, mr MediaResources) string {
	payload := p.Payload
	body := payload.Value.Content.BodyModel
	mentionedUsers := payload.MentionedUsers
	paras := body.Paragraphs
	secs := body.Sections

	processedParagraphs := make([]string, len(paras))

	if len(secs) > 0 {
		secs.process(cfg, processedParagraphs)
	}

	if len(paras) > 0 {
		var counter OrderListCounter

		for i, para := range paras {
			splittedParagraph := strings.Split(para.Text, "")

			for _, m := range para.Markups {
				m.process(cfg, splittedParagraph, mentionedUsers)
			}

			if para.Type == OrderedList {
				counter++
			} else {
				counter = 0
			}

			processedText := para.process(cfg, splittedParagraph, counter, mr, mentionedUsers)

			processedParagraphs[i] = processedParagraphs[i] + processedText
		}
	}

	return strings.Join(processedParagraphs, br+br)
}
