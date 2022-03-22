package main

// func main() {
// 	plugin.Serve(&plugin.ServeConfig{
// 		HandshakeConfig: grpcp.Handshake,
// 		Plugins: map[string]plugin.Plugin{
// 			"bearblog": &grpcp.GenericPlugin{Impl: bearblog.New()},
// 		},
// 		Logger: hclog.New(&hclog.LoggerOptions{
// 			Level:      hclog.Debug,
// 			Output:     os.Stderr,
// 			JSONFormat: true,
// 		}),
// 		GRPCServer: plugin.DefaultGRPCServer,
// 	})
// }
