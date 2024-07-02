rm -rf proto-go proto-ts proto-js
mkdir -p proto-go proto-ts proto-js/commonjs proto-js/es6 proto-ts/commonjs+dts proto-ts/typescript client server

# # go
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc chat.proto -I=./proto --go_out=proto-go
# protoc chat.proto -I=./proto --go-grpc_out=go
protoc chat.proto -I=./proto --go-grpc_out=require_unimplemented_servers=false:proto-go

cp -r ./proto-go/chat ./client
cp -r ./proto-go/chat ./server

# # javascript
# commonjs
protoc chat.proto -I=./proto \
  --js_out=import_style=commonjs,binary:./proto-js/commonjs --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js
# es6
protoc chat.proto -I=./proto \
  --js_out=import_style=es6,binary:./proto-js/es6 --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js

# # javascript / grpc-web
# commonjs
protoc chat.proto -I=./proto \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./proto-js/commonjs --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64
# --grpc-web_out=import_style=commonjs,mode=grpcweb:./proto-js/commonjs --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64

# # typescript
protoc chat.proto -I=./proto \
  --ts_out=./proto-ts/typescript --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts

# # typescript / grpc-web
# commonjs+dts
protoc chat.proto -I=./proto \
  --js_out=import_style=commonjs,binary:./proto-ts/commonjs+dts --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:./proto-ts/commonjs+dts --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64
# --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:./proto-ts/commonjs+dts --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64
# typescript
protoc chat.proto -I=./proto \
  --js_out=import_style=commonjs,binary:./proto-ts/typescript --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:./proto-ts/typescript --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64
# --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:./proto-ts/commonjs+dts --plugin=protoc-gen-grpc-web=./protoc-gen-grpc-web-1.5.0-linux-x86_64

echo "protoc finished"
