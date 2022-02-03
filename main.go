package main

import (
	"github.com/Gurv33r/go-env"
	"log"
	"os"

	//external or not built-in
	. "github.com/Gurv33r/Go-Gin-RESTfulAPI/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// for linux/mac: use curl http://localhost:port/dir
	// for windows: use Invoke-WebRequest -URI http://localhost:port/dir
	// web browsers will work too, just go to http://localhost:port/dir
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}
	// Basics of Go Gin:
	// all APIs, especially web APIs need a router
	// instantiate one with gin.Default()
	router := gin.Default()
	// A Gin router is special, where you can specify handlers to HTTP requests to a specfic subdirectory
	// this is useful because Gin abstracts over the routing issues for you and requires you to code the bare essentials of a RESTful API: create a handler and route a request to a handler
	// One way is to make immediate and specific routes to handler for each request with the Gin.Router's GET(), POST(), and PUT() methods
	// for now since the code only transfers data, set getAssignments to GET request handler for the /hw subdirectory
	router.GET("/hw", GetAssignments)
	// so far, all this is is letting someone else access an API endpoint for the hw subdirectory
	// let's add a home endpoint by adding a home page handler and then linking it to the '/' sub directory for GET requests
	router.GET("/", HomePageHandler)
	// you can also specify other http verbs
	router.POST("/new", NewAssignment)
	// now, just http://localhost:port will call the home page handler

	// start the router with router.Run("localhost:port")
	// grab the port number from the env file
	// os package has Getenv method that works well
	log.Fatal(router.Run("localhost:" + os.Getenv("PORT")))
}
