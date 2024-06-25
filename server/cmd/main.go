package main

import (
	"log"
	"server/db"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(wsHandler)
	router.Start("0.0.0.0:8080")

}
