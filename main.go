package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sebasromero/simple_rest_api/server"
)

func main() {
	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)                    //Create the channel to be able to communicate after its inicialization
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM) //Sending a signal through the channel

	//Define the port number creating the server
	server := server.New(":8080")

	//Running the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Server started")

	<-serverDoneChan     //Receiving the signal
	server.Shutdown(ctx) //Closing the server
	log.Println("Server stopped")
}
