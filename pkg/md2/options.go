package md2

type MarkupSymbol struct {
	Italic  string `yaml:"italic"`
	Section string `yaml:"section"`
}

type Options struct {
	MarkupSymbol     MarkupSymbol `yaml:"markupSymbol"`
	IsDownloadAssets bool         `yaml:"isDownloadAssets"`
}

func getDefaultOptions() *Options {
	conf := Options{
		MarkupSymbol{
			Italic:  "*",
			Section: "***",
		},
		false,
	}

	return &conf
}
