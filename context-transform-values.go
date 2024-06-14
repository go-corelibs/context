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
	"github.com/go-corelibs/values"
)

func toContextValue(input interface{}) interface{} {
	switch t := input.(type) {

	case Contexts, Context:
		// input is correct, assume nop recurse
		return t

	case []interface{}:
		slice := make([]interface{}, len(t))
		for idx, item := range t {
			slice[idx] = toContextValue(item)
		}
		return slice

	case []map[string]interface{}:
		slice := make([]Context, len(t))
		for idx, item := range t {
			slice[idx] = make(Context)
			for k, v := range item {
				slice[idx][k] = toContextValue(v)
			}
		}
		return slice

	case map[string]interface{}:
		m := make(Context)
		for k, v := range t {
			m[k] = toContextValue(v)
		}
		return m
	}
	return values.DeepCopy(input)
}

func toMapValue(input interface{}) interface{} {
	switch t := input.(type) {

	case Contexts:
		slice := make([]map[string]interface{}, t.Len())
		for idx, ctx := range t {
			slice[idx] = ctx.ToMap()
		}
		return slice

	case Context:
		return t.ToMap()

	case []interface{}:
		slice := make([]interface{}, len(t))
		for idx, item := range t {
			slice[idx] = toMapValue(item)
		}
		return slice

	case []map[string]interface{}:
		slice := make([]map[string]interface{}, len(t))
		for idx, item := range t {
			slice[idx] = make(map[string]interface{})
			for k, v := range item {
				slice[idx][k] = toMapValue(v)
			}
		}
		return slice

	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range t {
			m[k] = toMapValue(v)
		}
		return m

	}
	return values.DeepCopy(input)
}
