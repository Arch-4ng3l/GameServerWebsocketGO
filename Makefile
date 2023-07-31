build: 
	@go build -o ./bin/app

run: build
	@./bin/app

clean: 
	@rm ./bin/app

test:
	@go test ./server

bench:
	@go test ./server -bench=.
	
