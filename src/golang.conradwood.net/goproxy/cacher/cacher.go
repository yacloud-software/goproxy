package cacher

import (
	"context"
	"fmt"
	"golang.conradwood.net/goproxy/db"
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
