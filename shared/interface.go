package shared

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/louislef299/greeter-plugin/api"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"greet": &GreetPlugin{},
}

// Greet is the interface that we're exposing as a plugin
type Greet interface {
	Greet(name string) string
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
type GreetPlugin struct {
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Greet
}

func (p *GreetPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	api.RegisterGreeterServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GreetPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: api.NewGreeterClient(c)}, nil
}

// GRPCClient is an implementation of Greet that talks over RPC.
type GRPCClient struct {
	client api.GreeterClient
}

func (c *GRPCClient) Greet(name string) string {
	return fmt.Sprintf("hello, %s!", name)
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	api.UnimplementedGreeterServer
	// This is the real implementation
	Impl Greet
}

func (s *GRPCServer) Greet(ctx context.Context, p *api.Person) (*api.Greeting, error) {
	return &api.Greeting{
		Message: s.Impl.Greet(p.Name),
	}, nil
}
