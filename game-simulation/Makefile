.PHONY: all build run debug clean

all: build

build:
	docker build -t game-simulation .

run: build
	docker run -p 8081:80 game-simulation

debug: build
	docker run -p 8081:80 -e DEBUG=1 game-simulation

clean:
	docker rmi game-simulation
