package main

import (
	"context"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/goproxy"
	h2g "golang.conradwood.net/apis/h2gproxy"
	"golang.conradwood.net/goproxy/db"
	"strings"
)

func (sr *SingleRequest) CheckOverride(ctx context.Context, req *h2g.StreamRequest, srv streamer) (bool, error) {
	if strings.Contains(req.Host, "goproxy-dev") {
		return false, nil
	}
	p := req.Path
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimSuffix(p, "/@v/list")
	sr.Printf("Checking override for \"%s\"\n", p)
	t := 0
	if strings.HasSuffix(req.Path, "list") {
		t = 1
	}

	if t == 0 {
		return false, nil
	}

	ovs, err := db.DefaultDBOverride().ByPackage(ctx, p)
	if err != nil {
		return true, err
	}
	if len(ovs) == 0 {
		return false, nil
	}
	ov := ovs[0]
	err = sendBytes(srv, ov.List)
	if err != nil {
		return true, err
	}
	return false, nil
}

func (e *echoServer) AddOverride(ctx context.Context, req *pb.Override) (*common.Void, error) {
	ovs, err := db.DefaultDBOverride().ByPackage(ctx, req.Package)
	if err != nil {
		return nil, err
	}
	for _, ov := range ovs {
		err = db.DefaultDBOverride().DeleteByID(ctx, ov.ID)
		if err != nil {
			return nil, err
		}
	}
	_, err = db.DefaultDBOverride().Save(ctx, req)
	if err != nil {
		return nil, err
	}
	return &common.Void{}, nil
}


