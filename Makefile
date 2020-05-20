SHELL := /bin/bash

HASH_BUILD := $(shell git rev-parse --short HEAD)

run: generatesql
	@go run ./src/

build: generatesql
	go build -o ./bin/cr_linux-x86_64_$(HASH_BUILD) ./src
	@cp ./bin/cr_linux-x86_64_$(HASH_BUILD) ./bin/cr_linux-x86_64_latest
	@echo -e "\n[LAST] cr_linux-x86_64_$(HASH_BUILD) or cr_linux-x86_64_latest \n"
	@ls -sh1 ./bin/

generatesql:
	@clear
	go run build_sql.go
	@echo -e "'build_sql.go' finished."

clear:
	rm *.db

