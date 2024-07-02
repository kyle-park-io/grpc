# go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

grpcurl -plaintext -d '{"id": 1}' -proto ../proto/chat.proto 127.0.0.1:50051 kyle_chat.Chat/SendMsgTest

grpcurl -plaintext -d '{"id": 1}' -proto ../proto/chat.proto 127.0.0.1:50051 kyle_chat.Chat/GetMsgTest
