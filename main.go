package main

import (
	"log"
	"mygin"
	"net/http"
	"time"
)

func main() {
	r := mygin.New()
	r.Use(mygin.Logger())
	r.Get("/index", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	// 添加一个中间件
	v1.Use(func(c *mygin.Context) {
		log.Printf("[%d] %s in %v (v1 middleware)", c.StatusCode, c.Req.URL, time.Since(time.Now()))
	})
	{
		v1.Get("/hello/:name", func(c *mygin.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8001")
}
