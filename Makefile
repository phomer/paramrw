# Makefile
#

all: test
	@echo "Completed"

test: build
	./paramrw

build:
	#go build -a -v -o paramrw random.go
	go build -v -o paramrw random.go

