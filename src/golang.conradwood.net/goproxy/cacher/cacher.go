package cacher

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/goproxy/db"
	"time"
)

type Cache struct {
	path    string
	version string
	suffix  string
}

func NewCacher(ctx context.Context, path, version, suffix string) (*Cache, error) {
	c := &Cache{path: path, version: version, suffix: suffix}
	return c, nil
}

func (c *Cache) IsAvailable(ctx context.Context) bool {
	cms, err := db.DefaultDBCachedModule().ByPath(ctx, c.path)
	if err != nil {
		fmt.Printf("WARNING - cache error (%s)\n", err)
		return false
	}
	for _, cm := range cms {
		if cm.Path == c.path && cm.Version == c.version && cm.Suffix == c.suffix {
			return true
		}
	}
	return false
}

func (c *Cache) PutBytes(ctx context.Context, data []byte) error {
	cm := &pb.CachedModule{
		Path:    c.path,
		Version: c.version,
		Suffix:  c.suffix,
		Created: uint32(time.Now().Unix()),
	}
	_, err := db.DefaultDBCachedModule().Save(ctx, cm)
	if err != nil {
		return err
	}
	return nil
}
