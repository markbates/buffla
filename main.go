package main

import (
	"log"

	"github.com/markbates/buffla/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
