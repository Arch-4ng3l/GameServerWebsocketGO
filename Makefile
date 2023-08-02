build: 
	@echo "[*] Building"
	@go build -o ./bin/app
	@cd ./web/svelte-app && npm run build
	@echo "[+] Done Building"

run: build
	@./bin/app

clean: 
	@rm -rf ./bin
	@bash ./scripts/setups.sh -c
	@rm ./assets.zip

setup:
	@bash ./scripts/setups.sh -s

test:
	@go test ./server

bench:
	@go test ./server -bench=.
	
