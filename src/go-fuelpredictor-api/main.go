package main

import (
	"fmt"
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
		username string `form:"username" binding:"required, min=1, max=20"`
		pass     string `form:"password" binding:"required, min=7, max=20"`
		c_Pass   string `form:"confirmpassword" binding:"required, min=7, max=20"`
	}

	router.POST("/signup", func(c *gin.Context) {
		uname := c.PostForm("username")
		pass := c.PostForm("password")
		cPass := c.PostForm("confirmpassword")
		/*var json signup
		c.Bind(&json)
		if json.username == "abc" && json.pass == "abc123" && json.c_Pass == json.pass {
			c.JSON(http.StatusOK, gin.H{"status": "authorized"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "un-authorized"})
		}*/
		fmt.Println(uname, pass, cPass)
	})
}
