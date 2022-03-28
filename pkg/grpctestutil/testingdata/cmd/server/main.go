package main

import (
	"log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/grpctestutil"
)

func main() {
	app, err := grpctestutil.StandardSetup(true)
	//app, err := grpctestutil.GRPCSetup(true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	app.PluginSystem().SetGrant("hello", ambient.GrantPluginNeighborGrantRead)

	h, err := app.Handler()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Fatalln(app.ListenAndServe(h))
}
