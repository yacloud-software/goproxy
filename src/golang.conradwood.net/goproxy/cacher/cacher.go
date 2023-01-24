package cacher

import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/apis/objectstore"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/goproxy/db"
	"io"
	"sync"
	"time"
)

var (
	debug   = flag.Bool("debug_cacher", false, "debug cacher")
	enable  = flag.Bool("cacher_enable", true, "if true cache files")
	biglock sync.Mutex
)

type Cache struct {
	path    string
	version string
	suffix  string
	dbmod   *pb.CachedModule
}

func NewCacher(ctx context.Context, path, version, suffix string) (*Cache, error) {
	c := &Cache{path: path, version: version, suffix: suffix}
	return c, nil
}
func (c *Cache) String() string {
	return fmt.Sprintf("%s,%s,%s", c.path, c.version, c.suffix)
}
func (c *Cache) IsAvailable(ctx context.Context) bool {
	if !*enable {
		return false
	}
	return c.isAvailable(ctx)
}
func (c *Cache) isAvailable(ctx context.Context) bool {
	if c.getDBMatch(ctx) == nil {
		return false
	}
	return true
}
func (c *Cache) getDBMatch(ctx context.Context) *pb.CachedModule {
	if c.dbmod != nil {
		return c.dbmod
	}
	cms, err := db.DefaultDBCachedModule().ByPath(ctx, c.path)
	if err != nil {
		fmt.Printf("WARNING - cache error (%s)\n", err)
		return nil
	}
	for _, cm := range cms {
		if cm.ToBeDeleted {
			continue
		}
		if cm.Path == c.path && cm.Version == c.version && cm.Suffix == c.suffix {
			c.Debugf("is available\n")
			c.dbmod = cm
			return cm
		}
	}
	c.Debugf("is not available\n")
	return nil
}

func (c *Cache) PutBytes(ctx context.Context, data []byte) error {
	biglock.Lock()
	defer biglock.Unlock()
	if c.isAvailable(ctx) {
		return nil
	}
	c.Debugf("saving %d bytes\n", len(data))
	cm := &pb.CachedModule{
		Path:    c.path,
		Version: c.version,
		Suffix:  c.suffix,
		Created: uint32(time.Now().Unix()),
		Key:     "goproxy_cache_" + utils.RandomString(70),
	}
	_, err := db.DefaultDBCachedModule().Save(ctx, cm)
	if err != nil {
		return err
	}
	err = client.PutWithID(ctx, cm.Key, data)
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

func (c *Cache) Get(ctx context.Context, f func(data []byte) error) error {
	om := c.getDBMatch(ctx)
	if om == nil {
		return errors.NotFound(ctx, "not found")
	}
	key := om.Key
	if key == "" {
		return errors.NotFound(ctx, "key missing")
	}
	om.LastUsed = uint32(time.Now().Unix())
	err := db.DefaultDBCachedModule().Update(ctx, om)
	if err != nil {
		fmt.Printf("Failed to update: %s\n", err)
	}

	srv, err := client.GetObjectStoreClient().LGet(ctx, &objectstore.GetRequest{ID: key})
	if err != nil {
		return err
	}
	for {
		r, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = f(r.Content)
		if err != nil {
			return err
		}
	}
	return nil
}
