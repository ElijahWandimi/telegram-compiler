package src

import (
	"log"
	"net/http"
	"os"
)



func CreateServer() error{

	var serverPort = os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	// create engine instance
	engine, engineError := NewEngine()
	if engineError != nil {
		return engineError
	}

	// create http server
	http.HandleFunc("/", engine.HandleTelegramWebHook)

	log.Printf("Server is listening on port %s\n", serverPort)
	log.Printf("Press CTRL+C to stop the server\n")
	

	// start http server
	return http.ListenAndServe(":" + serverPort, nil)
}