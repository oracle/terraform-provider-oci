// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

// AddHttpResponseHeaderAction An object that represents the action of replacing or adding a header field.
// All prior occurrences of the header with the given name are removed and then the header field with specified value is added.
type AddHttpResponseHeaderAction struct {

	// A header field name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A header field value that conforms to RFC 7230.
	// Example: `example_value`
	Value *string `mandatory:"true" json:"value"`
}

func (m AddHttpResponseHeaderAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AddHttpResponseHeaderAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddHttpResponseHeaderAction AddHttpResponseHeaderAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAddHttpResponseHeaderAction
	}{
		"ADD_HTTP_RESPONSE_HEADER",
		(MarshalTypeAddHttpResponseHeaderAction)(m),
	}

	return json.Marshal(&s)
}
