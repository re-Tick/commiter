package main

import (
	"commiter/cmd"
	"log"
)

func main() {

	// run the committer CLI
	err := cmd.Execute()
	if err != nil {
		log.Panic("failed to start the CLI.", err)
	}
}
