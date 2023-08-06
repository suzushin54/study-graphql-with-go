.PHONY: help run gen
.DEFAULT_GOAL := help


install-tools:
	go install github.com/99designs/gqlgen@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

## start: Run the server
start:
	go run ./server.go

## gen: Generate Go code from GraphQL schema
gen:
	gqlgen generate

gen_orm:
	sqlboiler sqlite3

## help: Show this help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
