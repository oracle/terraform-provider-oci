// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type urlParts []interface{}

// request exists to support mocking.
// TODO: These marshal fns should be behind one single marshal
// func that takes a urlBuilder and an http.Method constant and
// returns a new http.Request via http.NewRequest to the requestor
// for further auth processing.
type request interface {
	marshalBody() ([]byte, error)
	marshalHeader() http.Header
	marshalURL(string, string, urlBuilderFn) (val string, e error)
}

// requestDetails is the concrete implementation of request.
// optional should always be a struct from request_options.go
// required should always be an anonymous struct that can,
// optionally, have one of the unexported structs from
// request_requirements.go embedded.
type requestDetails struct {
	region      string
	urlTemplate string
	ids         urlParts
	name        resourceName
	optional    interface{}
	required    interface{}
}

func objToJSONMap(val interface{}) (map[string]interface{}, error) {
	marshaled, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	jsonMap := make(map[string]interface{})
	if string(marshaled) == "null" {
		return jsonMap, nil
	}
	err = json.Unmarshal(marshaled, &jsonMap)
	return jsonMap, err
}

func (r *requestDetails) marshalBody() (marshaled []byte, e error) {
	if bm, ok := r.required.(bodyMarshaller); ok {
		return bm.body(), nil
	}

	required := r.required
	if required == nil {
		required = struct{}{}
	}
	requiredMap, err := objToJSONMap(required)
	if err != nil {
		return nil, err
	}

	opts := r.optional
	if opts == nil {
		opts = struct{}{}
	}

	optMap, err := objToJSONMap(opts)
	if err != nil {
		return nil, err
	}

	// Override options with required in case of overlap
	for k, v := range requiredMap {
		optMap[k] = v
	}

	marshaled, e = json.Marshal(optMap)
	return
}

func (r *requestDetails) marshalHeader() http.Header {

	// TODO: Error handling here.
	var rHeader, oHeader http.Header
	rHeader, _ = NewHeaderFromStruct(r.required)
	oHeader, _ = NewHeaderFromStruct(r.optional)

	for k, v := range rHeader {
		oHeader[k] = v
	}

	if md, ok := r.optional.(MetadataUnmarshallable); ok {
		// md may be nil and still be "ok"
		if md != (*PutObjectOptions)(nil) {
			prefix := "opc-meta-"
			for name, val := range md.GetMetadata() {
				oHeader[prefix+name] = []string{val}
			}
		}
	}

	return oHeader
}

func (r *requestDetails) marshalQueryString() (vals url.Values, e error) {
	var rVals url.Values
	if rVals, e = query.Values(r.required); e != nil {
		return
	}
	if vals, e = query.Values(r.optional); e != nil {
		return
	}
	for k, v := range rVals {
		vals[k] = v
	}
	return
}

func (r *requestDetails) marshalURL(urlTemplate string, region string, urlFn urlBuilderFn) (val string, e error) {
	var q url.Values
	if q, e = r.marshalQueryString(); e != nil {
		return
	}
	val = urlFn(urlTemplate, region, r.name, q, r.ids...)
	return
}
