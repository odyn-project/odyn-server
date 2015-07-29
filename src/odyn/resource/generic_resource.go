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

package resource

import (
    "fmt"
     "odyn/storage"
)

type GenericResource struct {
    conn storage.Connection
    json map[string]interface{}
    path string
}

// Get the resource's path, such as "device/Leela/Toaster"
func (res *GenericResource) Path() string {
    return res.path
}

func (res *GenericResource) StorageConnection() storage.Connection {
    return res.conn
}

// Get a property of the resource by name
func (res *GenericResource) Property(name string) (*Property, error) {
    _, ok := res.json[name].(map[string]interface{})
    if !ok {
        return nil, fmt.Errorf("Resource '%s' property '%s' not found.", 
                name, res.Path())
    }

    // TODO
    return nil, fmt.Errorf("Not implemented")
    //return jsonToProperty(propJson)
}

func ResourceFromJson(json map[string]interface{}) (Resource, error) {
    // TODO implement
    return nil, fmt.Errorf("Not implemented")
}
