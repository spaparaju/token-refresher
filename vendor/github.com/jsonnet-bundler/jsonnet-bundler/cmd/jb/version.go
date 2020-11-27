// Copyright 2018 jsonnet-bundler authors
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

// +build go1.12

package main

import (
	"os"
	"path/filepath"
	"runtime/debug"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func newApp() *kingpin.Application {
	a := kingpin.New(filepath.Base(os.Args[0]), "A jsonnet package manager")
	d, ok := debug.ReadBuildInfo()
	if ok {
		return a.Version(d.Main.Version)
	}
	return a
}
