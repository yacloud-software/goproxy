package cacher

import (
	"fmt"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/goproxy/db"
	"time"
)

func init() {
	go evict_loop()
}
func evict_loop() {
	t := time.Duration(3) * time.Second
	for {
		time.Sleep(t)
		t = time.Duration(63) * time.Second
		err := evict_now()
		if err != nil {
			fmt.Printf("could not evict: %s\n", err)
		}
	}
}
func evict_now() error {
	ctx := authremote.Context()
	cms, err := db.DefaultDBCachedModule().ByToBeDeleted(ctx, true)
	if err != nil {
		return err
	}
	for _, cm := range cms {
		err := client.EvictNoResult(ctx, cm.Key)
		if err != nil {
			fmt.Printf("failed to evict #%d: %s\n", cm.ID, err)
			continue
		}
		err = db.DefaultDBCachedModule().DeleteByID(ctx, cm.ID)
		if err != nil {
			fmt.Printf("Evicted #%d but failed to delete row: %s\n", cm.ID, err)
			continue
		}
	}
	return nil
}


