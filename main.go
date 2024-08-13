package main

import (
	"flag"
	"log"

	"github.com/darcoprogramador/go-chat/app"
	"github.com/darcoprogramador/go-chat/app/controllers"
)

func main() {
	hub := controllers.NewHub()

	go hub.RunHub()

	app := app.NewApplication(hub)

	// app.Static("/", "./static/index.html")

	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
