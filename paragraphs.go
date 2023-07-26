package main

import (
	"fmt"
	"log"
)

//func (p *Paragraph) process(sp []string, counter OrderListCounter, mrs MediaResources, mus MentionedUsers) string {

func (p *Paragraph) process(cfg *Config, sp []string, counter OrderListCounter, mrs MediaResources, mus MentionedUsers) {
	//if p.Type != EmbeddedLink {
	//	for _, m := range p.Markups {
	//		m.process(sp, mus)
	//	}
	//}
	//
	//mdText := strings.Join(sp, "")

	switch p.Type {
	//case Basic:
	//	return mdText
	case BigT:
		sp[0] = "# " + sp[0]
		//return "# " + mdText
	case SmallT:
		sp[0] = "## " + sp[0]
		//return "## " + mdText
	case Quote:
		sp[0] = "> " + sp[0]
		//return "> " + mdText
	case Image:
		// TODO: handle download image
		//var caption string
		imageID := p.Metadata.Id
		imageURL := mediumCdn + "/v2/resize:fit:950/" + imageID
		imageAlt := p.Metadata.Alt

		if len(sp) == 0 {
			sp = append(sp, fmt.Sprintf("![%s](%s)", imageAlt, imageURL))
		} else {
			sp[0] = fmt.Sprintf("![%s](%s)", imageAlt, imageURL) + br + br + cfg.MarkupSymbol.Italic + sp[0]
			sp[len(sp)-1] += cfg.MarkupSymbol.Italic
		}

		//if p.Text != "" {
		//caption = fmt.Sprintf("%s%s%s", mdText)
		//}
		//return fmt.Sprintf("![%s](%s)%s", imageAlt, imageURL, br+br) + caption
	case CodeBlock:
		sp[0] = fmt.Sprintf("```%s%s", p.CodeBlockMetadata.Lang, br)
		sp[len(sp)-1] += br + "````"
		//return fmt.Sprintf("```%s\n%s\n```", p.CodeBlockMetadata.Lang, mdText) + br
	case UnOrderedList:
		sp[0] = "- " + sp[0]
		//return "- " + mdText
	case OrderedList:
		sp[0] = fmt.Sprintf("%d. ", counter) + sp[0]
		//return fmt.Sprintf("%d. %s", counter, mdText)
	case Embed:
		mr := mrs.getMediaResource(p.Iframe.MediaResourceId)

		if mr.IframeSrc != "" {
			//orUrl := mr.getOriginalURL()
			//return fmt.Sprintf("<iframe src=\"%s\" width=\"%d\" height=\"%d\"></iframe>\n[Original URL](%s)",
			//	mr.IframeSrc,
			//	mr.IframeWidth,
			//	mr.IframeHeight,
			//	orUrl,
			//)
			orUrl := mr.getOriginalURL()
			iframe := fmt.Sprintf("<iframe src=\"%s\" width=\"%d\" height=\"%d\"></iframe>\n[Original URL](%s)",
				mr.IframeSrc,
				mr.IframeWidth,
				mr.IframeHeight,
				orUrl,
			)

			sp = append(sp, iframe)
		}
		//return ""
	case EmbeddedLink:
		sp = append(sp, fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.Href, p.MixtapeMetadata.Href))
		//return fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.MediaResourceId, p.MixtapeMetadata.Href)
		//return p.MixtapeMetadata.Href
	default:
		log.Printf("name: %s unkown paragraph type %d\n", p.Name, p.Type)
		//return mdText
	}

}
