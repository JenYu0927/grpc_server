# compile proto into golang interface
cd proto/
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative getMax.proto

##client test tool: grpcurl

# install:
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
  
use:
grpcurl -plaintext -d '{"num1":2, "num2":7}' localhost:50051 getMax.getMax/fromTwo

grpcurl -plaintext -d '{"nums":[2,4,6,8]}' localhost:50051 getMax.getMax/fromList
