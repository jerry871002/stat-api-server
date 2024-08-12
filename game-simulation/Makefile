.PHONY: all build run clean

all: build

build:
	docker build -t game-simulation .

run: build
	docker run -p 8080:8080 game-simulation

debug: build
	docker run -p 8080:8080 -e DEBUG=1 game-simulation

clean:
	docker rmi game-simulation
