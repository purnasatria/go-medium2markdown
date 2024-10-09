package main

import (
	"embed"
	"go-medium2markdown/internal/web"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	// cli.Exceute()
	web.Serve(staticFS, templateFS)
}
