package main

import (
	"example/gyy"
	"net/http"
)

func main()  {
	r := gyy.Default()
	r.GET("/", func(c *gyy.Context) {
		c.String(http.StatusOK, "Hello gyyktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gyy.Context) {
		names := []string{"gyyktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}