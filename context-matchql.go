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
	"fmt"
	"strings"

	"github.com/go-corelibs/context/cql"
	"github.com/go-corelibs/regexps"
	"github.com/go-corelibs/values"
)

var _rxc = regexps.NewCache()

// MatchQL checks if the given context query statement matches this context
func (c Context) MatchQL(query string) (matched bool, err error) {
	var stmnt *cql.Statement
	var pErr *cql.ParseError
	if stmnt, pErr = cql.Compile(query); pErr != nil {
		err = error(pErr)
		return
	}
	stmnt = stmnt.Render()
	matched, err = c.processQueryExpression(stmnt.Expression)
	return
}

func (c Context) processQueryExpression(expr *cql.Expression) (matched bool, err error) {
	switch {

	case expr.Condition != nil:
		matched, err = c.processQueryCondition(expr.Condition)

	case expr.Operation != nil:
		matched, err = c.processQueryOperation(expr.Operation)

	}
	return
}

func (c Context) processQueryCondition(cond *cql.Condition) (matched bool, err error) {
	if cond.Left != nil && cond.Right != nil {
		var leftMatch, rightMatch bool
		if leftMatch, err = c.processQueryExpression(cond.Left); err != nil {
			return
		}
		if rightMatch, err = c.processQueryExpression(cond.Right); err != nil {
			return
		}
		switch strings.ToUpper(cond.Type) {
		case "OR":
			matched = leftMatch || rightMatch
		case "AND":
			matched = leftMatch && rightMatch
		}
	}
	return
}

func (c Context) processQueryOperation(op *cql.Operation) (matched bool, err error) {
	switch op.Type {

	case "==":
		matched, err = c.processQueryOperationEquals(*op.Left, op.Right)

	case "!=":
		if matched, err = c.processQueryOperationEquals(*op.Left, op.Right); err == nil {
			matched = !matched
		}

	default:
		err = fmt.Errorf(`%v not implemented`, op.Type)

	}
	return
}

func (c Context) processQueryOperationEquals(key string, opValue *cql.Value) (matched bool, err error) {
	switch {

	case opValue.ContextKey != nil:
		lValue := c.Get(key)
		rValue := c.Get(*opValue.ContextKey)
		matched, err = values.Compare(lValue, rValue)

	case opValue.Regexp != nil:
		if value, ok := c.Get(key).(string); ok {
			if rx, e := _rxc.Compile(*opValue.Regexp); e != nil {
				err = fmt.Errorf("error compiling regular expression")
			} else {
				matched = rx.MatchString(value)
			}
		} else {
			err = fmt.Errorf("page.%v is of type %T, expected string", key, c.Get(key))
		}

	case opValue.String != nil:
		if value, ok := c.Get(key).(string); ok {
			matched = value == *opValue.String
		} else {
			err = fmt.Errorf("page.%v is of type %T, expected string", key, c.Get(key))
		}

	}
	return
}
