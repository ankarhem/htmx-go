package main

import (
	"fmt"
	"htmx/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Println(
		"Starting server on",
		server.Addr,
	)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
