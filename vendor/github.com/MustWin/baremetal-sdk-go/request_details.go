package baremetal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type urlParts []interface{}

type requestDetails struct {
	name     resourceName
	ids      urlParts
	required interface{}
	optional interface{}
}

func (r *requestDetails) query() (vals url.Values, e error) {
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

func (r *requestDetails) url(urlFn urlBuilderFn) (val string, e error) {
	var q url.Values
	if q, e = r.query(); e != nil {
		return
	}
	val = urlFn(r.name, q, r.ids...)
	return
}

func (r *requestDetails) header() http.Header {
	var rHeader, oHeader http.Header
	if rhd, ok := r.required.(HeaderGenerator); ok == true {
		rHeader = rhd.Header()
	} else {
		rHeader = http.Header{}
	}

	if ohd, ok := r.optional.(HeaderGenerator); ok == true {
		oHeader = ohd.Header()
	} else {
		oHeader = http.Header{}
	}
	for k, v := range rHeader {
		oHeader[k] = v
	}
	return oHeader
}

func (r *requestDetails) getBody() (marshaled []byte, e error) {
	if marshaled, e = json.Marshal(r.required); e != nil {
		return
	}

	if r.optional != nil {
		var oBody []byte
		if oBody, e = json.Marshal(r.optional); e != nil {
			return
		}
		marshaled = marshaled[:len(marshaled)-1]
		marshaled = append(marshaled, []byte(",")...)
		oBody = oBody[1:]
		marshaled = append(marshaled, oBody...)
	}

	return
}
