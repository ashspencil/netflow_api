package main

import (
	"log"
	"nmg/netflow/router"
)

func main() {
	log.Println("=== Job start ===")

	r := router.FlowRouter()
	r.Run(":8080")

	if r == nil {
		log.Fatal("[API_DEBUG]: Failed To Create Listening Session")
	}

	log.Println("=== Job finish ===")
}
