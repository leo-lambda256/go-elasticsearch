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
// https://github.com/elastic/elasticsearch-specification/tree/33e8a1c9cad22a5946ac735c4fba31af2da2cec2

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/templateformat"
)

// RoleTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/33e8a1c9cad22a5946ac735c4fba31af2da2cec2/specification/security/get_role/types.ts#L50-L53
type RoleTemplate struct {
	Format   *templateformat.TemplateFormat `json:"format,omitempty"`
	Template Script                         `json:"template"`
}

func (s *RoleTemplate) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRoleTemplate returns a RoleTemplate.
func NewRoleTemplate() *RoleTemplate {
	r := &RoleTemplate{}

	return r
}
