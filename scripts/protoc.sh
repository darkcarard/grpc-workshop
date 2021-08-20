cd ~/development/grpc/api

# Unary
protoc --go_out=plugins=grpc:. unary/unarypb/unary.proto

# Server Streaming
protoc --go_out=plugins=grpc:. server_streaming/serverstreamingpb/server_streaming.proto

# Client Streaming
protoc --go_out=plugins=grpc:. client_streaming/clientstreamingpb/client_streaming.proto

# Bidirectional Streaming
protoc --go_out=plugins=grpc:. bidirectional/bidirectionalpb/bidirectional.proto
