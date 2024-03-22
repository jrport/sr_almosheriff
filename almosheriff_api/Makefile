dev:
	go build main.go
	./main

build: 
	bash ./scripts/clean_docker.sh
	bash ./scripts/build_docker.sh
	go build -o main
	./main
