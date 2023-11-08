package main

import (
	"flag"
	"fmt"
	"goco2"
	"log"
	"os"
)

func main() {
	// read CLI arguments
	var host string
	flag.StringVar(&host, "host", os.Getenv("HTTP_HOST"), "host the microservice listen to")
	var port string
	flag.StringVar(&port, "port", os.Getenv("HTTP_PORT"), "port the microservice listen to")
	flag.Parse()

	// initialize the service
	svc := goco2.NewCO2Service(600)
	svc = goco2.NewLoggingService(svc)

	// initialize the api server
	apiServer := goco2.NewAPIServer(svc)
	// start the api server
	log.Printf("goco2 microservice started listening on %s:%s", host, port)
	if err := apiServer.Start(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatalf("API server failed to start, error: %s", err.Error())
	}
}
