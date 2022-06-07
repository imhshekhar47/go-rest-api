dependency:
	go get -u
	
clean:
	if [ -d "./build" ]; then rm -r "./build"; fi;

build: clean
	if [ ! -d "./build/bin" ];  then mkdir -p "./build/bin"; fi;
	go build -o ./build/bin/go-rest-api

run: build
	./build/bin/go-rest-api

image:
	docker build -t "go-rest-api:v1" .
