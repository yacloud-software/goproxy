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
		fmt.Printf("[cacher] WARNING - cache error (%s)\n", err)
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
		fmt.Printf("[cacher] ERROR!!!! Put for cachedmodule ID=%d failed: %s\n", cm.ID, utils.ErrorString(err))
		cm.PutFailed = true
		cm.PutError = utils.ErrorString(err)
	} else {
		cm.Size = uint64(len(data))
	}
	db.DefaultDBCachedModule().Update(context.Background(), cm)
	return err
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
		return errors.NotFound(ctx, "key not found")
	}
	key := om.Key
	if key == "" {
		return errors.NotFound(ctx, "key missing")
	}
	om.LastUsed = uint32(time.Now().Unix())
	err := db.DefaultDBCachedModule().Update(ctx, om)
	if err != nil {
		fmt.Printf("[cacher] Failed to update: %s\n", err)
	}

	srv, err := client.GetObjectStoreClient().LGet(ctx, &objectstore.GetRequest{ID: key})
	if err != nil {
		fmt.Printf("[cacher] failed to get key %s from objectstore - objectstore claims no knowledge of key: %s\n", key, err)
		c.objectstore_failed(ctx)
		return err
	}
	for {
		r, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("[cacher] failed to get key %s from objectstore - objectstore error in Recv(): %s\n", key, err)
			c.objectstore_failed(ctx)
			return err
		}
		err = f(r.Content)
		if err != nil {
			fmt.Printf("[cacher] failed to get key %s from objectstore - error from callback: %s\n", key, err)
			return err
		}
	}
	c.objectstore_succeeded(ctx)
	return nil
}

func (c *Cache) Description(ctx context.Context) (string, error) {
	om := c.getDBMatch(ctx)
	if om == nil {
		return "", errors.NotFound(ctx, "dbmatch not found")
	}
	return fmt.Sprintf("[cached: ID=%d, path=\"%s\", version=\"%s\", suffix=\"%s\", Created=%s, LastUsed=%s]", om.ID, om.Path, om.Version, om.Suffix, utils.TimestampString(om.Created), utils.TimestampString(om.LastUsed)), nil

}
func (c *Cache) Key(ctx context.Context) (string, error) {
	om := c.getDBMatch(ctx)
	if om == nil {
		return "", errors.NotFound(ctx, "dbmatch not found")
	}
	key := om.Key
	if key == "" {
		return "", errors.NotFound(ctx, "dbkey missing")
	}
	return key, nil
}

func (c *Cache) objectstore_failed(ctx context.Context) {
	om := c.getDBMatch(ctx)
	if om == nil {
		return
	}
	now_t := uint32(time.Now().Unix())
	if om.FailingSince == 0 {
		om.FailingSince = now_t
	}
	om.FailCounter = om.FailCounter + 1
	om.LastFailed = now_t

	// delete if all three are true:
	// 1. at least 10 failures (failcounter)
	// 2. first failure was at least 30 minutes ago
	// 3. last failure was no more than 5 minutes ago
	t := time.Unix(int64(om.FailingSince), 0)
	if time.Since(t) > time.Duration(30)*time.Minute {
		if om.FailCounter > 10 {
			om.ToBeDeleted = true
		}
	}
	err := db.DefaultDBCachedModule().Update(context.Background(), om)
	if err != nil {
		fmt.Printf("[cacher] Failed to update cachedmodule: %s\n", err)
	}
}
func (c *Cache) objectstore_succeeded(ctx context.Context) {
	om := c.getDBMatch(ctx)
	if om == nil {
		return
	}
	upd := false
	if om.FailingSince != 0 {
		upd = true
		om.FailingSince = 0
	}
	if om.FailCounter != 0 {
		upd = true
		om.FailCounter = 0
	}
	if upd {
		err := db.DefaultDBCachedModule().Update(context.Background(), om)
		if err != nil {
			fmt.Printf("[cacher] Failed to update cachedmodule: %s\n", err)
		}
	}
}



