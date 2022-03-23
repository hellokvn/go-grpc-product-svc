proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:. -I pkg/pb

run:
	go run cmd/main.go