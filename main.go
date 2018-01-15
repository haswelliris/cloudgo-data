package main

import (
	"os"

	"github.com/haswelliris/cloudgo-data/service"
	flag "github.com/spf13/pflag"
)

const (
	// PORT : default listening port if os doesn't have an env $PORT & -p not specified
	PORT string = "9090"
)

func main() {
	// first get $PORT
	port := os.Getenv("PORT")
	// if os doesn't have an env $PORT, set to const PORT
	if len(port) == 0 {
		port = PORT
	}

	// if specified -p, set listening port to the value of -p
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// start server
	server := service.NewServer()
	server.Run(":" + port)
}
