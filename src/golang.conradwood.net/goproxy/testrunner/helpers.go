package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cmdline"
	cm "golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/linux"
	"path/filepath"
	"time"
)

const (
// goproxy = "https://%s@l.conradwood.net,direct"
)

var (
	gocmd       = cmdline.GetYACloudDir() + "/ctools/dev/go/current/go/bin/go"
	goproxy     = flag.String("goproxy", "https://%s@goproxy.conradwood.net", "the goproxy to use")
	prober_user = flag.String("prober_user", "", "username for access")
	prober_pw   = flag.String("prober_password", "", "password for access")
	runtime     = flag.Duration("max_runtime", time.Duration(3)*time.Hour, "go stuff runtime")
	cur_path    string
)

func init() {
	cur_path, _ = filepath.Abs(".")
}

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
	l.SetMaxRuntime(*runtime)
	creds, err := get_auth()
	if err != nil {
		return err
	}
	l.SetEnvironment(go_env(creds))
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
	l.SetMaxRuntime(*runtime)
	creds, err := get_auth()
	if err != nil {
		return err
	}
	l.SetEnvironment(go_env(creds))
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
	l.SetMaxRuntime(*runtime)
	creds, err := get_auth()
	if err != nil {
		return err
	}
	l.SetEnvironment(go_env(creds))
	var com []string
	com = []string{gocmd, "get", "-u", "all"}
	com = []string{gocmd, "get", "-u", "./..."}
	out, err := l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		fmt.Printf("Output:\n%s\n", out)
		return err
	}
	return nil
}
func godir() string {
	godir := cur_path + "/gostuff"
	return godir
}
func go_env(c *creds) []string {
	authstring := fmt.Sprintf("%s.token:%s", c.userid, c.token)
	gop := "GOPROXY=" + fmt.Sprintf(*goproxy, authstring)
	if *debug {
		fmt.Println(gop)
		fmt.Printf("gostuff = %s\n", godir())
	}
	res := []string{
		gop,
		"GOSUMDB=off",
		"GOPATH=" + godir() + "/gopath",
		"GOTMP=" + godir() + "/gotmp",
		"GOMODCACHE=" + godir() + "/gomodcache",
		"GOCACHE=" + godir() + "/gocache",
		"GO111MODULE=on",
		"GOFLAGS=-modcacherw",
		"HOME=/tmp/x",
	}
	return res
}

type creds struct {
	userid string
	token  string
}

func get_auth() (*creds, error) {
	ctx := authremote.Context()
	if *prober_user != "" {
		apr := &auth.AuthenticatePasswordRequest{
			Email:    *prober_user,
			Password: *prober_pw,
		}
		cw, err := authremote.GetAuthClient().SignedGetByPassword(ctx, apr)
		if err != nil {
			fmt.Printf("Failed to get user \"%s\": %s\n", *prober_user, err)
			return nil, err
		}
		user := cm.VerifySignedUser(cw.User)
		if user == nil {
			s := fmt.Sprintf("could not verify user, signature incorrect or something up with auth!")
			fmt.Println(s)
			return nil, fmt.Errorf("%s", s)
		}
		return &creds{userid: user.ID, token: cw.Token}, nil

	}
	cw, err := authremote.GetAuthManagerClient().WhoAmI(ctx, &common.Void{})
	if err != nil {
		return nil, err
	}

	gtr := &auth.GetTokenRequest{DurationSecs: 600}
	r, err := authremote.GetAuthManagerClient().GetTokenForService(ctx, gtr)
	if err != nil {
		return nil, err
	}
	c := &creds{token: r.Token, userid: cw.ID}
	return c, nil
}



