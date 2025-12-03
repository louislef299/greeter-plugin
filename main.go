package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/louislef299/greeter-plugin/shared"
)

func main() {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", "./greet"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via gRPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("rpcClient type: %T\n", rpcClient)

	// Ping the client
	err = rpcClient.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully pinged the rpcClient")

	// Request the plugin
	raw, err := rpcClient.Dispense("greet")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	g := raw.(shared.Greeter)
	os.Args = os.Args[1:]
	log.Println("response from plugin:", g.Greet(os.Args[0]))
}
