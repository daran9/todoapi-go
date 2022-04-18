BINARY_NAME=todoapi
SRC_DIR=cmd/todo-server

compile:
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/${BINARY_NAME} ${SRC_DIR}/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/${BINARY_NAME} ${SRC_DIR}/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/${BINARY_NAME} ${SRC_DIR}/main.go

build:
	go build -o bin/${BINARY_NAME} ${SRC_DIR}/main.go

run:
	go build -o bin/${BINARY_NAME} ${SRC_DIR}/main.go
	bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}