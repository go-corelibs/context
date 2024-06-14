// Copyright (c) 2024  The Go-CoreLibs Authors
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

package context

import (
	"encoding/json"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"

	"github.com/go-corelibs/maps"
)

// ToMap returns a values.DeepCopy of this Context, transformed to a standard
// map[string]interface{} type, recursively (the output map has no references
// to the Context or Contexts types)
func (c Context) ToMap() (out map[string]interface{}) {
	out = make(map[string]interface{})
	for k, v := range c {
		out[k] = toMapValue(v)
	}
	return
}

// ToStringMap returns this Context as a transformed map[string]string
// structure, where each key's value is checked and if it's a string, use it
// as-is and if it's anything else, run it through fmt.Sprintf("%v") to make it
// a string
func (c Context) ToStringMap() (out map[string]string) {
	out = make(map[string]string)
	for k, v := range c {
		switch t := v.(type) {
		case string:
			out[k] = t
		default:
			out[k] = fmt.Sprintf("%v", t)
		}
	}
	return
}

// ToEnviron returns this Context as a transformed []string slice where each
// key is converted to SCREAMING_SNAKE_CASE and the value is converted to a
// string (similarly to ToStringMap) and the key/value pair is concatenated
// into a single "K=V" string and appended to the output slice, sorted by key in
// natural order, suitable for use in os.Environ cases.
func (c Context) ToEnviron() (out []string) {
	return maps.ToEnviron(c)
}

// ToJSON is a convenience wrapper around [json.MarshalIndent], indented with two
// spaces and no prefix
func (c Context) ToJSON() (data []byte, err error) {
	data, err = json.MarshalIndent(c, "", "  ")
	return
}

// ToTOML is a convenience wrapper around [toml.Marshal]
func (c Context) ToTOML() (data []byte, err error) {
	data, err = toml.Marshal(c)
	return
}

// ToYAML is a convenience wrapper around [yaml.Marshal]
func (c Context) ToYAML() (data []byte, err error) {
	data, err = yaml.Marshal(c)
	return
}
