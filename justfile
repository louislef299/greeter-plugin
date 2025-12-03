API_DIR := "./api"

proto: deepclean
    protoc -I={{API_DIR}} --go_out={{API_DIR}} --go_opt=paths=source_relative \
    --go-grpc_out={{API_DIR}} --go-grpc_opt=paths=source_relative \
     {{API_DIR}}/greeter.proto

run NAME: clean
    go build -o greet ./plugins
    time go run main.go {{NAME}}

clean:
    @rm greet greet-plugin-called || true

deepclean: clean
    rm api/*go greet || true
