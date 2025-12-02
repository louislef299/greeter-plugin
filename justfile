API_DIR := "./api"

proto: clean
    protoc -I={{API_DIR}} --go_out={{API_DIR}} --go_opt=paths=source_relative \
    --go-grpc_out={{API_DIR}} --go-grpc_opt=paths=source_relative \
     {{API_DIR}}/greeter.proto

run NAME: proto
    go build -o greet ./plugins
    go run main.go {{NAME}}

clean:
    rm api/*go greet || true
