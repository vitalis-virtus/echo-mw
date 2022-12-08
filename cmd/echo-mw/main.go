package main

import (
	"log"

	"github.com/vitalis-virtus/echo-mw/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
