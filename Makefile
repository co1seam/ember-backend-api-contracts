generate-auth: # Generate code for auth servicefrom proto files
	protoc -I proto proto/auth/*.proto \
      --go_out=./gen/go \
      --go_opt=paths=source_relative \
      --go-grpc_out=./gen/go \
      --go-grpc_opt=paths=source_relative

generate-media:
	protoc -I proto proto/media/*.proto \
          --go_out=./gen/go \
          --go_opt=paths=source_relative \
          --go-grpc_out=./gen/go \
          --go-grpc_opt=paths=source_relative