package ext

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/http"
	"io"
)

type exthandler struct {
	path string
}

func HandlerByPath(ctx context.Context, path string) (*exthandler, error) {
	res := &exthandler{path: path}
	return res, nil
}
func (e *exthandler) ModuleInfo() *pb.ModuleInfo {
	mi := &pb.ModuleInfo{ModuleType: pb.MODULETYPE_EXTERNALMODULE, Exists: true}
	return mi
}
func (e *exthandler) ListVersions(ctx context.Context) ([]*pb.VersionInfo, error) {
	fmt.Printf("No versions for ext yet")
	return nil, nil
}
func (e *exthandler) GetVersion(ctx context.Context, v string) (*pb.VersionInfo, error) {
	return nil, nil
}
func (e *exthandler) GetZip(ctx context.Context, w io.Writer, v string) error {
	h := http.HTTP{}
	hr := h.Get("https://proxy.golang.org/" + e.path + "/@v/" + v + ".zip")
	err := hr.Error()
	if err != nil {
		return err
	}
	_, err = w.Write(hr.Body())
	if err != nil {
		return err
	}
	return nil
}
func (e *exthandler) GetMod(ctx context.Context, version string) ([]byte, error) {
	h := http.HTTP{}
	hr := h.Get("https://proxy.golang.org/" + e.path + "/@v/" + version + ".mod")
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	return hr.Body(), nil
}
