package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	/*router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})*/

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	//index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	router.POST("/login")
	//login page
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "This is the login page...",
		})
	})

	//signup page
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"content": "This is the signup page...",
		})
	})

	//quote page
	router.GET("/quote", func(c *gin.Context) {
		c.HTML(http.StatusOK, "quote.html", gin.H{
			"content": "This is the quote page...",
		})

	})

	//client page
	router.GET("/client", func(c *gin.Context) {
		c.HTML(http.StatusOK, "client.html", gin.H{
			"content": "This is the client page...",
		})

	})
	router.Run("localhost:8080")

	type signup struct {
		username string `json:"username" binding:"required, min=1, max=20"`
		pass     string `json:"password" binding:"required, min=7, max=20"`
		c_pass   string `json:"confirmpassword" binding:"required, min=7, max=20"`
	}

}
