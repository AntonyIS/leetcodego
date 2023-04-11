build:
	go build -o bin/leetcodego

serve:build
	./bin/leetcodego

test:
	go test -v ./...
