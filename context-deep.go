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
	"sort"
	"strconv"

	"github.com/maruel/natural"
)

// DeepKeys returns a natural sorted list of .deep.keys representing the entire
// context
func (c Context) DeepKeys() (keys []string) {
	for k, v := range c {
		dk := "." + k
		keys = append(keys, dk)
		switch t := v.(type) {
		case Contexts:
			for idx, ctx := range t {
				for _, deeperKey := range ctx.DeepKeys() {
					keys = append(keys, dk+"["+strconv.Itoa(idx)+"]"+deeperKey)
				}
			}
		case Context:
			for _, deeper := range t.DeepKeys() {
				keys = append(keys, dk+deeper)
			}
		case map[string]interface{}:
			for _, deeper := range Context(t).DeepKeys() {
				keys = append(keys, dk+deeper)
			}
		}
	}
	sort.Sort(natural.StringSlice(keys))
	return
}

// AsDeepKeyed returns a deep-key flattened version of this context
//
// Examples:
//
//	Input   Context{"one": map[string]interface{}{"two": "deep"}}
//	Output  Context{".one.two": "deep"}
//
//	Input   Context{"one": Contexts{{"two": "deep"}}}
//	Output  Context{".one[0].two": "deep"}
func (c Context) AsDeepKeyed() (out Context) {
	out = Context{}
	for _, k := range c.Keys() {
		dk := "." + k
		switch t := c[k].(type) {
		case Contexts:
			for idx, ctx := range t {
				for deeperKey, deeperValue := range ctx.AsDeepKeyed() {
					out[dk+"["+strconv.Itoa(idx)+"]"+deeperKey] = deeperValue
				}
			}
		case Context:
			for deeperKey, deeperValue := range t.AsDeepKeyed() {
				out[dk+deeperKey] = deeperValue
			}
		case map[string]interface{}:
			for deeperKey, deeperValue := range Context(t).AsDeepKeyed() {
				out[dk+deeperKey] = deeperValue
			}
		default:
			out[dk] = t
		}
	}
	return
}
