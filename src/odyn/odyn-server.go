// Copyright 2015 Odyn Authors (see AUTHORS file for project)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "net/http"
    "odyn/log"
    "odyn/storage/fs"
    "odyn/webserver"
)

func main() {
    var err error
    log.Init("/var/log/odyn/server.log")

    // Spin up webserver
    rootMux := http.NewServeMux()
    launcher := webserver.NewLauncher()
    launcher.StartHTTPServer(":8080", rootMux)

    // Test storage engine
    engine := fs.NewEngine("/var/lib/odyn")
    err = engine.Prep()
    if err != nil {
        log.Error(err)
        return
    }

    conn, err := engine.Connect()
    if err != nil {
        log.Error(err)
        return
    }

    conn.Close()

    err = launcher.WaitForComplete()
    log.Info(err.Error())
}
