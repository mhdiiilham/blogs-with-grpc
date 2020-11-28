genproto:
	protoc protos/blog.proto --go_out=plugins=grpc:.

runserver:
	go run main.go

test:
	go test ./... $(v)
