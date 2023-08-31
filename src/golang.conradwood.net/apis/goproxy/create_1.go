// client create: GoProxyClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/goproxy/goproxy.proto
   gopackage : golang.conradwood.net/apis/goproxy
   importname: ai_0
   clientfunc: GetGoProxy
   serverfunc: NewGoProxy
   lookupfunc: GoProxyLookupID
   varname   : client_GoProxyClient_0
   clientname: GoProxyClient
   servername: GoProxyServer
   gsvcname  : goproxy.GoProxy
   lockname  : lock_GoProxyClient_0
   activename: active_GoProxyClient_0
*/

package goproxy

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_GoProxyClient_0 sync.Mutex
  client_GoProxyClient_0 GoProxyClient
)

func GetGoProxyClient() GoProxyClient { 
    if client_GoProxyClient_0 != nil {
        return client_GoProxyClient_0
    }

    lock_GoProxyClient_0.Lock() 
    if client_GoProxyClient_0 != nil {
       lock_GoProxyClient_0.Unlock()
       return client_GoProxyClient_0
    }

    client_GoProxyClient_0 = NewGoProxyClient(client.Connect(GoProxyLookupID()))
    lock_GoProxyClient_0.Unlock()
    return client_GoProxyClient_0
}

func GoProxyLookupID() string { return "goproxy.GoProxy" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("goproxy.GoProxy")
}
