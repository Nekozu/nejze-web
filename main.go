package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

  router = gin.Default()
  router.LoadHTMLGlob("public/*")

  router.GET("/", func(c *gin.Context) {

    c.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
        "title": "Home",
      },
    )

   })
  router.GET("/contact", func(c *gin.Context) {

    c.HTML(
      http.StatusOK,
      "contact.html",
      gin.H{
      },
    )

  })

  // Start serving the application
  router.Run()

}
