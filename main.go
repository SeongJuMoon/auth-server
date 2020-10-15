package main

import (
	"log"

	"github.com/seongjumoon/authServer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
