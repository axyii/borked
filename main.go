package main

import (
	"github.com/axyii/borked/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
)

type Post struct {
	Title   string
	Content template.HTML
}

func main() {
	utils.Genpages()
	allposts := utils.Getchrononames()
	path, err := os.Getwd()
	if err != nil {
		utils.Logger.Println(err)
	}
	r := gin.Default()
	r.Use(utils.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Delims("{{", "}}")
	r.SetFuncMap(template.FuncMap{
		"formatname": utils.Formatasname,
		"formatdate": utils.Formatasdate,
	})
	r.Use(static.Serve("/assets", static.LocalFile(path+"/assets", false)))
	r.LoadHTMLGlob(path + "/templates/*.tmpl.html")

	r.GET("/", func(c *gin.Context) {
		var posts []string
		posts = allposts[:2]

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"posts": posts,
		})
	})

	r.GET("/articles/:postName", func(c *gin.Context) {
		postName := c.Param("postName")
		htmlfile, err := os.ReadFile(path + "/articles/" + postName)
		if err != nil {
			utils.Logger.Println(err)
			c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
			return
		}
		postHTML := template.HTML(htmlfile)
		postit := utils.Formatasname(postName)

		post := Post{Title: postit, Content: postHTML}

		c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
			"Title":   post.Title,
			"Content": post.Content,
		})
	})
	r.GET("/articles", func(c *gin.Context) {

		c.HTML(http.StatusOK, "articles.tmpl.html", gin.H{
			"posts": allposts,
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
		return

	})

	r.Run(":8080")
}
