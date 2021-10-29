package cacher

import (
	"context"
	"fmt"
	"io"
	"os"
)

type GoCacher struct {
}

func (gc *GoCacher) Set(ctx context.Context, name string, r io.ReadSeeker) error {
	fmt.Printf("Setting \"%s\"\n", name)
	return nil
}
func (gc *GoCacher) Get(ctx context.Context, name string) (io.ReadCloser, error) {
	fmt.Printf("Getting \"%s\"\n", name)
	return nil, os.ErrNotExist
}
