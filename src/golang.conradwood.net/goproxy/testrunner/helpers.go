package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/auth"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"time"
)

const (
	gocmd   = "/opt/yacloud/ctools/dev/go/current/go/bin/go"
	goproxy = "https://%s@goproxy.conradwood.net,direct"
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
	creds, err := get_auth()
	if err != nil {
		fmt.Printf("Error getting auth: %s\n", utils.ErrorString(err))
		panic("no auth")
	}
	authstring := fmt.Sprintf("%s.token:%s", creds.userid, creds.token)
	res := []string{
		"GOPROXY=" + fmt.Sprintf(goproxy, authstring),
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

	cw, err := authremote.GetAuthManagerClient().WhoAmI(ctx, &common.Void{})
	if err != nil {
		return nil, err
	}

	gtr := &auth.GetTokenRequest{DurationSecs: 600}
	r, err := authremote.GetAuthManagerClient().GetTokenForMe(ctx, gtr)
	if err != nil {
		return nil, err
	}
	c := &creds{token: r.Token, userid: cw.ID}
	return c, nil
}
