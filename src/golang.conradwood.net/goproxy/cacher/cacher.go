package cacher

import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/goproxy/db"
	"time"
)

var (
	debug  = flag.Bool("debug_cacher", false, "debug cacher")
	enable = flag.Bool("cacher_enable", false, "if true cache files")
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
	if !*enable {
		return false
	}
	return c.isAvailable(ctx)
}
func (c *Cache) isAvailable(ctx context.Context) bool {
	cms, err := db.DefaultDBCachedModule().ByPath(ctx, c.path)
	if err != nil {
		fmt.Printf("WARNING - cache error (%s)\n", err)
		return false
	}
	for _, cm := range cms {
		if cm.Path == c.path && cm.Version == c.version && cm.Suffix == c.suffix {
			c.Debugf("is available\n")
			return true
		}
	}
	c.Debugf("is not available\n")
	return false
}

func (c *Cache) PutBytes(ctx context.Context, data []byte) error {
	if c.isAvailable(ctx) {
		return nil
	}
	c.Debugf("saving %d bytes\n", len(data))
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

func (c *Cache) Debugf(format string, args ...interface{}) {
	if !*debug {
		return
	}
	s1 := fmt.Sprintf("[cacher %s@%s.%s] ", c.path, c.version, c.suffix)
	s2 := fmt.Sprintf(format, args...)
	fmt.Print(s1 + s2)
}
