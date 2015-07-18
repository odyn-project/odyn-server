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

package dal

// Data Abstraction Layer

type ResourceType int
const (
    UnknownResource = iota

    DeviceResource
    OrgResource
    PolicyResource
    TeamResource
    UserResource
)

type DAL interface {
    Connect() (Connection, error)

    EraseDb() error

    PrepDb() error

    MigrateDb(keyspace, startVersion, endVersion string) error
}

type Connection interface {
    Close()

    CreateResource(path string) (Resource, error)

    DeleteResource(res Resource)

    Resource(path string) (Resource, error)

    ResourceExists(path string) (bool, error)
}

type Property interface {
    Attribute(attr string) (PropVal, error)

    Datatype() PropDatatype

    SetAttribute(attr string, val PropVal)

    SetValue(val PropVal) error

    Child(name string) (Property, error)

    Value() PropVal
}

type Resource interface {
    // Get the resource's path, such as "device/Leela/Toaster"
    Path() string

    // Get a property of the resource by name
    Property(name string) (Property, error)

    Type() ResourceType

    // Reload all properties.  Existing Property objects will be orphaned.
    Refresh() (error)

    // Write all modified properties to the database
    Save() (error)
}
