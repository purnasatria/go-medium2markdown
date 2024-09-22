package main

func (s Sections) process(cfg *Config, processedParagraphs []string) {
	if len(s) > 0 {
		for _, secs := range s {
			if secs.StartIndex > 0 {
				processedParagraphs[secs.StartIndex] = cfg.MarkupSymbol.Section + br + br + processedParagraphs[secs.StartIndex]
			}
		}
	}
}
