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

	"github.com/go-corelibs/maths"
)

type Contexts []Context

func (c Contexts) Len() (count int) {
	return len(c)
}

func (c Contexts) FindQL(query string) (found Contexts) {
	for _, ctx := range c {
		if matched, _ := ctx.MatchQL(query); matched {
			found = append(found, ctx)
		}
	}
	return
}

func (c Contexts) SelectValues(keys ...string) (values [][]interface{}) {
	count := len(keys)
	for _, ctx := range c {
		if row := ctx.SelectValues(keys...); len(row) == count {
			values = append(values, row)
		}
	}
	return
}

func (c Contexts) SelectStringValues(keys ...string) (values [][]string) {
	count := len(keys)
	for _, ctx := range c {
		if row := ctx.SelectStringValues(keys...); len(row) == count {
			values = append(values, row)
		}
	}
	return
}

func (c Contexts) FirstValue(key string) interface{} {
	for _, ctx := range c {
		if v := ctx.Get(key); v != nil {
			return v
		}
	}
	return nil
}

func (c Contexts) Values(key string) (values []interface{}) {
	for _, ctx := range c {
		if value := ctx.Get(key); value != nil {
			values = append(values, value)
		}
	}
	return
}

func (c Contexts) FirstIntValue(key string) int {
	if v := c.FirstValue(key); v != nil {
		return maths.ToInt(v, math.MaxInt)
	}
	return math.MaxInt
}

func (c Contexts) IntValues(key string) (values []int) {
	for _, ctx := range c {
		if v := ctx.Get(key); v != nil {
			values = append(values, maths.ToInt(v, math.MaxInt))
		}
	}
	return
}

func (c Contexts) FirstInt64Value(key string) int64 {
	if v := c.FirstValue(key); v != nil {
		return maths.ToInt64(v, math.MaxInt64)
	}
	return math.MaxInt64
}

func (c Contexts) Int64Values(key string) (values []int64) {
	for _, ctx := range c {
		if v := ctx.Get(key); v != nil {
			values = append(values, maths.ToInt64(v, math.MaxInt64))
		}
	}
	return
}

func (c Contexts) FirstStringValue(key string) string {
	if v := c.FirstValue(key); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (c Contexts) StringValues(key string) (values []string) {
	for _, ctx := range c {
		if v := ctx.Get(key); v != nil {
			if value, ok := v.(string); ok {
				values = append(values, value)
			}
		}
	}
	return
}
