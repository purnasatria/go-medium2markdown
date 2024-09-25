package core

import (
	"archive/zip"
	"fmt"
	"go-medium2markdown/internal/pkg"
	"log"
	"path/filepath"
	"strings"
)

func (p *Paragraph) Parse(w *zip.Writer, opt *MediumConverterOptions, sp []string, counter OrderListCounter, mrs MediaResources, mus MentionedUsers) string {
	log.Printf("Parsing paragraph. Type: %v, IsDownloadAssets: %v", p.Type, opt.IsDownloadAssets)
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
		// TODO: handle download image
		var caption string
		imageID := p.Metadata.Id
		imageURL := mediumCdn + "/v2/resize:fit:950/" + imageID
		imageAlt := p.Metadata.Alt
		log.Printf("Processing image. ID: %s, URL: %s", p.Metadata.Id, imageURL)

		if p.Text != "" {
			caption = fmt.Sprintf("%s%s%s", opt.MarkupSymbol.Italic, mdText, opt.MarkupSymbol.Italic)
		}

		imageLink := fmt.Sprintf("![%s](%s)%s", imageAlt, imageURL, newLine) + caption + newLine

		if opt.IsDownloadAssets {
			log.Println("Attempting to download image...")
			imageData, err := pkg.DownloadFile(imageURL)
			if err != nil {
				log.Printf("Failed to download image: %v", err)
				return imageLink
			}

			log.Printf("Image downloaded. Size: %d bytes", len(imageData))

			if len(imageData) == 0 {
				log.Println("Downloaded image data is empty")
				return imageLink
			}

			imgPath := filepath.Join("assets", p.Name+"_"+imageID)
			imgWriter, err := w.Create(imgPath)
			if err != nil {
				log.Printf("Failed to create zip entry for image: %s", imageURL)
				return imageLink
			}

			n, err := imgWriter.Write(imageData)
			if err != nil {
				log.Printf("Failed to write image to zip: %v", err)
				return imageLink
			}
			if n != len(imageData) {
				log.Printf("failed to write all image data: wrote %d of %d bytes", n, len(imageData))
				return imageLink
			}

			log.Printf("Successfully added image to zip: %s", imgPath)
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
		return fmt.Sprintf("[%s](%s)", p.MixtapeMetadata.Href, p.MixtapeMetadata.Href)
	default:
		log.Printf("name: %s unkown paragraph type %d\n", p.Name, p.Type)
		return mdText
	}
}
