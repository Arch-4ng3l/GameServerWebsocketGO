build: 
	@echo -e "\033[1;34m[*] Building\033[0m\n"
	@go build -o ./bin/app
	@cd ./web/svelte-app && npm run build
	@echo -e "\033[1;32m[+] Done Building\033[0m\n"

run: build
	@echo -e "\033[1;34m[*] Running Server\033[0m\n"
	@./bin/app

clean: 
	-@rm ./assets.zip
	-@find ./assets -type f ! -name "*.go" -exec rm {} \;
	-@rm ./logs/*
	-@rm -rf ./bin
	-@cd ./scripts && bash ./setups.sh -c

setup:
	@bash ./scripts/setups.sh -s

test:
	@go test ./server

bench:
	@go test ./server -bench=.


	
