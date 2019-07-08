simple-server:
	@cd ./cmd/simple/server && ./server

simple-client:
	@cd ./cmd/simple/client && ./client

simple-build:
	@cd ./cmd/simple/server && go build .
	@cd ./cmd/simple/client && go build .

simple-bench:
	@cd ./cmd/simple/client && time seq 10 | xargs -n 1 ./client
	@cd ./cmd/simple/client && time seq 100 | xargs -n 1 ./client
	@cd ./cmd/simple/client && time seq 300 | xargs -n 1 ./client

rest-server:
	@cd ./cmd/rest/server && ./server

rest-client:
	@cd ./cmd/rest/client && ./client

rest-build:
	@cd ./cmd/rest/server && go build .
	@cd ./cmd/rest/client && go build .

rest-bench:
	@cd ./cmd/rest/client && time seq 10 | xargs -n 1 ./client
	@cd ./cmd/rest/client && time seq 100 | xargs -n 1 ./client
	@cd ./cmd/rest/client && time seq 300 | xargs -n 1 ./client

prot:
	@protoc --go_out=plugins=grpc:. ./proto/*/*.proto
