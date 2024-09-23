package core

func (s Sections) Parse(opt *MediumConverterOptions, processedParagraphs []string) {
	if len(s) > 0 {
		for _, secs := range s {
			if secs.StartIndex > 0 {
				processedParagraphs[secs.StartIndex] = opt.MarkupSymbol.Section + newLine + newLine + processedParagraphs[secs.StartIndex]
			}
		}
	}
}
