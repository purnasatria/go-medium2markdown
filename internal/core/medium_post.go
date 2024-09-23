package core

import (
	"strings"
)

type OrderListCounter int

func (p *MediumPost) Parse(opt *MediumConverterOptions, mr MediaResources) string {
	payload := p.Payload
	body := payload.Value.Content.BodyModel
	mentionedUsers := payload.MentionedUsers
	paras := body.Paragraphs
	secs := body.Sections

	processedParagraphs := make([]string, len(paras))

	if len(secs) > 0 {
		secs.Parse(opt, processedParagraphs)
	}

	if len(paras) > 0 {
		var counter OrderListCounter

		for i, para := range paras {
			splittedParagraph := strings.Split(para.Text, "")

			for _, m := range para.Markups {
				m.Parse(opt, splittedParagraph, mentionedUsers)
			}

			if para.Type == OrderedList {
				counter++
			} else {
				counter = 0
			}

			processedText := para.Parse(opt, splittedParagraph, counter, mr, mentionedUsers)

			processedParagraphs[i] = processedParagraphs[i] + processedText
		}
	}

	return strings.Join(processedParagraphs, newLine+newLine)
}
