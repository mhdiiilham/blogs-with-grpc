genproto:
	protoc protos/blog.proto --go_out=plugins=grpc:server
	protoc protos/blog.proto --go_out=plugins=grpc:client

runserver:
	go run server/server.go
