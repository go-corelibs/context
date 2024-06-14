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

// Package cql provides a syntax parser for a "Context Query Language"
package cql

import (
	"sort"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/maruel/natural"

	"github.com/go-corelibs/maps"
)

const (
	gIdent      = `\b([a-zA-Z][.a-zA-Z0-9]*)\b`
	gInteger    = `\b(\d+)\b`
	gFloat      = `\b(\d*\.\d+)\b`
	gString     = `'[^']*'|"[^"]*"`
	gRegexp     = `/(.+?)/|\!(.+?)\!|\@(.+?)\@|\~(.+?)\~`
	gOperators  = `==|=\~|\!=|\!\~|[.,()]`
	gWhitespace = `\s+`
)

var (
	gLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: `Keyword`, Pattern: `(?i)\b(TRUE|FALSE|NULL|IS|NOT|AND|OR|IN)\b`},
		{Name: `Ident`, Pattern: gIdent},
		{Name: `Int`, Pattern: gInteger},
		{Name: `Float`, Pattern: gFloat},
		{Name: `String`, Pattern: gString},
		{Name: `Regexp`, Pattern: gRegexp},
		{Name: `Operators`, Pattern: gOperators},
		{Name: `whitespace`, Pattern: gWhitespace},
	})
	gParser = participle.MustBuild[Statement](
		participle.Lexer(gLexer),
		participle.CaseInsensitive("Keyword"),
	)
)

func EBNF() (ebnf string) {
	return gParser.String()
}

func Compile(query string) (stmnt *Statement, err *ParseError) {
	err = nil
	query = strings.TrimSpace(query)

	var participleError error
	if stmnt, participleError = gParser.ParseString("cql", query); participleError != nil && participleError.Error() != "" {
		err = newParseError(query, participleError)
		return
	}

	var extract func(expr *Expression) (keys []string)
	extract = func(expr *Expression) (keys []string) {
		unique := make(map[string]bool)
		switch {
		case expr.Operation != nil:
			unique[*expr.Operation.Left] = true
			if expr.Operation.Right.ContextKey != nil {
				unique[*expr.Operation.Right.ContextKey] = true
			}
		case expr.Condition != nil:
			for _, key := range extract(expr.Condition.Left) {
				unique[key] = true
			}
			for _, key := range extract(expr.Condition.Right) {
				unique[key] = true
			}
		}
		keys = maps.Keys(unique)
		return
	}

	contextKeys := extract(stmnt.Expression)
	sort.Sort(natural.StringSlice(contextKeys))
	stmnt.ContextKeys = contextKeys
	return
}
