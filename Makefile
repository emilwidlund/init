BINARY_NAME=init
 
build:
	go mod tidy
	go build -o bin/${BINARY_NAME} main.go
 
clean:
	go clean
	rm bin/${BINARY_NAME}