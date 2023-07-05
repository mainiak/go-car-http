package internal

import (
	"embed"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

// see https://pkg.go.dev/embed
//
//go:embed templates/*
var embedFS embed.FS

// https://gin-gonic.com/docs/examples/html-rendering/#custom-template-renderer
func get_templates() multitemplate.Renderer {
	base_tmpl, _ := embedFS.ReadFile("templates/base.tmpl")
	base_tmpl_str := string(base_tmpl)

	index_tmpl, _ := embedFS.ReadFile("templates/index.tmpl")
	index_tmpl_str := string(index_tmpl)

	files_tmpl, _ := embedFS.ReadFile("templates/files.tmpl")
	files_tmpl_str := string(files_tmpl)

	mtr := multitemplate.NewRenderer()
	//mtr.AddFromString("index", string(base_tmpl)) // XXX

	mtr.AddFromStringsFuncs("index", template.FuncMap{}, base_tmpl_str, index_tmpl_str)
	mtr.AddFromStringsFuncs("files", template.FuncMap{}, base_tmpl_str, files_tmpl_str)

	return mtr
}
