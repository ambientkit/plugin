package main

import (
	"log"

	"github.com/ambientkit/plugin/pkg/grpctestutil"
)

func main() {
	app, err := grpctestutil.Setup2(true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	h, err := app.Handler()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Fatalln(app.ListenAndServe(h))
}
