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
//                  ":datatype" : "string",
//                  "value" : "leela@PlanetExpress.com"
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

type FsDal struct {
}

type FsResource struct {
    path string
}

func (dal *FsDal) Connect() (Connection, error) {
}

func (res *FsResource) Path() string {
    return res.path
}

func (res *FsResource) Path() string {
    return res.path
}
