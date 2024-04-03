package main

import (
	"fmt"
	"oAuth/internal/auth"
	"oAuth/internal/server"
)

func main() {

	auth.NewAuth()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}

type BasicData struct {
	Title       string
	Description string
}

type AdvancedDate struct {
	Title       string
	Description string
	IsExtended  bool
}
