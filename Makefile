proto-build:
	protoc --proto_path api/proto/ --go_out=api/gen/protogo/ --go_opt=paths=source_relative --go-grpc_out=api/gen/protogo/ --go-grpc_opt=paths=source_relative api/proto/auth.proto
