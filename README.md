# README

## Compile stubs

Will place the stubs next to the *.proto files.

```{shell}
protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative message/message.proto 
```
## Running the examples

From the project directory start the server.

```{shell}
go run ./server/main.go
```

From another terminal run the client.

```{shell}
go run ./client/main.go
```
