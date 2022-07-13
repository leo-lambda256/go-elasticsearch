// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

// Package boundaryscanner
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/highlighting.ts#L26-L30
package boundaryscanner

import "strings"

type BoundaryScanner struct {
	name string
}

var (
	Chars = BoundaryScanner{"chars"}

	Sentence = BoundaryScanner{"sentence"}

	Word = BoundaryScanner{"word"}
)

func (b BoundaryScanner) MarshalText() (text []byte, err error) {
	return []byte(b.String()), nil
}

func (b *BoundaryScanner) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "chars":
		*b = Chars
	case "sentence":
		*b = Sentence
	case "word":
		*b = Word
	default:
		*b = BoundaryScanner{string(text)}
	}

	return nil
}

func (b BoundaryScanner) String() string {
	return b.name
}