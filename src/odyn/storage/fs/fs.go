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

package fs

// Filesystem storage engine
//
// This storage backend is dependency-free and relatively high performance, but
// non-scalable.

import (
     "code.google.com/p/go-uuid/uuid"
     "encoding/json"
     "fmt"
     "io/ioutil"
     "odyn/storage"
     "os"
     "strings"
)

// Documents are stored as directories on the local filesystem, keyed by
// internal UUID.
//
//      /var/odyn-server/data/res/UUID/
//
// Properties are stored in the __doc file
//
//      /var/odyn-server/data/res/UUID/__doc
//
//      {
//          "system" : {
//              "username" : {
//                  ":datatype" : "string",
//                  "value" : "Leela"
//              },
//              "email" : {
//                  ":datatype" : "string",
//                  "value" : "leela@PlanetExpress.com"
//              }
//              "password" : {
//                  ":datatype" : "password",
//                  "value" : "zXdt5d4ug78jige"
//              }
//          }
//      }
//
// Document paths contain directories & files with the UUID lookup
//
//      /var/odyn-server/data/device/Leela/Toaster/__uuid
//
//          2fe4651e-fec5-474f-84b4-0792bbe0382a
//

// Internal structure for Filesystem Storage Engine.
// Implements storage.Engine
type FsEngine struct {
    odynDir string
}

// Internal structure for Filesystem Storage Connection.
// Implements storage.Connection
type FsConnection struct {
    engine *FsEngine
    dataDir string
}

func (engine *FsEngine) Connect() (storage.Connection, error) {
    // Nothing needs to be done to connect
    return &FsConnection{
        engine,
        engine.odynDir + "/data",
    }, nil
}

func (engine *FsEngine) Erase() error {
    // The (+ "/data") prevents misconfiguration from wiping the whole
    // filesystem.
    return os.RemoveAll(engine.odynDir + "/data")
}

func (engine *FsEngine) Prep() error {
    // No prep needed.
    return nil
}

func (engine *FsEngine) Migrate(start, end string) error {
    return fmt.Errorf("No migration support for filesystem db")
}

func (conn *FsConnection) Close() {
    // Nothing needs to be done
}

func (conn *FsConnection) createUUID(path string) (string, error) {
    // Generate random UUID
    id := uuid.New()

    // MkdirAll for path
    err := os.MkdirAll(conn.dataDir + "/" + path, 0644)
    if err != nil {
        return "", err
    }

    // Write UUID to __uuid file for path
    filename := conn.dataDir + "/" + path + "/__uuid"
    err = ioutil.WriteFile(filename, []byte(id), 0644)
    if err != nil {
        return "", err
    }

    return id, nil
}

func (conn *FsConnection) lookupUUID(path string) (string, error) {
    // Read UUID from path file
    buf, err := ioutil.ReadFile(conn.dataDir + path + "/__uuid")
    if (err != nil) {
        return "", err
    }
    s := string(buf)

    // verify that it resembles uuid
    if len(s) != 36 {
        return "", fmt.Errorf(conn.dataDir + path + 
                "/__uuid file contents is not a UUID")
    }

    return s, nil
}

func (conn *FsConnection) lookupOrCreateUUID(path string) (string, error) {
    id, err := conn.lookupUUID(path)
    if err == nil {
        return id, nil
    }

    id, err = conn.createUUID(path)
    if err == nil {
        return id, nil
    }

    return "", err
}

func (conn *FsConnection) DeleteDocument(path string) error {
    // Lookup the UUID
    id, err := conn.lookupUUID(path)
    if err != nil {
        return err
    }

    // Delete the document file & directory
    err = os.RemoveAll(conn.dataDir + path)
    if err != nil {
        return err
    }

    // Delete the lookup file
    err = os.RemoveAll(conn.dataDir + "/res/" + id)
    if err != nil {
        return err
    }

    return nil
}

func (conn *FsConnection)LoadDocument(path string) (interface{}, error) {
    // Lookup the UUID
    id, err := conn.lookupUUID(path)
    if err != nil {
        return nil, err
    }

    // Read the document file
    buf, err := ioutil.ReadFile(conn.dataDir + "/res/" + id + "/__doc")
    if (err != nil) {
        return nil, err
    }

    // Parse the JSON
    var doc map[string]interface{}
    decoder := json.NewDecoder(strings.NewReader(string(buf)))
    err = decoder.Decode(&doc)
    if err != nil {
        return nil, err
    }
    
    return doc, nil

}

func (conn *FsConnection)SaveDocument(path string, doc interface{}) (error) {
    var id string
    var err error

    // Lookup the UUID
    id, err = conn.lookupUUID(path)
    if err != nil {
        // UUID not found for this resource.  Create it.
        id, err = conn.createUUID(path)
        if err != nil {
            return err
        }
    }

    // Serialize to JSON
    jsonBytes, err := json.MarshalIndent(doc, "", "    ")
    if err != nil {
        return err
    }

    // Save to document file
    filename := conn.dataDir + "/res/" + id + "/__doc"
    err = ioutil.WriteFile(filename, jsonBytes, 0644)
    if err != nil {
        return err
    }

    return nil
}

// Initialize a new Filesystem Storage Engine object that uses <storageDir> for
// data storage.
func NewEngine(storageDir string) storage.Engine {
    return &FsEngine{
        odynDir : storageDir,
    }
}
