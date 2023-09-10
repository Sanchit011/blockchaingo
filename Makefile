build:
	go build -o ./bin/blockchaingo

run: build
	./bin/blockchaingo

test:
	go test -v ./...