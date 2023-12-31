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
// alternative to https://gin-gonic.com/docs/examples/bind-single-binary-with-template/
func get_templates() multitemplate.Renderer {
	base_tmpl, _ := embedFS.ReadFile("templates/base.tmpl")
	base_tmpl_str := string(base_tmpl)

	about_tmpl, _ := embedFS.ReadFile("templates/about.tmpl")
	about_tmpl_str := string(about_tmpl)

	files_tmpl, _ := embedFS.ReadFile("templates/files.tmpl")
	files_tmpl_str := string(files_tmpl)

	mtr := multitemplate.NewRenderer()
	//mtr.AddFromString("index", string(base_tmpl)) // XXX

	mtr.AddFromStringsFuncs("about", template.FuncMap{}, base_tmpl_str, about_tmpl_str)
	mtr.AddFromStringsFuncs("files", template.FuncMap{}, base_tmpl_str, files_tmpl_str)

	return mtr
}
