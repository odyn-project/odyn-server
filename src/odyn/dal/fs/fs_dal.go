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

// Filesystem backend for DAL
//
// This backend is dependency-free and relatively high performance, but
// non-scalable.

// Resources are stored as directories on the local filesystem, keyed by
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
// Resource paths contain directories & files with the UUID lookup
//
//      /var/odyn-server/data/device/Leela/Toaster/__uuid
//
//          2fe4651e-fec5-474f-84b4-0792bbe0382a
//

type FsBackend struct {
    odynDir string
}

type FsConnection struct {
    backend FsBackend
}

func (backend *FsBackend) Connect() (Connection, error) {
    // Nothing needs to be done to connect
    return &FsConnection{
        backend
    }, nil
}

func (backend *FsBackend) Erase() error {
    // The (+ "/data") prevents misconfiguration from wiping the whole
    // filesystem.
    return os.RemoveAll(backend.odynDir + "/data")
}

func (backend *FsBackend) Prep() error {
    // No prep needed.
    return nil
}

func (backend *FsBackend) Migrate(start, end string) error {
    return fmt.Errorf("No migration support for filesystem db")
}

func (conn *FsConnection) Close() {
    // Nothing needs to be done
}

func (conn *FsConnection) DeleteResource(path string) {
    // Lookup the UUID

    // Delete the document file

    // Delete the lookup file
}

func LoadResource(path string) (Resource, error) {
    // Lookup the UUID

    // Read the document file
    
    // Convert the JSON contents into a dal.Resource object.
}

func SaveResource(path string) (Resource, error) {
    // Lookup the UUID

    // Read the document file
    
    // Convert the JSON contents into a dal.Resource object.
}
