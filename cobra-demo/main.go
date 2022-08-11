package main

import (
	"log"

	"github.com/go-demos/cobra-demo/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
