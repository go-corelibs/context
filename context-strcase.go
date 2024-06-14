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
	"github.com/iancoleman/strcase"
)

// CamelizeKeys transforms all keys within the Context to be of CamelCased form
func (c Context) CamelizeKeys() {
	var remove []string
	for k, v := range c {
		if vc, ok := v.(Context); ok {
			vc.CamelizeKeys()
		} else if vm, ok := v.(map[string]interface{}); ok {
			Context(vm).CamelizeKeys()
		}
		if modified := strcase.ToCamel(k); k != modified {
			remove = append(remove, k)
			c.SetSpecific(modified, v)
		}
	}
	c.DeleteKeys(remove...)
}

// LowerCamelizeKeys transforms all keys within the Context to be of lowerCamelCased form
func (c Context) LowerCamelizeKeys() {
	var remove []string
	for k, v := range c {
		if vc, ok := v.(Context); ok {
			vc.LowerCamelizeKeys()
		} else if vm, ok := v.(map[string]interface{}); ok {
			Context(vm).LowerCamelizeKeys()
		}
		if modified := strcase.ToLowerCamel(k); k != modified {
			remove = append(remove, k)
			c.SetSpecific(modified, v)
		}
	}
	c.DeleteKeys(remove...)
}

// KebabKeys transforms all keys within the Context to be of kebab-cased form
func (c Context) KebabKeys() {
	var remove []string
	for k, v := range c {
		if vc, ok := v.(Context); ok {
			vc.KebabKeys()
		} else if vm, ok := v.(map[string]interface{}); ok {
			Context(vm).KebabKeys()
		}
		if modified := strcase.ToKebab(k); k != modified {
			remove = append(remove, k)
			c.SetSpecific(modified, v)
		}
	}
	c.DeleteKeys(remove...)
}
