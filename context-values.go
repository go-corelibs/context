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
	"math"
	"time"

	"github.com/go-corelibs/maths"
	"github.com/go-corelibs/values"
)

// Bytes returns the key's value as a byte slice, returning the given default if
// not found or not actually a byte slice value.
func (c Context) Bytes(key string, def ...[]byte) []byte {
	if v := c.Get(key); v != nil {
		if s, ok := v.([]byte); ok {
			return s
		}
	}
	for _, d := range def {
		if d != nil {
			return d
		}
	}
	return nil
}

// String returns the key's value as a string, returning the given default if
// not found or not actually a string value.
func (c Context) String(key string, def ...string) string {
	if v := c.Get(key); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	for _, d := range def {
		if d != "" {
			return d
		}
	}
	return ""
}

// StringOrStrings returns the key's value as a list of strings and if the key's
// actual value is not a list of strings, return that as a list of one string
func (c Context) StringOrStrings(key string) (values []string) {
	if v := c.Get(key); v != nil {
		if s, ok := v.(string); ok {
			values = []string{s}
			return
		}
		if vi, ok := v.([]interface{}); ok {
			for _, i := range vi {
				if s, ok := i.(string); ok {
					values = append(values, s)
				}
			}
			return
		}
	}
	return
}

// Strings returns the key's value as a list of strings, returning an empty list
// if not found or not actually a list of strings
func (c Context) Strings(key string) (values []string) {
	if v := c.Get(key); v != nil {
		if vs, ok := v.([]string); ok {
			values = vs
			return
		} else if vi, ok := v.([]interface{}); ok {
			for _, vii := range vi {
				if viis, ok := vii.(string); ok {
					values = append(values, viis)
				}
			}
			return
		}
	}
	return
}

// DefaultStrings is a wrapper around Strings() and returns the given default
// list of strings if the key is not found
func (c Context) DefaultStrings(key string, def ...[]string) []string {
	if v := c.Get(key); v != nil {
		if s, ok := v.([]string); ok {
			return s
		}
	}
	for _, d := range def {
		if d != nil {
			return d
		}
	}
	return nil
}

func (c Context) Slice(key string) (list []interface{}, ok bool) {
	if v := c.Get(key); v != nil {
		list, ok = v.([]interface{})
	}
	return
}

func (c Context) Bool(key string, def ...bool) bool {
	if v := c.Get(key); v != nil {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	for _, d := range def {
		if d {
			return d
		}
	}
	return false
}

func (c Context) Boolean(key string) (value, ok bool) {
	v := c.Get(key)
	if ok = v != nil; ok {
		value, ok = values.ToBoolValue(v)
	}
	return
}

func (c Context) ValueAsInt(key string, def ...int) int {
	if v := c.Get(key); v != nil {
		if i := maths.ToInt(v); i != math.MaxInt {
			return i
		}
	}
	for _, d := range def {
		return d
	}
	return 0
}

func (c Context) ValueAsInt64(key string, def ...int64) int64 {
	if v := c.Get(key); v != nil {
		if i := maths.ToInt64(v); i != math.MaxInt64 {
			return i
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (c Context) Int(key string, def ...int) int {
	if v := c.Get(key); v != nil {
		return maths.ToInt(v, def...)
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxInt
}

func (c Context) Int64(key string, def ...int64) int64 {
	if v := c.Get(key); v != nil {
		return maths.ToInt64(v, def...)
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxInt64
}

func (c Context) Uint(key string, def ...uint) uint {
	if v := c.Get(key); v != nil {
		return maths.ToUint(v, def...)
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxUint
}

func (c Context) Uint64(key string, def ...uint64) uint64 {
	if v := c.Get(key); v != nil {
		return maths.ToUint64(v, def...)
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxUint64
}

func (c Context) Float64(key string, def ...float64) float64 {
	if v := c.Get(key); v != nil {
		return maths.ToFloat64(v, def...)
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxFloat64
}

func (c Context) Time(key string, def ...time.Time) time.Time {
	if v := c.Get(key); v != nil {
		if t, ok := v.(time.Time); ok {
			return t
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return time.Time{}
}

func (c Context) TimeDuration(key string, def ...time.Duration) time.Duration {
	if v := c.Get(key); v != nil {
		if t, ok := v.(time.Duration); ok {
			return t
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

// Context looks for the given key and if the value is of Context type, returns it
func (c Context) Context(key string) (ctx Context) {
	if v := c.Get(key); v != nil {
		var ok bool
		if ctx, ok = v.(Context); !ok {
			ctx, _ = v.(map[string]interface{})
		}
	}
	if ctx == nil {
		ctx = New()
	}
	return
}

func (c Context) FirstString(key string) (value string, ok bool) {
	var v interface{}
	var list []string
	if v, ok = c[key]; v != nil {
		if value, ok = v.(string); ok {
		} else if list, ok = v.([]string); ok && len(list) > 0 {
			value = list[0]
		}
	}
	return
}

func (c Context) SelectValues(keys ...string) (selected []interface{}) {
	for _, key := range keys {
		if value := c.Get(key); value != nil {
			selected = append(selected, value)
		}
	}
	return
}

func (c Context) SelectStringValues(keys ...string) (selected []string) {
	var found []string
	for _, key := range keys {
		if value := c.String(key, ""); value != "" {
			found = append(found, value)
		}
	}
	if len(keys) == len(found) {
		selected = found
	}
	return
}

func (c Context) Select(keys ...string) (selected map[string]interface{}) {
	selected = make(map[string]interface{})
	for _, key := range keys {
		if v, ok := c[key]; ok {
			selected[key] = v
		}
	}
	return
}
