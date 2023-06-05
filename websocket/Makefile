test:
	echo "Hello"


protoc:
	mkdir -p go client server
	protoc ex.proto -I=./proto --go_out=go  
	protoc ex.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:go
	cp -r ./go/websocket ./client
	cp -r ./go/websocket ./server

	echo "protoc finished"