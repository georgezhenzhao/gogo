GO ?= latest

build:
	go build -o gogo main/main.go
	@echo "Done building."
	@echo "Run \"gogo\" to launch gogo."