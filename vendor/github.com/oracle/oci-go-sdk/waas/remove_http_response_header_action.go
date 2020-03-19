// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// RemoveHttpResponseHeaderAction An object that represents the action of removing from a response all occurrences of header fields
// with a specified name.
type RemoveHttpResponseHeaderAction struct {

	// A header field name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`
}

func (m RemoveHttpResponseHeaderAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RemoveHttpResponseHeaderAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemoveHttpResponseHeaderAction RemoveHttpResponseHeaderAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRemoveHttpResponseHeaderAction
	}{
		"REMOVE_HTTP_RESPONSE_HEADER",
		(MarshalTypeRemoveHttpResponseHeaderAction)(m),
	}

	return json.Marshal(&s)
}
