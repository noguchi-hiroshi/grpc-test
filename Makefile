# simple
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

# rest
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

# server stream
serverstream-server:
	@cd ./cmd/serverstream/server && ./server

serverstream-client:
	@cd ./cmd/serverstream/client && ./client

serverstream-build:
	@cd ./cmd/serverstream/server && go build .
	@cd ./cmd/serverstream/client && go build .

# client stream
clientstream-server:
	@cd ./cmd/clientstream/server && ./server

clientstream-client:
	@cd ./cmd/clientstream/client && ./client

clientstream-build:
	@cd ./cmd/clientstream/server && go build .
	@cd ./cmd/clientstream/client && go build .

# bi stream
bistream-server:
	@cd ./cmd/bistream/server && ./server

bistream-client:
	@cd ./cmd/bistream/client && ./client

bistream-build:
	@cd ./cmd/bistream/server && go build .
	@cd ./cmd/bistream/client && go build .

prot:
	@protoc --go_out=plugins=grpc:. ./proto/simple/*.proto
	@protoc --go_out=plugins=grpc:. ./proto/serverstream/*.proto
	@protoc --go_out=plugins=grpc:. ./proto/clientstream/*.proto
	@protoc --go_out=plugins=grpc:. ./proto/bistream/*.proto
