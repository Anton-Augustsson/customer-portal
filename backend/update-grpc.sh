
#grpc_update()

protoc --go_out=robot-remote-controller-service --go_opt=paths=source_relative \
    --go-grpc_out=robot-remote-controller-service --go-grpc_opt=paths=source_relative \
    api/robot-remote-controller-service/robot-remote-controller-service.proto

protoc --go_out=robot-remote-controller-service --go_opt=paths=source_relative \
    --go-grpc_out=robot-remote-controller-service --go-grpc_opt=paths=source_relative \
    api/subscription-service/subscription-service.proto

protoc --go_out=subscription-service --go_opt=paths=source_relative \
    --go-grpc_out=subscription-service --go-grpc_opt=paths=source_relative \
    api/subscription-service/subscription-service.proto
