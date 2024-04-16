package main

import (
	"html/template"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func (app *application) parseFileIntoPost(p *Post, file string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	c := string(contents)

	// Extract front matter
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

	p.Title = metaMap["title"]
	p.Slug = metaMap["slug"]
	p.Catergory = metaMap["catergory"]

	html := mdToHTML([]byte(mu))
	p.Content = template.HTML(html)

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
