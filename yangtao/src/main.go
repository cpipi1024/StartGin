package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler_post(c *gin.Context) {
	c.HTML(http.StatusOK, "post/index.html", gin.H{
		"title": "post index.html",
	})
}

func handler_users(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"title": "users inde.html",
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/**/*")
	r.GET("/post/index", handler_post)
	r.GET("/users/index", handler_users)
	r.Run(":8080")
}
