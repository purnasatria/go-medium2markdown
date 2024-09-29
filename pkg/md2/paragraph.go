package md2

import (
	"archive/zip"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func (p *Paragraph) Parse(w *zip.Writer, sp []string, counter OrderListCounter, mus MentionedUsers, opt Options) string {
	mdText := strings.Join(sp, "")

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
		var caption string
		imageID := p.Metadata.Id
		imageURL := mediumCdnUrl + "/v2/" + imageID
		imageAlt := p.Metadata.Alt

		if p.Text != "" {
			caption = fmt.Sprintf("%s%s%s", opt.MarkupSymbol.Italic, mdText, opt.MarkupSymbol.Italic)
		}

		imageLink := fmt.Sprintf("![%s](%s)%s", imageAlt, imageURL, newLine) + caption + newLine

		if opt.IsDownloadAssets {
			imageData, err := downloadFile(imageURL)
			if err != nil {
				return imageLink
			}

			if len(imageData) == 0 {
				return imageLink
			}

			imgPath := filepath.Join("assets", p.Name+"_"+imageID)
			imgWriter, err := w.Create(imgPath)
			if err != nil {
				return imageLink
			}

			n, err := imgWriter.Write(imageData)
			if err != nil {
				return imageLink
			}
			if n != len(imageData) {
				return imageLink
			}

			return fmt.Sprintf("![%s](%s)%s", imageAlt, "./assets/"+p.Name+"_"+imageID, newLine) + caption + newLine
		} else {
			return imageLink
		}

	case CodeBlock:
		return fmt.Sprintf("```%s\n%s\n```", p.CodeBlockMetadata.Lang, mdText)

	case UnOrderedList:
		return "- " + mdText

	case OrderedList:
		return fmt.Sprintf("%d. %s", counter, mdText)

	case Embed:
		return p.handleEmbed()

	case EmbeddedLink:
		return fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.Href, p.MixtapeMetadata.Href)
	default:
		log.Printf("name: %s unkown paragraph type %d\n", p.Name, p.Type)
		return mdText
	}
}
