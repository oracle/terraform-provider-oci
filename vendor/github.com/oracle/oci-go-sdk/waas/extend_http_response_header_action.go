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

// ExtendHttpResponseHeaderAction An object that represents the action of adding a header field to a response.
// If the header with specified value already exists, nothing will be added.
// If the header exists with different value, additional header name:value pair will be added.
// Comma separated header values are not considered individually (instead as a single value) when adding a new header field.
type ExtendHttpResponseHeaderAction struct {

	// A header field name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A header field value that conforms to RFC 7230.
	// Example: `example_value`
	Value *string `mandatory:"true" json:"value"`
}

func (m ExtendHttpResponseHeaderAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExtendHttpResponseHeaderAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExtendHttpResponseHeaderAction ExtendHttpResponseHeaderAction
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeExtendHttpResponseHeaderAction
	}{
		"EXTEND_HTTP_RESPONSE_HEADER",
		(MarshalTypeExtendHttpResponseHeaderAction)(m),
	}

	return json.Marshal(&s)
}
