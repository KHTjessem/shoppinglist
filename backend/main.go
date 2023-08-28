package main

import (
	"github.com/gin-gonic/gin"
)

func startListening(router *gin.Engine) {
	println("Listening on port 8080")
	// Start listening
	router.Run(":8080")
}

// Global datbase access
// done in this demonstration for simplicity.
var DATABASE *Database

func main() {
	// Initialize database
	var err error
	DATABASE, err = NewDatabase()
	if err != nil {
		panic(err)
	}
	DATABASE.ExecuteScript("shoppinglist.sql")

	// endpoints
	router := gin.Default()
	registerPages(router)

	// Start serving
	startListening(router)
}
