package main

import (
	"log"

	"github.com/shinya-ml/go-distributional-application/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
