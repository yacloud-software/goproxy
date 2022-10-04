.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} goproxy-client.go
server:
	cd server && go install ${LDFLAGS} goproxy-server.go `ls -1 *.go|grep -v goproxy-server.go`
