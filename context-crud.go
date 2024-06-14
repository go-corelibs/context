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

	"github.com/go-corelibs/maps"
)

// Set CamelCases the given key and sets that within this Context
func (c Context) Set(key string, value interface{}) Context {
	key = strcase.ToCamel(key)
	c[key] = value
	return c
}

// SetSpecific is like Set(), without CamelCasing the key
func (c Context) SetSpecific(key string, value interface{}) Context {
	c[key] = value
	return c
}

// SetKV is a convenience wrapper around maps.SetKV
func (c Context) SetKV(key string, value interface{}) (err error) {
	err = maps.SetKV(c, key, value)
	return
}

// Get is a convenience wrapper around GetKV
func (c Context) Get(key string) (value interface{}) {
	_, value = c.GetKV(key)
	return
}

// GetKV is a convenience wrapper around maps.GetKV
func (c Context) GetKV(key string) (k string, v interface{}) {
	k, v = maps.GetKV(c, key)
	return
}

// Delete is a convenience wrapper around maps.DeleteKV
func (c Context) Delete(key string) (deleted bool) {
	return maps.DeleteKV(c, key)
}

// DeleteKeys is a batch wrapper around Delete()
func (c Context) DeleteKeys(keys ...string) {
	for _, key := range keys {
		_ = c.Delete(key)
	}
}
