package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	h2g "golang.conradwood.net/apis/h2gproxy"
)

type getpath_streamer struct {
	ctx  context.Context
	srv  pb.GoProxy_GetPathServer
	size int
}

func (e *echoServer) GetPath(req *pb.GetPathRequest, srv pb.GoProxy_GetPathServer) error {
	sr := &h2g.StreamRequest{Path: req.Path}
	sv := &getpath_streamer{ctx: srv.Context(), srv: srv}
	err := e.streamHTTP(sr, sv)
	if err != nil {
		fmt.Printf("Error %s for path %s\n", err, req.Path)
	}
	fmt.Printf("Sent %d bytes for path \"%s\"\n", sv.size, req.Path)
	return err

}

func (gs *getpath_streamer) Send(d *h2g.StreamDataResponse) error {
	bd := &pb.BinData{Data: d.Data}
	gs.size = gs.size + len(bd.Data)
	err := gs.srv.Send(bd)
	return err
}
func (gs *getpath_streamer) Context() context.Context {
	return gs.ctx
}





