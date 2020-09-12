package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()
  // Call the HTML method of the Context to render a template
  render(c, gin.H{
    "title":   "Home Page",
    "payload": articles,    
  },"index.html")
  // c.HTML(
  //   // Set the HTTP status to 200 (OK)
  //   http.StatusOK,
  //   // Use the index.html template
  //   "index.html",
  //   // Pass the data that the page uses
  //   gin.H{
  //     "title":   "Home Page",
  //     "payload": articles,
  //   },
  // )

}

func getArticle (c *gin.Context){
  if articleId,err := strconv.Atoi(c.Param("article_id"));err==nil{
    if article,err := getArticleByID(articleId);err==nil{
      c.HTML(
        http.StatusOK,
        "article.html",
        gin.H{
          "title" : article.Title,
          "payload" : article,
        },
      )
    } else{
      c.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    c.AbortWithStatus(http.StatusNotFound)
  }
}