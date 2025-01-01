all: init

init:
	cp ./default/main.go main.go
	cp ./default/main_test.go main_test.go

test:
	go test

arch:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a name using 'make archive name=<name>'."; \
		exit 1; \
	fi
	@echo "Archiving files as $(name)..."
	cp main.go ./archive/$(name).go
	cp main_test.go ./archive/$(name)_test.go
