package main

import (
	"errors"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdHtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func (app *application) parseFileIntoPost(p *Post, file string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrPostNotFound
		}

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
			field := strings.ToLower(strings.Trim(fieldVal[0], " \""))
			value := strings.Trim(fieldVal[1], " \"")
			metaMap[field] = value
		}
	}

	p.Title = metaMap["title"]
	p.Slug = metaMap["slug"]
	p.Category = metaMap["catergory"]

	// Build tags slice
	tagStrings := strings.Split(metaMap["tags"], ",")
	p.Tags = make([]Tag, len(tagStrings))

	for i, tag := range tagStrings {
		p.Tags[i] = Tag(strings.Trim(tag, " \""))
	}

	html := app.mdToHTML([]byte(mu))
	p.Content = template.HTML(html)

	return nil
}

func (app *application) mdToHTML(md []byte) []byte {
	// Create MD parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.SuperSubscript | parser.Attributes
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// Create HTML render with extensions
	htmlFlags := mdHtml.CommonFlags | mdHtml.HrefTargetBlank
	opts := mdHtml.RendererOptions{Flags: htmlFlags, RenderNodeHook: app.myRenderHook}
	render := mdHtml.NewRenderer(opts)

	return markdown.Render(doc, render)
}

/**
 * Syntax highlighting
 */
func (app *application) myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.Code); ok {
		lang := string("go")
		applyHighlighting(w, app.htmlInlineFormatter, app.highlightStyle, string(code.Literal), lang, app.defaultLang)
		return ast.GoToNext, true
	}

	if codeBlock, ok := node.(*ast.CodeBlock); ok {
		lang := string(codeBlock.Info)
		applyHighlighting(w, app.htmlBlockFormatter, app.highlightStyle, string(codeBlock.Literal), lang, app.defaultLang)
		return ast.GoToNext, true
	}

	return ast.GoToNext, false
}

func applyHighlighting(w io.Writer, formatter chroma.Formatter, highlightStyle *chroma.Style, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return formatter.Format(w, highlightStyle, it)
}
