dbup:
	docker pull mongo:4.0.4 && docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4

run:
	go run main.go

test:
	go test -v ./router