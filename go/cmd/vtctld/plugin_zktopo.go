/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// Imports and register the 'zookeeper' topo.Server and its Explorer.

import (
	"github.com/youtube/vitess/go/vt/servenv"
	"github.com/youtube/vitess/go/vt/vtctld"
	"github.com/youtube/vitess/go/vt/zktopo"
)

func init() {
	// Wait until flags are parsed, so we can check which topo server is in use.
	servenv.OnRun(func() {
		if zkServer, ok := ts.Impl.(*zktopo.Server); ok {
			vtctld.HandleExplorer("zk", zktopo.NewZkExplorer(zkServer.GetZConn()))
		}
	})
}
