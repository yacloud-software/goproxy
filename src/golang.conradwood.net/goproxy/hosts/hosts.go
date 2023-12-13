package hosts

import (
	"context"
	//	"fmt"
	"golang.conradwood.net/goproxy/config"
	"strings"
)

var ()

// TODO: use ipmanager/or dns lookup mechanism or so sort this out instead of hardcoded list
func IsOneOfUs(ctx context.Context, path string) (bool, error) {
	//fmt.Printf("Checking path...\"%s\"\n", path)
	ho := strings.ToLower(path)
	HOSTS := config.GetConfig().LocalHosts
	for _, h := range HOSTS {
		if strings.HasPrefix(ho, h) {
			return true, nil
		}
	}
	return false, nil
}





