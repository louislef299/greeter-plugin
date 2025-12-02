# greeter-plugin

```bash
$ just run louis
rm api/*go greet || true
protoc -I=./api --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative ./api/greeter.proto
go build -o greet ./plugins
go run main.go louis
2025-12-02T17:10:39.652-0600 [DEBUG] plugin: starting plugin: path=/bin/sh args=["sh", "-c", "./greet"]
2025-12-02T17:10:39.653-0600 [DEBUG] plugin: plugin started: path=/bin/sh pid=91692
2025-12-02T17:10:39.653-0600 [DEBUG] plugin: waiting for RPC address: plugin=/bin/sh
2025-12-02T17:10:40.139-0600 [DEBUG] plugin: using plugin: version=1
2025-12-02T17:10:40.139-0600 [DEBUG] plugin.sh: plugin address: address=/var/folders/gj/2j8kvlpn0wv4y9ytwcwppczr0000gp/T/plugin1611998675 network=unix timestamp=2025-12-02T17:10:40.139-0600
2025/12/02 17:10:40 *plugin.GRPCClient
hello, louis!
2025-12-02T17:10:40.140-0600 [TRACE] plugin.stdio: waiting for stdio data
2025-12-02T17:10:40.141-0600 [DEBUG] plugin.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
2025-12-02T17:10:40.141-0600 [INFO]  plugin: plugin process exited: plugin=/bin/sh id=91692
2025-12-02T17:10:40.141-0600 [DEBUG] plugin: plugin exited
```
