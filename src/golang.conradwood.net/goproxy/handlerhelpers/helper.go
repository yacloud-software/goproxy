package handlerhelpers

import (
	"fmt"
	"strconv"
	"strings"
)

func Filename2ZipFilename(modname, version, filename string) string {
	vid, _ := ParseIDFromString(version)
	filename = strings.TrimPrefix(filename, modname)
	filename = strings.TrimPrefix(filename, "/")
	res := fmt.Sprintf("%s@v1.1.%d/%s", modname, vid, filename)
	fmt.Printf("Filename: %s -> %s\n", filename, res)
	return res
}
func ParseIDFromString(v string) (uint64, error) {
	sx := strings.Split(v, ".")
	if len(sx) != 3 {
		return 0, fmt.Errorf("invalid string \"%s\" (%d)\n", v, len(sx))
	}
	// new ones are 1.0.x
	res, err := strconv.ParseUint(sx[2], 10, 64)
	if err != nil {
		return 0, err
	}
	if res != 0 {
		return res, nil
	}
	// there are some old protos with version 0.x.0
	res, err = strconv.ParseUint(sx[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

