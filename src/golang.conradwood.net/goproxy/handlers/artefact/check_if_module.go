package artefact

import (
	"context"
	"fmt"
	artefact "golang.conradwood.net/apis/artefact"
	// pb "golang.conradwood.net/apis/goproxy"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/utils"
	"time"
)

var (
	af_mod_cache  = cache.New("artefact_is_mod_cache", time.Duration(24)*time.Hour, 1000)
	max_neg_cache = time.Duration(30) * time.Second
)

type af_mod_cache_entry struct {
	err     error
	res     bool
	created time.Time
}

func (af *afhandler) check_if_has_file(ctx context.Context, build, artefactid uint64, filename string) (bool, error) {
	key := fmt.Sprintf("%d_%d", build, artefactid)
	or := af_mod_cache.Get(key)
	if or != nil {
		of := or.(*af_mod_cache_entry)
		if of.err == nil || time.Since(of.created) < max_neg_cache {
			return of.res, of.err
		}
	}
	b, err := af.verify_has_file(ctx, build, artefactid, filename)
	of := &af_mod_cache_entry{
		err:     err,
		created: time.Now(),
		res:     b,
	}
	af_mod_cache.Put(key, of)
	if err != nil {
		return false, err
	}
	return b, nil
}

// ask artefact server for file
func (af *afhandler) verify_has_file(ctx context.Context, build, artefactid uint64, filename string) (bool, error) {
	//TODO: fix gitbuilder to create zip files with v1.1.%d instead of 0.1.%d (or both??)
	//v := &pb.VersionInfo{VersionName: fmt.Sprintf("v0.1.%d", build)}

	dlr := &artefact.DirListRequest{
		Build:      build,
		ArtefactID: artefactid,
		Dir:        filename,
	}
	ds, err := artefact.GetArtefactClient().GetDirListing(ctx, dlr)

	if err != nil {
		af.Printf("Build #%d does not have a module in \"%s\" (%s)\n", build, dlr.Dir, utils.ErrorString(err))
		return false, nil
	}
	if len(ds.Files) == 0 {
		af.Printf("Build #%d does not have any of the required files for modules in \"%s\"", build, dlr.Dir)
		return false, nil
	}
	af.Printf("Build #%d added as module in \"%s\"\n", build, dlr.Dir)
	return true, nil

}





