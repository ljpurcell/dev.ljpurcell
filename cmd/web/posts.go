package main

import (
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
