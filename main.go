package main

import (
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	// "math/rand"
	"html/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "/")
}


func main() {
	e := echo.New()

	t := &Template{
			templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}

	e.Renderer = t
	e.GET("/", Home)
	e.Logger.Fatal(e.Start(":5000"))
}