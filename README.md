[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/context)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/context)](https://goreportcard.com/report/github.com/go-corelibs/context)

# context - the weaving together of data in structure

This package primarily provides a simple map type, called "Context":

``` go
type Context map[string]interface{}
```

The Context can be used like a normal map type, which also means that these
sorts of things are not concurrecy-safe. The purpose of these structures are
intended for ephemeral cases such as the unmarshalling of JSON data and
accessing the information as conveniently as possible.

Like most [Go-CoreLibs], [Go-Curses] and [Go-Enjin] projects, Context follows a
developer-focused design - "make it nice to work with". However we need to
define what "nice" means in this context because of the limitations of human
language.

A project is defined as "nice to work with" when (in no particular order):

- it works well with the Go standard library
- it does not make decisions for the developer
- it is convenient to type into any text editor
- it is convenient to conceptualize it's use cases
- it is designed with a balance of security and convenience
- it does not require arcane, occult or estoric knowledge to use effectively

## What is a Context anyways?

Merriam-Webster has a lovely blurb on the meaning of the word "context", copied
here (without permission) because of it's relevance and because I love their
work:

```
Did you know?

In its earliest uses (documented in the 15th century), context meant "the
weaving together of words in language." This sense, now obsolete, developed
logically from the word's source in Latin, contexere "to weave or join
together." Context now most commonly refers to the environment or setting in
which something (whether words or events) exists. When we say that something
is contextualized, we mean that it is placed in an appropriate setting, one in
which it may be properly considered.
```
  -- https://www.merriam-webster.com/dictionary/context
  -- please support your favorite dictionary, even if it's not Merriam-Webster
  -- the languages we use in life defines the limits of what can be thought of

For this project in specific, the meaning of the word "context" is actually
both the obsolete and modern meanings.

Within the [Go-Enjin] project, where this package manifested first, there are
a number of use cases:

- it is the front-matter portion of any parsed page
- it is the means for parsing page editor form submissions
- it is a type used for unmarshalling JSON data
- it is a type used for scanning DB query rows

Now that this project has been migrated from it's original home at:
`github.com/go-enjin/be/pkg/context` to this new home as a formal [Go-CoreLibs]
project, this package will end up being used anytime a type for managing the
weaving together of data in structure makes sense.

# Installation

``` shell
> go get github.com/go-corelibs/context@latest
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
