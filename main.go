package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)
var router *gin.Engine

type Page struct{
	Title string
	Body []byte
}
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page,error){
	filename := title + ".txt"
	body,err := ioutil.ReadFile(filename)
	fmt.Println(body)
	if err !=nil{
		log.Fatal(err)
		return nil,err 
	}
	return &Page{Title : title, Body:body},nil
}

func handler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"Hi there,I love %s",r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter , r *http.Request){
	vars:=mux.Vars(r)
	title := vars["fileName"]
	p,_ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
	  // Respond with JSON
	  c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
	  // Respond with XML
	  c.XML(http.StatusOK, data["payload"])
	default:
	  // Respond with HTML
	  c.HTML(http.StatusOK, templateName, data)
	}
  
  }


func main() {
    //p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    //p1.save()
    //p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", handler)
	//router.HandleFunc("/view/{fileName}",viewHandler)
	//log.Fatal(http.ListenAndServe(":8000", router))
	
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	// router.GET("/", func(c *gin.Context) {

	// 	// Call the HTML method of the Context to render a template
	// 	c.HTML(
	// 		// Set the HTTP status to 200 (OK)
	// 		http.StatusOK,
	// 		// Use the index.html template
	// 		"index.html",
	// 		// Pass the data that the page uses (in this case, 'title')
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	  
	//   })
	router.Run(":2020")

}