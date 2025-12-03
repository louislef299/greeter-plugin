# greeter-plugin

```bash
$ just run louis
go build -o greet ./plugins
time go run main.go louis
2025-12-02T19:01:12.808-0600 [DEBUG] plugin: starting plugin: path=/bin/sh args=["sh", "-c", "./greet"]
2025-12-02T19:01:12.810-0600 [DEBUG] plugin: plugin started: path=/bin/sh pid=3590
2025-12-02T19:01:12.810-0600 [DEBUG] plugin: waiting for RPC address: plugin=/bin/sh
2025-12-02T19:01:13.072-0600 [DEBUG] plugin: using plugin: version=1
2025-12-02T19:01:13.073-0600 [DEBUG] plugin.sh: plugin address: address=/var/folders/gj/2j8kvlpn0wv4y9ytwcwppczr0000gp/T/plugin2574164733 network=unix timestamp=2025-12-02T19:01:13.072-0600
2025/12/02 19:01:13 rpcClient type: *plugin.GRPCClient
2025-12-02T19:01:13.074-0600 [TRACE] plugin.stdio: waiting for stdio data
2025/12/02 19:01:13 successfully pinged the rpcClient
2025/12/02 19:01:13 response from plugin: hello louis
2025-12-02T19:01:13.076-0600 [DEBUG] plugin.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
2025-12-02T19:01:13.076-0600 [INFO]  plugin: plugin process exited: plugin=/bin/sh id=3590
2025-12-02T19:01:13.076-0600 [DEBUG] plugin: plugin exited

real    0m0.729s
user    0m0.161s
sys     0m0.592s
```
