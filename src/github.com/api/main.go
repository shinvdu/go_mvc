package main

import (
	"github.com/api/app"
	"github.com/api/config"
	// "./app"
	// "./config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

// GOPATH=/home/silas/Downloads/temp/  go build main.go

// https://www.golangprograms.com/advance-programs/golang-restful-api-using-grom-and-gorilla-mux.html