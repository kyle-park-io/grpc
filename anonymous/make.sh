rm -rf proto-go
mkdir -p proto-go client server

protoc anonymouschat.proto -I=./proto --go_out=proto-go
protoc anonymouschat.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:proto-go

cp -r ./proto-go/chat ./client
cp -r ./proto-go/chat ./server

echo "protoc finished"
