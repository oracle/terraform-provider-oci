package baremetal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/gotascii/go-http-header/header"
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
	marshalURL(urlBuilderFn) (val string, e error)
}

// requestDetails is the concrete implementation of request.
// optional should always be a struct from request_options.go
// required should always be an anonymous struct that can,
// optionally, have one of the unexported structs from
// request_requirements.go embedded.
type requestDetails struct {
	ids      urlParts
	name     resourceName
	optional interface{}
	required interface{}
}

func (r *requestDetails) marshalBody() (marshaled []byte, e error) {
	if bm, ok := r.required.(bodyMarshaller); ok {
		return bm.body(), nil
	}

	if marshaled, e = json.Marshal(r.required); e != nil {
		return
	}

	if r.optional != nil {
		var oBody []byte
		if oBody, e = json.Marshal(r.optional); e != nil {
			return
		}
		if len(oBody) > 2 {
			marshaled = marshaled[:len(marshaled)-1]
			marshaled = append(marshaled, []byte(",")...)
			oBody = oBody[1:]
			marshaled = append(marshaled, oBody...)
		}
	}

	return
}

func (r *requestDetails) marshalHeader() http.Header {

	// TODO: Error handling here.
	var rHeader, oHeader http.Header
	rHeader, _ = header.NewFromStruct(r.required)
	oHeader, _ = header.NewFromStruct(r.optional)

	for k, v := range rHeader {
		oHeader[k] = v
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

func (r *requestDetails) marshalURL(urlFn urlBuilderFn) (val string, e error) {
	var q url.Values
	if q, e = r.marshalQueryString(); e != nil {
		return
	}
	val = urlFn(r.name, q, r.ids...)
	return
}
