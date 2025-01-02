all: init

init:
	cp ./default/main.go ./sol/sol.go
	cp ./default/main_test.go ./sol/sol_test.go

test:
	go test

arch:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a name using 'make archive name=<name>'."; \
		exit 1; \
	fi
	@echo "Archiving files as $(name)..."
	cp ./sol/sol.go ./archive/$(name).go
	cp ./sol/sol_test.go ./archive/$(name)_test.go
