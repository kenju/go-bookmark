.PHONY: setup
setup:
	go install golang.org/dl/go1.17.3
	go1.17.3 download

.PHONY: test
test:
	go1.17.3 test -v ./...

.PHONY: build
build:
	mkdir -p build
	go1.17.3 build -o build/bookmark

.PHONY: locate
locate: build
	mv build/bookmark ~/bin/bookmark
