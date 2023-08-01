package main

import (
	"fmt"
	"log"
	"strings"
)

func (p *Paragraph) process(cfg *Config, sp []string, counter OrderListCounter, mrs MediaResources, mus MentionedUsers) string {
	mdText := strings.Join(sp, "")

// func (p *Paragraph) process(cfg *Config, sp []string, counter OrderListCounter, mrs MediaResources, mus MentionedUsers) {
	//if p.Type != EmbeddedLink {
	//	for _, m := range p.Markups {
	//		m.process(sp, mus)
	//	}
	//}
	//
	//mdText := strings.Join(sp, "")

	switch p.Type {
	case Basic:
		return mdText

	case BigT:
		return "# " + mdText

	case SmallT:
		return "## " + mdText

	case Quote:
		return "> " + mdText

	case Image:
		// TODO: handle download image
		var caption string
		imageID := p.Metadata.Id
		imageURL := mediumCdn + "/v2/resize:fit:950/" + imageID
		imageAlt := p.Metadata.Alt

		if p.Text != "" {
			caption = fmt.Sprintf("%s%s%s", cfg.MarkupSymbol.Italic, mdText, cfg.MarkupSymbol.Italic)
		}

		return fmt.Sprintf("![%s](%s)%s", imageAlt, imageURL, br+br) + caption + br

	case CodeBlock:
		return fmt.Sprintf("```%s\n%s\n```", p.CodeBlockMetadata.Lang, mdText)

	case UnOrderedList:
		return "- " + mdText

	case OrderedList:
		return fmt.Sprintf("%d. %s", counter, mdText)

	case Embed:
		mr := mrs.getMediaResource(p.Iframe.MediaResourceId)

		if mr.IframeSrc != "" {
			orUrl := mr.getOriginalURL()
			return fmt.Sprintf("<iframe src=\"%s\" width=\"%d\" height=\"%d\"></iframe>\n[Original URL](%s)",
				mr.IframeSrc,
				mr.IframeWidth,
				mr.IframeHeight,
				orUrl,
			)
		}
		return ""

	case EmbeddedLink:
		// return fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.MediaResourceId, p.MixtapeMetadata.Href)
		return fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.Href, p.MixtapeMetadata.Href)
	default:
		log.Printf("name: %s unkown paragraph type %d\n", p.Name, p.Type)
		return mdText
	}

}
