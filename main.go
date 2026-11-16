package main

import (
	"gowebuiystu/app/apiserver"
	"log"
)

var config = apiserver.Config{
	Addr: ":8080",
}

func main() {
	err := apiserver.StartServer(config)
	if err != nil {
		log.Fatal(err)
	}
}
