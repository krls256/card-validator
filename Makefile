lint:
	golangci-lint run -v

test-card:
	go test ./card

gen-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
	 --go-grpc_opt=require_unimplemented_servers=false \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/grpc/card_validator.proto

docker-build:
	docker build -t card-validator .

docker-run:
	docker run -p 6600:6600 -p 6610:6610 card-validator