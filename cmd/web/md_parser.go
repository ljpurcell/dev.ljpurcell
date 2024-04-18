package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdHtml "github.com/gomarkdown/markdown/html"
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
	htmlFlags := mdHtml.CommonFlags | mdHtml.HrefTargetBlank
	opts := mdHtml.RendererOptions{Flags: htmlFlags, RenderNodeHook: myRenderHook}
	render := mdHtml.NewRenderer(opts)

	return markdown.Render(doc, render)
}

/**
 * Syntax highlighting
 */
var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

func init() {
	htmlFormatter = html.New(html.WithClasses(false), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}

	// Options:
	// RosePint
	// CatppuccinMocha
	// GitHibDark
	styleName := styles.GitHubDark.Name
	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := "go"
	lang := string(codeBlock.Info)
	applyHighlighting(w, string(codeBlock.Literal), lang, defaultLang)
}

func applyHighlighting(w io.Writer, source, lang, defaultLang string) error {
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
	return htmlFormatter.Format(w, highlightStyle, it)
}
