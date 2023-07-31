build: 
	@go build -o ./bin/app

run: build
	@./bin/app

clean: 
	@rm ./bin/app
	@bash ./scripts/setups.sh -c

setup:
	@bash ./scripts/setups.sh -s

test:
	@go test ./server

bench:
	@go test ./server -bench=.
	
