package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-medium2markdown/pkg/md2"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	URL              string `form:"url"`
	ItalicSymbol     string `form:"italicSymbol"`
	SectionSymbol    string `form:"sectionSymbol"`
	IsDownloadAssets bool   `form:"isDownloadAssets"`
}

type FAQ struct {
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func Serve() {
	e := echo.New()

	// Serve static files
	e.Static("/static", "static")

	// Template
	t := &Template{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"formatAnswer": formatAnswer,
		}).ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", handleIndex)
	e.POST("/convert", handleConvert)

	e.Logger.Fatal(e.Start(":8080"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func handleIndex(c echo.Context) error {
	faqData, err := os.ReadFile("./faq.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var faq FAQ
	err = json.Unmarshal(faqData, &faq)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"FAQ": faq,
	})
}

func handleConvert(c echo.Context) error {
	var formData FormData
	if err := c.Bind(&formData); err != nil {
		return err
	}

	log.Print(formData)

	buf := new(bytes.Buffer)
	mco := md2.Options{
		MarkupSymbol: md2.MarkupSymbol{
			Italic:  formData.ItalicSymbol,
			Section: formData.SectionSymbol,
		},
		IsDownloadAssets: formData.IsDownloadAssets,
	}

	mc := md2.NewConverter(buf, mco)
	err := mc.Convert(formData.URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error converting: "+err.Error())
	}

	filename := mc.Metadata.Slug
	if filename == "" {
		filename = "converted"
	}

	if mco.IsDownloadAssets {
		c.Response().Header().Set(echo.HeaderContentType, "application/zip")
		c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+mc.Metadata.Slug+".zip")
	} else {
		c.Response().Header().Set(echo.HeaderContentType, "text/markdown")
		c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+mc.Metadata.Slug+".md")
	}

	c.Response().Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	return c.Blob(http.StatusOK, c.Response().Header().Get("Content-Type"), buf.Bytes())
}

func formatAnswer(answer string) template.HTML {
	// First, escape the HTML to prevent XSS attacks
	safe := template.HTMLEscapeString(answer)

	// Convert newlines to <br> tags
	safe = strings.Replace(safe, "\n", "<br>", -1)

	// Return the safe HTML
	return template.HTML(safe)
}
