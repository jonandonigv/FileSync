clean: 
	rm -rf bin/

build:
	go build -o /bin/filesync ./cmd/filesync/main.go
