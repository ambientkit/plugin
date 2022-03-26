package main

import (
	"log"

	"github.com/ambientkit/ambient/pkg/grpcp"
	"github.com/ambientkit/ambient/pkg/hclogadapter"
	"github.com/ambientkit/plugin/generic/pluginmanager"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/hashicorp/go-plugin"
)

func main() {
	p := pluginmanager.New()

	zlog, err := zaplogger.New().Logger(p.PluginName(), p.PluginVersion(), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: grpcp.Handshake,
		Plugins: map[string]plugin.Plugin{
			p.PluginName(): &grpcp.GenericPlugin{Impl: p},
		},
		Logger:     hclogadapter.New(p.PluginName(), zlog),
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
