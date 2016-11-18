package main

import (
	"log"
	"os"

	"bitbucket.org/c3sr/p3sr-demo-webserver-bluemix/web"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	log.Printf("Starting app on port %+v\n", port)
	web.Start(":1323")
}
