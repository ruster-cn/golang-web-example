all: clean build

build:
	go build -o bin/op-server cmd/main.go
	go build -o bin/cluster-import-cli tools/cluster-import-cli/main.go

run:
	go run cmd/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/op-server-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/op-server-linux-arm64 cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/op-server-freebsd-386 cmd/main.go

clean:
	rm -rf bin
