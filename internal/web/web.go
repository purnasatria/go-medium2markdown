package web

import (
	"bytes"
	"embed"
	"encoding/json"
	"go-medium2markdown/pkg/md2"
	"html/template"
	"io"
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

func Serve(staticFS embed.FS, templateFS embed.FS) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	// Serve static files
	e.StaticFS("/static", echo.MustSubFS(staticFS, "static"))

	// Template
	t := &Template{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"formatAnswer": formatAnswer,
		}).ParseFS(templateFS, "templates/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", handleIndex(staticFS))
	e.POST("/convert", handleConvert)
	// Serve robots.txt
	e.GET("/robots.txt", func(c echo.Context) error {
		robotsTxt := `User-agent: *
   Allow: /

   Sitemap: https://md2.blocka.dev/sitemap.xml`
		return c.String(http.StatusOK, robotsTxt)
	})

	// Serve sitemap.xml
	e.GET("/sitemap.xml", func(c echo.Context) error {
		sitemap := `<?xml version="1.0" encoding="UTF-8"?>
   <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
     <url>
       <loc>https://md2.blocka.dev/</loc>
       <lastmod>2024-10-09</lastmod>
       <changefreq>monthly</changefreq>
       <priority>1.0</priority>
     </url>
   </urlset>`
		c.Response().Header().Set(echo.HeaderContentType, "application/xml")
		return c.String(http.StatusOK, sitemap)
	})

	e.Logger.Info("listening on", port)
	e.Logger.Fatal(e.Start(":" + port))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func handleIndex(staticFS embed.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		faqData, err := staticFS.ReadFile("static/faq.json")
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		var faq FAQ
		err = json.Unmarshal(faqData, &faq)
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"FAQ": faq,
		})
	}
}

func handleConvert(c echo.Context) error {
	var formData FormData
	if err := c.Bind(&formData); err != nil {
		c.Logger().Error(err)
		setToastHeader(c, "Error", "Can't read form data: "+err.Error(), ToastError)
		return err
	}

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
		c.Logger().Error(err)
		setToastHeader(c, "Error", "Failed to convert Medium post: "+err.Error(), ToastError)
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
	setToastHeader(c, "Success", "Download will start automatically...", ToastSuccess)

	return c.Blob(http.StatusOK, c.Response().Header().Get("Content-Type"), buf.Bytes())
}

func formatAnswer(answer string) template.HTML {
	// First, escape the HTML to prevent XSS attacks
	safe := template.HTMLEscapeString(answer)

	// Convert newlines to <br> tags
	safe = strings.ReplaceAll(safe, "\n", "<br>")

	// Return the safe HTML
	return template.HTML(safe)
}

// ToastType represents the type of toast message
type ToastType string

const (
	ToastSuccess ToastType = "success"
	ToastError   ToastType = "error"
	ToastInfo    ToastType = "info"
	ToastWarning ToastType = "warning"
)

// setToastHeader sets the HX-Trigger header for showing a toast message
func setToastHeader(c echo.Context, title, message string, toastType ToastType) {
	toast := struct {
		Title   string    `json:"title"`
		Message string    `json:"message"`
		Type    ToastType `json:"type"`
	}{
		Title:   title,
		Message: message,
		Type:    toastType,
	}

	toastJSON, err := json.Marshal(map[string]interface{}{
		"showToast": toast,
	})
	if err != nil {
		// Handle error (in this case, we'll just not set the header)
		return
	}

	c.Response().Header().Set("HX-Trigger", string(toastJSON))
}
