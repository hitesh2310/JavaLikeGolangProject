package main

import (
	"fmt"
	"main/config"
	"main/pkg/router"
)

func init() {
	fmt.Println("Starting ECom App....")
	config.SetUpApplication()
}

func main() {

	server := router.NewServer()
	server.ListenAndServe()
}
