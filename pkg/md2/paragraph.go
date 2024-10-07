package md2

import (
	"archive/zip"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
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
		var imageSize string
		imageLayout := p.Layout
		imageID := p.Metadata.Id
		imageURL := mediumCdnUrl + "/v2/format:webp/" + imageID
		imageAlt := p.Metadata.Alt

		if imageLayout != 0 {
			switch imageLayout {
			case 1:
				imageSize = "680"
			case 3:
				imageSize = "1110"
			}
		}

		if imageSize != "" {
			imageURL = mediumCdnUrl + fmt.Sprintf("/v2/resize:fit:%s/format:webp/", imageSize) + imageID
		}

		if p.Text != "" {
			caption = fmt.Sprintf("%s%s%s", opt.MarkupSymbol.Italic, mdText, opt.MarkupSymbol.Italic)
		}

		defaultImageURL := fmt.Sprintf("![%s](%s)", imageAlt, imageURL) + newLine + newLine + caption + newLine

		if opt.IsDownloadAssets {
			imageData, err := downloadFile(imageURL)
			if err != nil {
				return defaultImageURL
			}

			if len(imageData) == 0 {
				return defaultImageURL
			}

			imgPath := filepath.Join("assets", filepath.Base(p.Name+"_"+imageID))
			// imgWriter, err := w.Create(imgPath)
			imgWriter, err := w.CreateHeader(&zip.FileHeader{
				Name:     imgPath,
				Modified: time.Now(),
				Method:   zip.Store,
			})
			if err != nil {
				return defaultImageURL
			}

			// _, err = io.Copy(imgWriter, bytes.NewReader(imageData))
			n, err := imgWriter.Write(imageData)
			if err != nil {
				return defaultImageURL
			}

			if n != len(imageData) {
				return defaultImageURL
			}

			return fmt.Sprintf("![%s](%s)", imageAlt, "./assets/"+p.Name+"_"+imageID) + newLine + newLine + caption + newLine
		} else {
			return defaultImageURL
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
