package main

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func parseMdFile(content []byte, p *post, md *[]byte) error {
	c := string(content)
	seperators := 2
	if !strings.HasPrefix(c, "---") {
		seperators = 1
	}
	sections := strings.SplitN(c, "---", seperators+1)

	metaBlock, mu := sections[len(sections)-2], sections[len(sections)-1]

	metaLines := strings.Split(metaBlock, "\n")

	metaMap := make(map[string]string)
	for _, line := range metaLines {
		fieldVal := strings.SplitN(line, ":", 2)
		if len(fieldVal) == 2 {
			field := strings.ToLower(strings.TrimSpace(fieldVal[0]))
			value := strings.TrimSpace(fieldVal[1])
			metaMap[field] = value
		}
	}

	p.title = metaMap["title"]
	p.slug = metaMap["slug"]
	p.catergory = metaMap["catergory"]

	*md = mdToHTML([]byte(mu))
	return nil
}

func mdToHTML(md []byte) []byte {
	// Create MD parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// Create HTML render with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	render := html.NewRenderer(opts)

	return markdown.Render(doc, render)
}
