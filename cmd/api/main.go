package main

import (
	"log"

	"github.com/yudhisrana/go-hospital/internal/interface/http"
)

func main() {
	srv := http.NewServer()

	println("Starting server on port 8080")
	if err := srv.Start("8080"); err != nil {
		log.Fatal(err)
	}
}
