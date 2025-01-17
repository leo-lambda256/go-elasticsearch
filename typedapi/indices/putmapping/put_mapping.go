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

// Updates the index mappings.
package putmapping

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewPutMapping type alias for index.
type NewPutMapping func(index string) *PutMapping

// NewPutMappingFunc returns a new instance of PutMapping with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutMappingFunc(tp elastictransport.Interface) NewPutMapping {
	return func(index string) *PutMapping {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Updates the index mappings.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-put-mapping.html
func New(tp elastictransport.Interface) *PutMapping {
	r := &PutMapping{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutMapping) Raw(raw io.Reader) *PutMapping {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutMapping) Request(req *Request) *PutMapping {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutMapping: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mapping")

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutMapping) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutMapping query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putmapping.Response
func (r PutMapping) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the PutMapping headers map.
func (r *PutMapping) Header(key, value string) *PutMapping {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names the mapping should be added to
// (supports wildcards); use `_all` or omit to add the mapping on all indices.
// API Name: index
func (r *PutMapping) Index(index string) *PutMapping {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *PutMapping) AllowNoIndices(allownoindices bool) *PutMapping {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *PutMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutMapping {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *PutMapping) IgnoreUnavailable(ignoreunavailable bool) *PutMapping {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *PutMapping) MasterTimeout(duration string) *PutMapping {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *PutMapping) Timeout(duration string) *PutMapping {
	r.values.Set("timeout", duration)

	return r
}

// WriteIndexOnly When true, applies mappings only to the write index of an alias or data
// stream
// API name: write_index_only
func (r *PutMapping) WriteIndexOnly(writeindexonly bool) *PutMapping {
	r.values.Set("write_index_only", strconv.FormatBool(writeindexonly))

	return r
}

// DateDetection Controls whether dynamic date detection is enabled.
// API name: date_detection
func (r *PutMapping) DateDetection(datedetection bool) *PutMapping {
	r.req.DateDetection = &datedetection

	return r
}

// Dynamic Controls whether new fields are added dynamically.
// API name: dynamic
func (r *PutMapping) Dynamic(dynamic dynamicmapping.DynamicMapping) *PutMapping {
	r.req.Dynamic = &dynamic

	return r
}

// DynamicDateFormats If date detection is enabled then new string fields are checked
// against 'dynamic_date_formats' and if the value matches then
// a new date field is added instead of string.
// API name: dynamic_date_formats
func (r *PutMapping) DynamicDateFormats(dynamicdateformats ...string) *PutMapping {
	r.req.DynamicDateFormats = dynamicdateformats

	return r
}

// DynamicTemplates Specify dynamic templates for the mapping.
// API name: dynamic_templates
func (r *PutMapping) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *PutMapping {
	r.req.DynamicTemplates = dynamictemplates

	return r
}

// FieldNames_ Control whether field names are enabled for the index.
// API name: _field_names
func (r *PutMapping) FieldNames_(fieldnames_ *types.FieldNamesField) *PutMapping {

	r.req.FieldNames_ = fieldnames_

	return r
}

// Meta_ A mapping type can have custom meta data associated with it. These are
// not used at all by Elasticsearch, but can be used to store
// application-specific metadata.
// API name: _meta
func (r *PutMapping) Meta_(metadata types.Metadata) *PutMapping {
	r.req.Meta_ = metadata

	return r
}

// NumericDetection Automatically map strings into numeric data types for all fields.
// API name: numeric_detection
func (r *PutMapping) NumericDetection(numericdetection bool) *PutMapping {
	r.req.NumericDetection = &numericdetection

	return r
}

// Properties Mapping for a field. For new fields, this mapping can include:
//
// - Field name
// - Field data type
// - Mapping parameters
// API name: properties
func (r *PutMapping) Properties(properties map[string]types.Property) *PutMapping {

	r.req.Properties = properties

	return r
}

// Routing_ Enable making a routing value required on indexed documents.
// API name: _routing
func (r *PutMapping) Routing_(routing_ *types.RoutingField) *PutMapping {

	r.req.Routing_ = routing_

	return r
}

// Runtime Mapping of runtime fields for the index.
// API name: runtime
func (r *PutMapping) Runtime(runtimefields types.RuntimeFields) *PutMapping {
	r.req.Runtime = runtimefields

	return r
}

// Source_ Control whether the _source field is enabled on the index.
// API name: _source
func (r *PutMapping) Source_(source_ *types.SourceField) *PutMapping {

	r.req.Source_ = source_

	return r
}
