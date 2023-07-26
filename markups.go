package main

import (
	"fmt"
	"log"
)

func (m *Markup) process(cfg *Config, splittedParagraph []string, mus MentionedUsers) {
	switch m.Type {
	case Bold:
		m.addMarkup(splittedParagraph, "**", "**")
	case Italic:
		m.addMarkup(splittedParagraph, cfg.MarkupSymbol.Italic, cfg.MarkupSymbol.Italic)
	case LinkOrMention:
		if m.UserId == "" {
			m.addMarkup(splittedParagraph, "[", fmt.Sprintf("](%s)", m.Href))
		} else {
			user := mus.getUserData(m.UserId)
			if user.UserID != "" {
				m.addMarkup(splittedParagraph, "[", fmt.Sprintf("](medium.com/@%s)", user.Username))
			}
		}
	case Highlight:
		m.addMarkup(splittedParagraph, "`", "`")
	default:
		log.Printf("id: %s unkown markup type %d\n", m.Title, m.Type)
	}
}

func (m *Markup) addMarkup(arrText []string, startSymbol, endSymbol string) {
	if startSymbol != "" {
		arrText[m.Start] = startSymbol + arrText[m.Start]
	}

	if endSymbol != "" {
		arrText[m.End-1] += endSymbol
	}
}
