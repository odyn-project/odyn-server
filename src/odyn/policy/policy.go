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

package policy

// Odyn's permission system.
//
// Every Resource and every Property has a Policy which governs who can access
// it (the "Actor"), and what actions they can perform on it.
//
//
// Each Property has an ACL (Access Control List).  The ACL lists the Users,
// Teams, Organizations, and other Devices who have access.
//
// A PropertyPolicy document is structured like this:
//
//  // For device/Leela/Toaster:
//
//  {
//      "acl" : {
//          "@self" : "*",
//          "@owner" : "*",
//          "@admin" : "*",
//          "user/Leela" : "gsmcd",
//          "device/PlanetExpress/Refrigerator" : "g",
//          "team/PlanetExpress/Crew" : "g",
//          "team/PlanetExpress/Execs" : "G",
//          "user/doorman" : "20150803202208:gs",
//          "user/doorman2" : ["g", "20150803202208-20150804235959:smcd"],
//      }
//  }
//
//  "*" = All access rights
//  "g" = Get property value
//  "s" = Set property value
//  "m" = Modify property metadata
//  "c" = Clear history
//  "d" = Delete property
//  "G" = Get property value in aggregate/anonymized form only
//
//  A permission may start with a timestamp (expiry date) or have a timestamp
//  range.
//
//
//  APPLICATION ACCESS:
//
//  If an App is accessing a property on behalf of a User, first it is checked
//  to see if the User has access.  Then the User's AppPolicy document is
//  checked.
//
//  // For user/Leela:
//
//  {
//      "app_perms" : {
//          "@all" : "location,diagnostics",
//          "device/PlanetExpress/Refridgerator" : "*",
//      }
//  }
//
//  Each property has a ":scope" metadata attribute that lists the application
//  scope.
//
//  User can control which devices & which scopes to grant 
//

type Actor struct
    ActorPath string
    AppPath string
}

type ResourceACL interface {
    CanReadValue(actor Actor)
}

type PropertyPermissions struct {
    CanGetValue bool
    CanSetValue bool
    CanSetMetadata bool
    CanClearHistory bool
    CanDelete bool
}

type PropertyUserPermissions struct {
    UserPath string
    Perms PropertyPermissions
}

type PropertyACL interface {
    UserPermsissionList()[]
}
