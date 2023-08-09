package main

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"path/filepath"
)

type testrun1 struct {
	name string
	dir  string
}

func NewModTestRun(name string, dir string) *testrun1 {
	return &testrun1{name: name, dir: dir}
}
func (t *testrun1) Name() string {
	return t.name
}
func (t *testrun1) Sections() int {
	return 3
}
func (t *testrun1) Run(section int) error {
	dir, err := utils.FindFile("extra/tests/" + t.dir)
	if err != nil {
		return err
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	if section == 0 {
		t.Printf("Copying go.mod in dir \"%s\"\n", dir)
		err = copy_file(dir+"/go.mod.orig", dir+"/go.mod")
		if err != nil {
			return err
		}
		return go_mod_tidy(dir)
	} else if section == 1 {
		return go_update_all(dir)
	} else if section == 2 {
		return go_compile(dir, "test-compile.go")
	}
	return fmt.Errorf("no such section %d", section)
}
func (t *testrun1) Printf(format string, args ...interface{}) {
	fmt.Printf("["+t.Name()+"] "+format, args...)
}
