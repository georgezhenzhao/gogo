GO ?= latest

build:
	go build -o gogo.exe main/main.go
	@echo "Done building."
	@echo "Run \"gogo.exe\" to launch gogo."