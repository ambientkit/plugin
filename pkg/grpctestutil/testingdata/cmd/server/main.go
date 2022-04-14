package main

import (
	"context"
	stdlog "log"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/grpctestutil"
)

func main() {
	//app, log, err := grpctestutil.StandardSetup(true)
	app, log, err := grpctestutil.GRPCSetup(true)
	if err != nil {
		stdlog.Fatalln(err.Error())
	}

	app.PluginSystem().SetGrant("hello", ambient.GrantPluginNeighborGrantRead)

	h, err := app.Handler(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	app.SetLogLevel(ambient.LogLevelDebug)

	log.Fatal("%v", app.ListenAndServe(h))
}
