package main

import (
	"gitlab/vote-crossbow/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()
	errRouter := http.ListenAndServe(":8082", r)
	if errRouter != nil {
		log.Fatal("Not connect server")
	}
	//r.Run(":8080")
}
