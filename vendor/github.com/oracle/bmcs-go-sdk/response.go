// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type response struct {
	header http.Header
	body   []byte
}

func (r *response) unmarshal(resource interface{}) (e error) {
	var val interface{}

	if c, ok := resource.(Container); ok {
		val = c.GetList()
	} else {
		val = resource
	}

	if pc, ok := resource.(NextPageUnmarshallable); ok {
		pc.SetNextPage(r.header.Get(headerOPCNextPage))
	}

	if cs, ok := resource.(BodyUnmarshallable); ok {
		if e = cs.SetBody(r.body, val); e != nil {
			return
		}
	} else if len(r.body) == 0 {
		// Continue without error. This is usually caused by a 204 response
	} else if e = json.Unmarshal(r.body, val); e != nil {
		return
	}

	if rr, ok := resource.(OPCRequestIDUnmarshallable); ok {
		rr.SetRequestID(r.header.Get(headerOPCRequestID))
	}

	if crr, ok := resource.(OPCClientRequestIDUnmarshallable); ok {
		crr.SetClientRequestID(r.header.Get(headerOPCClientRequestID))
	}

	if wrr, ok := resource.(OPCWorkRequestIDUnmarshallable); ok {
		wrr.SetWorkRequestID(r.header.Get(headerOPCWorkRequestID))
	}

	if et, ok := resource.(ETagUnmarshallable); ok {
		et.SetETag(r.header.Get(headerETag))
	}

	// TODO: Unmarshal this string into a Time.
	// if et, ok := resource.(LastModifiedUnmarshallable); ok {
	// 	et.SetLastModified(r.header.Get(headerLastModified))
	// }

	if cr, ok := resource.(ContentUnmarshallable); ok {
		cr.SetContentEncoding(r.header.Get(headerContentEncoding))
		cr.SetContentLanguage(r.header.Get(headerContentLanguage))
		cr.SetContentMD5(r.header.Get(headerContentMD5))
		cr.SetContentType(r.header.Get(headerContentType))

		lengthStr := r.header.Get(headerContentLength)
		if lengthStr != "" {
			if length, err := strconv.Atoi(lengthStr); err != nil {
				e = err
				return
			} else {
				cr.SetContentLength(uint64(length))
			}
		}
	}

	if md, ok := resource.(MetadataUnmarshallable); ok {
		prefix := "opc-meta-"
		meta := make(map[string]string)
		for name, headers := range r.header {
			name = strings.ToLower(name)
			if strings.HasPrefix(name, prefix) {
				for _, h := range headers {
					meta[strings.Replace(name, prefix, "", 1)] = h
				}
			}
		}
		md.SetMetadata(meta)
	}

	return
}
