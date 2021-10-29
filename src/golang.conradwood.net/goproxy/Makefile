.PHONY:	client server
all:	client server
client:
	cd client && go install goproxy-client.go
server:
	cd server && go install goproxy-server.go `ls -1 *.go|grep -v goproxy-server.go`
