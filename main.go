package main

import (
	"go-gin-sample/server"
)

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
