package main

import (
	"os"

	"github.com/ambientkit/ambient/pkg/grpcp"
	"github.com/ambientkit/plugin/generic/pluginmanager"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: grpcp.Handshake,
		Plugins: map[string]plugin.Plugin{
			"pluginmanager": &grpcp.GenericPlugin{Impl: pluginmanager.New()},
		},
		Logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Output:     os.Stderr,
			JSONFormat: true,
		}),
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
