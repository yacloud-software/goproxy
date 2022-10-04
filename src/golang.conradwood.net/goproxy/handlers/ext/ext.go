package ext

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/http"
	"io"
	"strings"
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
	h := http.HTTP{}
	url := "https://proxy.golang.org/" + e.path + "/@v/list"
	hr := h.Get(url)
	err := hr.Error()
	if err != nil {
		fmt.Printf("Failed to access \"%s\": %s (%d)\n", url, err, hr.HTTPCode())
		fmt.Printf("body:\n%s\n\n", string(hr.Body()))
		return nil, err
	}
	body := hr.Body()
	var res []*pb.VersionInfo
	for _, line := range strings.Split(string(body), "\n") {
		v := &pb.VersionInfo{VersionName: line}
		res = append(res, v)
	}
	return res, nil

}
func (e *exthandler) GetLatestVersion(ctx context.Context) (*pb.VersionInfo, error) {
	return nil, errors.NotFound(ctx, "no latest version implemented for ext")
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
