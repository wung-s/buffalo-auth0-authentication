package main

import (
	"log"

	"github.com/wung-s/buffalo-auth0-authentication/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
