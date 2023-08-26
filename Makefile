tidy:
	go mod tidy
run-dev:
	docker-compose up --build
test:
	go test -v ./...
