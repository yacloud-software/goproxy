package handlers

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/goproxy/handlers/ext"
	"golang.conradwood.net/goproxy/handlers/proto"
	"io"
)

type Handler interface {
	ModuleInfo() *pb.ModuleInfo
	ListVersions(ctx context.Context) ([]*pb.VersionInfo, error)
	// return a versioninfo from a go string (e.g. "v0.120.0")
	GetVersion(ctx context.Context, version string) (*pb.VersionInfo, error)
	// get the zip file for a version
	GetZip(ctx context.Context, w io.Writer, version string) error
	// get the go.mod (url like .../@v/[version].mod
	GetMod(ctx context.Context, version string) ([]byte, error)
}

// returns a handler or a defaulthandler if it cannot determine which one
func HandlerByPath(ctx context.Context, path string) (Handler, error) {
	var err error
	h1, err := proto.HandlerByPath(ctx, path)
	if err != nil {
		return nil, err
	}
	//	fmt.Printf("Handler: %#v\n", h)
	if h1 != nil {
		return h1, nil
	}

	h2, err := ext.HandlerByPath(ctx, path)
	if err != nil {
		return nil, err
	}
	//	fmt.Printf("Handler: %#v\n", h)
	if h2 != nil {
		return h2, nil
	}
	fmt.Printf("Returning default handler\n")
	return &defaulthandler{}, nil
}

type defaulthandler struct {
}

func (d *defaulthandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	panic("do not call me")
}
func (d *defaulthandler) ModuleInfo() *pb.ModuleInfo {
	return &pb.ModuleInfo{ModuleType: pb.MODULETYPE_UNKNOWN}
}
func (d *defaulthandler) GetVersion(ctx context.Context, v string) (*pb.VersionInfo, error) {
	panic("do not call me")
}
func (d *defaulthandler) GetZip(ctx context.Context, w io.Writer, v string) error {
	panic("do not call me")
}
func (d *defaulthandler) GetMod(ctx context.Context, version string) ([]byte, error) {
	panic("do not call me")
}
