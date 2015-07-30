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

package webserver

import (
    "github.com/gorilla/context"
    "net/http"
    "reflect"
)

// Utililty library that simplifies the launch of a webserver on HTTP/HTTPS

type Webserver interface {
    // Launch a goroutine that runs an HTTP server.
    // <host> specifies the hostname and/or port for listening, for example:
    // "localhost:8080" or ":80".
    // <handler> is the HTTP handler to use.
    //
    // Call WaitForComplete to get error results.
    StartHTTPServer(host string, handler http.Handler)

    // Launch a goroutine that runs an HTTPS server.
    // Similar to StartHTTPServe.
    // <certFile> is filename of server certificate.
    // <privKeyFile> is filename of server private key.
    //
    // Call WaitForComplete to get error results.
    StartHTTPSServer(host, certFile, privKeyFile string, handler http.Handler)

    // Block the current goroutine indefinitely until any of the running web
    // server(s) halt.
    WaitForComplete() error
}

type WebserverObj struct {
    resultChans [](chan error)
}

func (server *WebserverObj) StartHTTPServer(host string, handler http.Handler) {
    resultChan := make(chan error)
    server.resultChans = append(server.resultChans, resultChan)
    go func() {
        srv := &http.Server{
            Addr: host,
            Handler: context.ClearHandler(handler),
        }
        err := srv.ListenAndServe()
        resultChan <- err
        close(resultChan)
    }()
}

func (server *WebserverObj)StartHTTPSServer(host, certFile, privKeyFile string, 
        handler http.Handler) {

    resultChan := make(chan error)
    server.resultChans = append(server.resultChans, resultChan)
    go func() {
        srv := &http.Server{
            Addr: host,
            Handler: context.ClearHandler(handler),
        }
        err := srv.ListenAndServeTLS(certFile, privKeyFile)
        resultChan <- err
        close(resultChan)
    }()
}

func (server *WebserverObj)WaitForComplete() error {
    cases := make([]reflect.SelectCase, len(server.resultChans))
    for i, ch := range server.resultChans {
        cases[i] = reflect.SelectCase{
            Dir: reflect.SelectRecv, 
            Chan: reflect.ValueOf(ch),
        }
    }
    _, value, _ := reflect.Select(cases)
    return value.Interface().(error)
}

func NewLauncher() Webserver {
    return &WebserverObj{[](chan error){}}
}
