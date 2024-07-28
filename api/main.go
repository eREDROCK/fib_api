package main

import (
	"api/router"
	"fmt"
	"log"
)

func main() {
	e := router.NewRouter()
	fmt.Println("Server started at http://localhost:8888")
	if err := e.Start(":8888"); err != nil {
		log.Fatalf("Error Starting Fibserver: %v", err)
	}
}
