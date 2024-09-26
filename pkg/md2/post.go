package md2

import (
	"archive/zip"
	"strings"
)

type OrderListCounter int

func (p *MediumPost) Parse(w *zip.Writer, opt Options) string {
	payload := p.Payload
	body := payload.Value.Content.BodyModel
	mentionedUsers := payload.MentionedUsers
	paras := body.Paragraphs
	secs := body.Sections

	processedParagraphs := make([]string, len(paras))

	if len(secs) > 0 {
		secs.Parse(processedParagraphs, opt)
	}

	if len(paras) > 0 {
		var counter OrderListCounter

		for i, para := range paras {
			splittedParagraph := strings.Split(para.Text, "")

			for _, m := range para.Markups {
				m.Parse(splittedParagraph, mentionedUsers, opt)
			}

			if para.Type == OrderedList {
				counter++
			} else {
				counter = 0
			}

			processedText := para.Parse(w, splittedParagraph, counter, mentionedUsers, opt)

			processedParagraphs[i] = processedParagraphs[i] + processedText
		}
	}

	return strings.Join(processedParagraphs, newLine+newLine)
}
