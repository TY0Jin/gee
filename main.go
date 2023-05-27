package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	engine := gee.New()
	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./static")

	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	v2 := engine.Group("/v2")
	v2.Use(gee.Logger())
	{
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
		})
	}
	engine.GET("/panic", func(c *gee.Context) {
		names := []string{"jinxw"}
		c.String(http.StatusOK, names[200])
	})
	engine.RUN(":9999")
}
