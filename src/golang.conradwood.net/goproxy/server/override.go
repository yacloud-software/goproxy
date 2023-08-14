package main

import (
	"context"
	h2g "golang.conradwood.net/apis/h2gproxy"
)

func CheckOverride(ctx context.Context, req *h2g.StreamRequest, srv streamer) (bool, error) {
	return false, nil
}
