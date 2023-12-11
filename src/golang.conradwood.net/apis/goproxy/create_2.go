// client create: GoProxyTestRunnerClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/goproxy/goproxy.proto
   gopackage : golang.conradwood.net/apis/goproxy
   importname: ai_1
   clientfunc: GetGoProxyTestRunner
   serverfunc: NewGoProxyTestRunner
   lookupfunc: GoProxyTestRunnerLookupID
   varname   : client_GoProxyTestRunnerClient_1
   clientname: GoProxyTestRunnerClient
   servername: GoProxyTestRunnerServer
   gsvcname  : goproxy.GoProxyTestRunner
   lockname  : lock_GoProxyTestRunnerClient_1
   activename: active_GoProxyTestRunnerClient_1
*/

package goproxy

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_GoProxyTestRunnerClient_1 sync.Mutex
  client_GoProxyTestRunnerClient_1 GoProxyTestRunnerClient
)

func GetGoProxyTestRunnerClient() GoProxyTestRunnerClient { 
    if client_GoProxyTestRunnerClient_1 != nil {
        return client_GoProxyTestRunnerClient_1
    }

    lock_GoProxyTestRunnerClient_1.Lock() 
    if client_GoProxyTestRunnerClient_1 != nil {
       lock_GoProxyTestRunnerClient_1.Unlock()
       return client_GoProxyTestRunnerClient_1
    }

    client_GoProxyTestRunnerClient_1 = NewGoProxyTestRunnerClient(client.Connect(GoProxyTestRunnerLookupID()))
    lock_GoProxyTestRunnerClient_1.Unlock()
    return client_GoProxyTestRunnerClient_1
}

func GoProxyTestRunnerLookupID() string { return "goproxy.GoProxyTestRunner" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("goproxy.GoProxyTestRunner")
   AddService("goproxy.GoProxyTestRunner")
}


