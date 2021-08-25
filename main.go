package main

import (
	"github.com/jairoteran6/twittor/bd"
	"github.com/jairoteran6/twittor/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a la BD!!s")
		return
	}
	handlers.Manejadores()
}
