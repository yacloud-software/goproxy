package handlers

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/goproxy/cacher"
	af "golang.conradwood.net/goproxy/handlers/artefact"
	"golang.conradwood.net/goproxy/handlers/ext"
	"golang.conradwood.net/goproxy/handlers/proto"
	"golang.conradwood.net/goproxy/handlers/proxy"
	"golang.conradwood.net/goproxy/hosts"
	"io"
)

type Handler interface {
	ModuleInfo() *pb.ModuleInfo
	ListVersions(ctx context.Context) ([]*pb.VersionInfo, error)
	// return the latest versioninfo from a go string (e.g. "v0.120.0")
	GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error)
	// get the zip file for a version
	GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, version string) error
	// get the go.mod (url like .../@v/[version].mod
	GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error)
	// only serve from cache if this is enabled
	CacheEnabled() bool
}

// returns a handler or a defaulthandler if it cannot determine which one
func HandlerByPath(ctx context.Context, normalisedpath, fullpath string) (Handler, error) {
	var err error

	// check with protorenderer and well-known hosts
	h1, err := proto.HandlerByPath(ctx, normalisedpath, hosts.IsOneOfUs)
	if err != nil {
		return nil, err
	}
	if h1 != nil {
		return h1, nil
	}

	// check with artefact server
	h3, err := af.HandlerByPath(ctx, normalisedpath)
	if err != nil {
		return nil, err
	}
	if h3 != nil {
		return h3, nil
	}

	h4, err := proxy.HandlerByPath(ctx, normalisedpath, fullpath)
	if err != nil {
		return nil, err
	}
	if h4 != nil {
		return h4, nil
	}

	// this one must come last. it accepts almost everything and will
	// try to use http and git to directly retrieve stuff
	// check external (probably, most likely if it is not any of the others...)
	h2, err := ext.HandlerByPath(ctx, normalisedpath)
	if err != nil {
		return nil, err
	}
	if h2 != nil {
		return h2, nil
	}

	// recheck with protorenderer if that might be available as a proto even if it is not a well-known-host
	h1, err = proto.HandlerByPath(ctx, normalisedpath, allHosts)
	if err != nil {
		return nil, err
	}
	if h1 != nil {
		return h1, nil
	}

	fmt.Printf("Returning default handler\n")
	return &defaulthandler{}, nil
}

func allHosts(ctx context.Context, path string) (bool, error) {
	return true, nil
}

type defaulthandler struct {
}

func (d *defaulthandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	panic("do not call me")
}
func (d *defaulthandler) ModuleInfo() *pb.ModuleInfo {
	return &pb.ModuleInfo{ModuleType: pb.MODULETYPE_UNKNOWN}
}
func (d *defaulthandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	panic("do not call me")
}
func (d *defaulthandler) GetZip(ctx context.Context, c *cacher.Cache, w io.Writer, v string) error {
	panic("do not call me")
}
func (d *defaulthandler) GetMod(ctx context.Context, c *cacher.Cache, version string) ([]byte, error) {
	panic("do not call me")
}
func (d *defaulthandler) CacheEnabled() bool {
	return false
}


