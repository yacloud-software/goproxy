package hosts

import (
	"context"
	"fmt"
	"strings"
)

var (
	HOSTS = []string{
		"golang.conradwood.net",
		"golang.singingcat.net",
		"golang.yacloud.eu",
	}
)

func IsOneOfUs(ctx context.Context, path string) (bool, error) {
	fmt.Printf("Checking path...\"%s\"\n", path)
	ho := strings.ToLower(path)
	for _, h := range HOSTS {
		if strings.HasPrefix(ho, h) {
			return true, nil
		}
	}
	return false, nil
}
