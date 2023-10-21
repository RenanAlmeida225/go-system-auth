package main

import (
	"github.com/RenanAlmeida225/go-system-auth/config"
	"github.com/RenanAlmeida225/go-system-auth/router"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	router.Initialize()
}
