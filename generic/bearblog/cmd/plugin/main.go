package main

import (
	"context"
	"log"
	"os"

	"github.com/ambientkit/ambient/pkg/grpcp"
	"github.com/ambientkit/ambient/pkg/hclogadapter"
	"github.com/ambientkit/plugin/generic/bearblog"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/hashicorp/go-plugin"
)

func main() {
	p := bearblog.New(os.Getenv("AMB_PASSWORD_HASH"))
	ctx := context.Background()
	pluginName := p.PluginName(ctx)

	zlog, err := zaplogger.New().Logger(ctx, pluginName, p.PluginVersion(ctx), nil)
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
