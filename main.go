package main

import (
	"probable-potato-go/main.go/server"
)

func main() {
	r := server.SetupServer()
	r.Run(":8080")
}
