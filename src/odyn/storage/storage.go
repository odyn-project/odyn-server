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

package storage

// Storage interface.  Database backends must implement the interfaces below.

import (
    "odyn/dal"
)

// Storage Engine interface.
type Engine interface {
    Connect() (Connection, error)

    Erase() error

    Prep() error

    Migrate(startVersion, endVersion string) error
}

// Storage Connection interface for saving and loading resources.
type Connection interface {
    Close()

    DeleteResource(path string) error

    LoadResource(path string) (dal.Resource, error)

    // Create or update a resource
    SaveResource(path string, res dal.Resource) (error)
}
