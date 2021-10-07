package main

import (
    "fmt"
    "html/template"
    "github.com/gamingbeast36/borked/utils"
    "io/ioutil"
//    "log"
    "net/http"
    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    "github.com/russross/blackfriday"
)

type Post struct {
    Title   string
    Content template.HTML
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.Use(gin.Logger())
    r.Delims("{{", "}}")
    r.SetFuncMap(template.FuncMap{
        "formatname": utils.Formatasname,
        "formatdate": utils.Formatasdate,
    })
    r.Use(static.Serve("/assets", static.LocalFile("./assets", false)))
    r.LoadHTMLGlob("./templates/*.tmpl.html")

    r.GET("/", func(c *gin.Context) {
        var posts []string
        posts = utils.Getchrononames()
        posts = posts[:2]

        c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
            "posts": posts,
        })
    })

    r.GET("/articles/:postName", func(c *gin.Context) {
        postName := c.Param("postName")
        mdfile, err := ioutil.ReadFile("./markdown/" + postName)
        fmt.Println(postName)
        // if the file can not be found
        if err != nil {
            fmt.Println(err)
            c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
            return
        }
        postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))
        postit := utils.Formatasname(postName) 

        post := Post{Title: postit, Content: postHTML}

        c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
            "Title":   post.Title,
            "Content": post.Content,
        })
    })
    r.GET("/articles", func(c *gin.Context){

        posts := utils.Getchrononames()
        c.HTML(http.StatusOK, "articles.tmpl.html", gin.H{
            "posts": posts,
        })
    })
    r.NoRoute(func(c *gin.Context) {
        c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
        return

    })

    r.Run(":8080")
}
