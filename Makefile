server: server.go
	CGO_ENABLED=0  GOOS=linux  go build -a -a -ldflags '-extldflags "-static"' -o server server.go


clean:
	rm -f server
