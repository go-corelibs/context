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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeep(t *testing.T) {
	Convey("DeepKeys", t, func() {
		ctx := Context{
			"one": Contexts{
				{"two": "many"},
			},
			"more": map[string]interface{}{
				"this": "that",
			},
		}
		So(ctx.DeepKeys(), ShouldEqual, []string{
			".more",
			".more.this",
			".one",
			".one[0].two",
		})
	})

	Convey("AsDeepKeyed", t, func() {
		ctx := Context{
			"one": Contexts{
				{"two": "many"},
			},
			"more": map[string]interface{}{
				"this": "that",
			},
		}
		So(ctx.AsDeepKeyed(), ShouldEqual, Context{
			".more.this":  "that",
			".one[0].two": "many",
		})
	})
}
