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

package cql

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Statement struct {
	Expression  *Expression `parser:"@@" json:"expressions,omitempty"`
	ContextKeys []string    `parser:"" json:"context-keys,omitempty"`
	rendered    bool        `parser:""`
}

func (s *Statement) Render() (out *Statement) {
	out = new(Statement)
	if s.Expression != nil {
		out.Expression = s.Expression.Render()
	}
	out.ContextKeys = append(out.ContextKeys, s.ContextKeys...)
	out.rendered = true
	return
}

func (s *Statement) String() (query string) {
	if s.rendered {
		return
	}
	var compile func(expr *Expression)
	compile = func(expr *Expression) {
		switch {
		case expr.Operation != nil:
			var right string
			switch {
			case expr.Operation.Right.ContextKey != nil:
				right = "." + *expr.Operation.Right.ContextKey
			case expr.Operation.Right.String != nil:
				right = *expr.Operation.Right.String
			case expr.Operation.Right.Regexp != nil:
				right = "m" + *expr.Operation.Right.Regexp
			case expr.Operation.Right.Int != nil:
				right = fmt.Sprintf("%v", *expr.Operation.Right.Int)
			case expr.Operation.Right.Float != nil:
				right = fmt.Sprintf("%v", *expr.Operation.Right.Float)
			case expr.Operation.Right.Bool != nil:
				right = fmt.Sprintf("%v", *expr.Operation.Right.Bool)
			case expr.Operation.Right.Nil != nil:
				right = "nil"
			}
			query += fmt.Sprintf("(.%s %s %s)", *expr.Operation.Left, expr.Operation.Type, right)

		case expr.Condition != nil:
			query += "("
			compile(expr.Condition.Left)
			query += " " + strings.ToUpper(expr.Condition.Type) + " "
			compile(expr.Condition.Right)
			query += ")"
		}
	}
	compile(s.Expression)
	return
}

func (s *Statement) Stringify() (out string) {
	b, _ := json.MarshalIndent(s.Render(), "", "  ")
	out = string(b)
	return
}
