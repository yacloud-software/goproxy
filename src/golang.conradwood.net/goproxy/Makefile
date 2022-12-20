.PHONY:	client server testrunner
all:	client server testrunner
client:
	cd client && go install ${LDFLAGS} goproxy-client.go
server:
	cd server && go install ${LDFLAGS} goproxy-server.go `ls -1 *.go|grep -v goproxy-server.go`
testrunner:
	cd testrunner && go install ${LDFLAGS} goproxy-testrunner.go `ls -1 *.go|grep -v goproxy-testrunner.go`
