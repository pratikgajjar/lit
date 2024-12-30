all: init

init:
	cp ./default/main.go main.go
	cp ./default/main_test.go main_test.go

test:
	go test
