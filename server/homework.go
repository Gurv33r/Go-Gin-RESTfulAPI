package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Assignment struct {
	Id     uint64 `json:"id,omitempty"`
	Title  string `json:"title"`
	Prompt string `json:"prompt"`
}

// Assignments = simulates a database
var Assignments = []Assignment{
	{1, "Implement Stack", "Implement a stack data structure in a programming language of your choice"},
	{2, "Implement Queue", "Implement a queue data structure in a programming language of your choice"},
	{3, "Implement Tree", "Implement a tree data structure in a programming language of your choice"},
}

// GetAssignments = accessor for assignments slice -> returns assignments slice as a single JSON string
// Gin framework's cornerstone is gin.Context, which is an abstraction over a datastruct that carries HTTP request data and also converts to and from JSON format
// to convert a struct slice to a JSON string, use gin.Context.IndentedJSON, which takes a http.StatusCode and the Struct slice to convert
// to deny requests, set the status code to http.StatusBadRequest and the struct slice to nil
func GetAssignments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Assignments)
}

// HomePageHandler = accessor for homepage (http://localhost:port/)
// returns "hello, world"
func HomePageHandler(ctx *gin.Context) {
	//sending a string gets handled by gin
	ctx.JSON(http.StatusOK, "Hello, World")
}

// NewAssignment = should only accept a post request
// add the incoming assignment to the Assignments 'db'
func NewAssignment(ctx *gin.Context) {
	// to decode json data, you need to preemptively create an Assignment object to parse the JSON data into
	// like you are catching JSON from someone
	var catcher Assignment        // create a default assignment (will be empty)
	err := ctx.BindJSON(&catcher) // catch it by passing its memory address to the json decoder
	// handle errors, probably means there's a problem in the request
	if err != nil {
		ctx.String(http.StatusBadRequest, "malformed request")
	}
	// id shouldn't be a
	catcher.Id = uint64(len(Assignments) + 1)
	// update 'db'
	Assignments = append(Assignments, catcher)
	ctx.String(http.StatusCreated, "")
}
