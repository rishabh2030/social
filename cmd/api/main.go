package main

import (
	"log"

	"github.com/rishabh2030/social/interanl/env"
)

func main() {
	app := &application{}
	app.LoadEnv()

	cfg := config{
		addr: env.GetString("PORT", ":3000"),
	}
	app.config = cfg

	mux := app.mount()
	app.LoadEnv()

	log.Fatal(app.run(mux))
}
