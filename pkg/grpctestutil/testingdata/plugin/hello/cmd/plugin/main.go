package main

import (
	"os"

	"github.com/ambientkit/ambient/pkg/grpcp"
	"github.com/ambientkit/plugin/pkg/grpctestutil/testingdata/plugin/hello"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: grpcp.Handshake,
		Plugins: map[string]plugin.Plugin{
			"hello": &grpcp.GenericPlugin{Impl: hello.New()},
		},
		Logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Debug,
			Output:     os.Stderr,
			JSONFormat: true,
		}),
		//GRPCServer: plugin.DefaultGRPCServer,
		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
			opts = append(opts, grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
			opts = append(opts, grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()))
			return grpc.NewServer(opts...)
		},
	})
}
