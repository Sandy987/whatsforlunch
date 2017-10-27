package main

import (
	"log"
	"os"
)

func main() {
	app := App{}

	log.Println("Initialising App")
	app.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	log.Printf("Starting API on %s\n", port)
	app.StartAPI(port)
}
