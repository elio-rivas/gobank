package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", store)

	server := NewAPIServer(":3000", store)
	server.Run()
}

//time 13:23-45.41
//How to build a complete json api in Golang(JWT, Postgres, and Docker) part 2
