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

// Package context provides a means for the weaving together of data in
// structure
package context

import (
	"database/sql"
	"strings"

	"github.com/go-corelibs/maps"
	clStrings "github.com/go-corelibs/strings"
	"github.com/go-corelibs/values"
)

// Context is a convenient map[string]interface{} type, used by Go-Enjin for
// representing page front-matter and other contextual cases
//
// Context is not a replacement for the standard Go context library
type Context map[string]interface{}

// New constructs a new Context instance
func New() (ctx Context) {
	ctx = make(Context)
	return
}

// NewFromMap deep copies and transforms the given map into a Context with all
// values converted to Context types
func NewFromMap(m map[string]interface{}) Context {
	ctx := New()
	for k, v := range m {
		ctx[k] = toContextValue(v)
	}
	return ctx
}

// NewFromOsEnviron constructs a new Context from os.Environ() string K=V slices,
// unquoting any quoted values, all the keys will be exactly as present in the
// input environ strings
func NewFromOsEnviron(environs ...[]string) (c Context) {
	c = New()
	for _, environ := range environs {
		for _, pair := range environ {
			if key, value, ok := strings.Cut(pair, "="); ok {
				c[key] = clStrings.TrimQuotes(value)
			}
		}
	}
	return
}

// Len returns the number of keys in the Context
func (c Context) Len() (count int) {
	count = len(c)
	return
}

// Empty returns true if there is nothing stored in the Context
func (c Context) Empty() (empty bool) {
	empty = c.Len() == 0
	return
}

// Keys is a convenience wrapper around [maps.SortedKeys]
func (c Context) Keys() (keys []string) {
	return maps.SortedKeys(c)
}

// Copy returns a deep-copy of this Context using values.DeepCopy
//
// For problematic type cases, implementing the [deepcopy.Copyable] interface
// will enable the correct behaviour
func (c Context) Copy() (ctx Context) {
	ctx, _ = values.DeepCopy(c).(Context)
	return
}

// Apply takes a list of contexts and merges their contents into this one
func (c Context) Apply(contexts ...Context) {
	for _, cc := range contexts {
		if cc != nil {
			for k, v := range cc {
				c.Set(k, v)
			}
		}
	}
	return
}

// ApplySpecific takes a list of contexts and merges their contents into this one, keeping the keys specifically
func (c Context) ApplySpecific(contexts ...Context) {
	for _, cc := range contexts {
		if cc != nil {
			for k, v := range cc {
				c.SetSpecific(k, v)
			}
		}
	}
	return
}

// Has returns true if the given Context key exists and is not nil
func (c Context) Has(key string) (present bool) {
	present = c.Get(key) != nil
	return
}

// HasExact returns true if the specific Context key given exists and is not nil
func (c Context) HasExact(key string) (present bool) {
	if v, ok := c[key]; ok {
		present = v != nil
	}
	return
}

func (c Context) PruneEmpty() (pruned Context) {
	pruned = c.Copy()
	for k, v := range c {
		if v == nil || k == "" {
			delete(pruned, k)
			continue
		}
		switch t := v.(type) {
		case bool:
			if !t {
				delete(pruned, k)
			}
		case string:
			if t == "" {
				delete(pruned, k)
			}
		case []byte:
			if len(t) == 0 {
				delete(pruned, k)
			}
		case sql.NullTime:
			if !t.Valid {
				delete(pruned, k)
			}
		}
	}
	return
}
