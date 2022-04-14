package main

import (
	"context"
	"log"

	"github.com/ambientkit/ambient/pkg/grpcp"
	"github.com/ambientkit/ambient/pkg/hclogadapter"
	"github.com/ambientkit/plugin/generic/bearcss"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/hashicorp/go-plugin"
)

func main() {
	p := bearcss.New()
	ctx := context.Background()
	pluginName := p.PluginName(ctx)

	zlog, err := zaplogger.New().Logger(pluginName, p.PluginVersion(), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: grpcp.Handshake,
		Plugins: map[string]plugin.Plugin{
			pluginName: &grpcp.GenericPlugin{Impl: p},
		},
		Logger:     hclogadapter.New(pluginName, zlog),
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
