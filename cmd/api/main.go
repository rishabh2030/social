package main

import (
	"log"

	"github.com/rishabh2030/social/interanl/db"
	"github.com/rishabh2030/social/interanl/env"
	"github.com/rishabh2030/social/interanl/store"
)

func main() {
	app := &application{}
	app.LoadEnv()

	cfg := config{
		addr: env.GetString("PORT", ":3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		log.Panic("Failed to connect to the database:", err)
	}
	defer db.Close()

	log.Println("Database connection established")

	store := store.NewPostgresStorage(db)

	app.config = cfg
	app.store = store

	mux := app.mount()

	log.Fatal(app.run(mux))
}
