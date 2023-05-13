package main

import (
	"log"
	"notes"
)

func main() {
	srv := new(notes.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server:%s", err)
	}
}
