default: start

start:
	go build main.go
	./main
build:
	docker build -t cloudsgift/id_generator .