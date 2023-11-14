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

type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

var blogs = []Blog {
	{Title: "Title 1", Description: "Description 1"},
	{Title: "Title 2", Description: "Description 2"},
	{Title: "Title 3", Description: "Description 3"},
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "/")
}

func Feed(c echo.Context) error {
	return c.Render(http.StatusOK, "feed.html", "/")
}

func Create(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", "/")
}

func getBlogs(c echo.Context) error {
	return c.Render(http.StatusOK, "cards.html", blogs)
}


func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	blog := Blog {
		Title: title,
		Description: description,
	}

	blogs = append(blogs, blog)
	return c.Render(http.StatusOK, "cards.html", blogs)
}

func removeBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "cards.html", nil)
}

func updateBlog(c echo.Context) error {
	return c.Render(http.StatusOK, "cards.html", nil)
}

func main() {
	e := echo.New()

	t := &Template{
			templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}

	e.Renderer = t
	e.GET("/", Home)
	e.GET("/feed", Feed)
	e.GET("/create", Create)
	e.GET("/blogs", getBlogs)
	e.POST("/addBlog", addBlog)
	e.Logger.Fatal(e.Start(":5000"))
}