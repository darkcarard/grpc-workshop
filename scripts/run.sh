cd ~/development/grpc/api

# Unary
go run unary/server/server.go
go run unary/client/client.go

# Server Streaming
go run server_streaming/server/server.go
go run server_streaming/client/client.go

# Client Streaming
go run client_streaming/server/server.go
go run client_streaming/client/client.go

# Bidi Streaming
go run bidirectional/server/server.go
go run bidirectional/client/client.go
