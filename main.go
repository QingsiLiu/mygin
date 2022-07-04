package main

import (
	"mygin"
	"net/http"
)

func main() {
	r := mygin.New()
	r.Get("/index", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.Get("/", func(c *mygin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello MyGin</h1>")
		})

		v1.Get("/hello", func(c *mygin.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *mygin.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.Post("/login", func(c *mygin.Context) {
			c.Json(http.StatusOK, mygin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":8001")
}
