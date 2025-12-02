package main

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/louislef299/greeter-plugin/shared"
)

type Greet struct{}

func (Greet) Greet(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"greet": &shared.GreetPlugin{Impl: &Greet{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
