package main

import (
	"html/template"
	"net/url"
	"os"
	"path/filepath"
)

func (app *application) newPostCache() (map[string]*Post, error) {
	cache := make(map[string]*Post)

	const POSTS_DIR = "./data/posts/markdown"

	files, err := os.ReadDir(POSTS_DIR)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f := filepath.Join(POSTS_DIR, file.Name())
		p := &Post{}
		app.parseFileIntoPost(p, f)
		cache[p.Slug] = p
	}

	return cache, nil
}

func newTagCache(posts map[string]*Post) map[Tag][]Post {
	cache := map[Tag][]Post{}

	for _, post := range posts {
		for _, t := range post.Tags {
			tag := Tag(url.QueryEscape(string(t)))
			tagPosts, ok := cache[tag]
			if ok {
				cache[tag] = append(tagPosts, *post)
			} else {
				cache[tag] = []Post{*post}
			}
		}
	}

	return cache
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
