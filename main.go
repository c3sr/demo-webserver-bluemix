package main

import (
	"fmt"
	"log"
	"os"

	"bitbucket.org/c3sr/p3sr-demo-webserver-bluemix/web"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = DEFAULT_PORT
	}

	log.Printf("Starting app on port %+v\n", port)
	web.Start(fmt.Sprintf(":%s", port))
}
