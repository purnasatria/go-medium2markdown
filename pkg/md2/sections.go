package md2

func (s Sections) Parse(processedParagraphs []string, opt Options) {
	if len(s) > 0 {
		for _, secs := range s {
			if secs.StartIndex > 0 {
				processedParagraphs[secs.StartIndex] = opt.MarkupSymbol.Section + newLine + newLine + processedParagraphs[secs.StartIndex]
			}
		}
	}
}
