// Copyright 2016 Google Inc. All Rights Reserved.
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

// This program generates tmpldata.go
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// template files to parse.
// map keys are format names.
var files = map[string]struct {
	file string
	html bool
}{
	"html":    	    {"template.html", true},
	"htmlElements": {"template-elements.html", true},
	"md":           {"template.md", false},
	"offline":      {"template-offline.html", true},
}

func main() {
	log.SetFlags(0)
	w, err := os.Create("tmpldata.go")
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewWriter(w)
	defer func() {
		buf.Flush()
		w.Close()
	}()

	buf.WriteString(tmpldataHead)
	for f, t := range files {
		b, err := ioutil.ReadFile(t.file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(buf, "\t%q: &template{\n", f)
		fmt.Fprintf(buf, "\t\thtml: %v,\n", t.html)
		fmt.Fprintf(buf, "\t\tbytes: []byte{\n")
		writeBytes(buf, b)
		fmt.Fprintf(buf, "\t\t},\n\t},\n")
	}
	buf.WriteString(tmpldataFoot)
}

func writeBytes(w io.Writer, b []byte) {
	for i, x := range b {
		if i%10 == 0 {
			fmt.Fprint(w, "\t\t\t")
		}
		fmt.Fprintf(w, "%#x,", x)
		if i%10 == 9 || i == len(b)-1 {
			fmt.Fprint(w, "\n")
		}
	}
}

const tmpldataHead = `// Copyright 2016 Google Inc. All Rights Reserved.
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

// This file is auto-generated by gen-tmpldata.go.
// All modifications will be lost.

package render

var tmpldata = map[string]*template{
`

const tmpldataFoot = `}
`
