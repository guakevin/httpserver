tag=latest

.PHONY: build docker-build 
build: 
	go build -o dist/server .

docker-build:
	podman build -t docker.io/guakevin/httpserver:${tag}

clean:
	rm dist/server
	
.DEFAULT_GOAL := build