package main

import (
	"flag"
	"golang.conradwood.net/go-easyops/linux"
	"time"
	// "golang.conradwood.net/go-easyops/utils"
	"fmt"
)

const (
	gocmd   = "/opt/yacloud/ctools/dev/go/current/go/bin/go"
	goproxy = "https://goproxy.conradwood.net,direct"
)

var (
	runtime = flag.Duration("max_runtime", time.Duration(600)*time.Second, "go stuff runtime")
)

func copy_file(src, dest string) error {
	return linux.Copy(src, dest)
	/*
		b, err := utils.ReadFile(src)
		if err != nil {
			return err
		}
		err = utils.WriteFile(dest, b)
		if err != nil {
			return err
		}
		return nil
	*/
}

// compile filename in dir
func go_compile(dir, filename string) error {
	l := linux.New()
	l.SetRuntime(*runtime)
	l.SetEnvironment(go_env())
	com := []string{gocmd, "install", filename}
	out, err := l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		fmt.Println(out)
		return err
	}
	return nil
}

// run go mod in dir
func go_mod_tidy(dir string) error {
	l := linux.New()
	l.SetRuntime(*runtime)
	l.SetEnvironment(go_env())
	com := []string{gocmd, "mod", "tidy"}
	out, err := l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		fmt.Println(out)
		return err
	}
	return nil
}

// run go mod in dir
func go_update_all(dir string) error {
	l := linux.New()
	l.SetRuntime(*runtime)
	l.SetEnvironment(go_env())
	com := []string{gocmd, "get", "-u", "all"}
	out, err := l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		fmt.Println(out)
		return err
	}
	return nil
}

func go_env() []string {
	res := []string{
		"GOPROXY=" + goproxy,
	}
	return res
}
