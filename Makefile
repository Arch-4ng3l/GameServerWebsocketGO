build: 
	@echo -e "\033[1;34m[*] Building\033[0m\n"
	@go build -o ./bin/app
	@cd ./web/svelte-app && npm run build
	@echo -e "\033[1;32m[+] Done Building\033[0m\n"

run: build
	@echo -e "\033[1;34m[*] Running Server\033[0m\n"
	@./bin/app

clean: 
	-@cd ./scripts && bash ./setups.sh -c
	-@rm -rf ./bin
	-@rm ./assets.zip
	-@rm ./logs/*

setup:
	@bash ./scripts/setups.sh -s

test:
	@go test ./server

bench:
	@go test ./server -bench=.


	
