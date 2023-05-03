package main

import (
	"log"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
