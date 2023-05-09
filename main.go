package main

import (
	"crud/config"
	"crud/route"
)

// get all users

func main() {
	// create a new echo instance
	config.Init()
	e := route.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
