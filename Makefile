genproto:
	protoc protos/blog.proto --go_out=plugins=grpc:.

runserver:
	go run server/server.go
