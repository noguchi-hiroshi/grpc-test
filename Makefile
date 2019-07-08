simple-run-server:
	@cd ./cmd/simple/server && go run .

simple-run-client:
	@cd ./cmd/simple/client && go run .

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