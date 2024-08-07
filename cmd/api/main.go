package main

import (
	"fmt"

	"portfolio/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
