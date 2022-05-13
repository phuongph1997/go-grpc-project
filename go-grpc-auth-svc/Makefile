proto-auth:
	protoc --proto_path=./pkg/pb --go_out=./pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/pb --go-grpc_opt=paths=source_relative \
    ./pkg/pb/*.proto


docker-auth:
    docker run --name auth-svc --network micro-service-network -p 8080:8080 -e DB_URL="postgres://root:secret@postgres12:5432/auth_svc"  auth-svc:latest
    
