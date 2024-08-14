.PHONY: all build run clean

all: build

build:
	docker build -t stat-api-server .

run: build
	docker run -p 8082:80 stat-api-server

clean:
	docker rmi stat-api-server
