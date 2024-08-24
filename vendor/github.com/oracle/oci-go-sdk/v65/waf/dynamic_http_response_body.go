// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicHttpResponseBody Allows returning dynamically generated HTTP response body based on the provided template.
// The template allows variable interpolation by specifying variable name between the '${' and '}' delimiters.
// Escape sequences using '\' are supported to allow usage of '\\' and '\${' in the template to return '\' and '\${' in final response.
// The following variables are supported:
// * http.request.id - the HTTP request ID. For example: "d5fa953f75ef417e4c8008ef9336d779".
// Example:
//
//	{
//	  "type": "DYNAMIC",
//	  "template": "{\n  \"code\": 403,\n  \"message\":\"Unauthorised\",\n  \"incidentId\": \"${http.request.id}\"\n}"
//	}
//
// Example with escape sequence:
//
//	{
//	  "type": "DYNAMIC",
//	  "template": "\\${Returned as plain text}"
//	}
type DynamicHttpResponseBody struct {

	// Dynamic response body
	Template *string `mandatory:"true" json:"template"`
}

func (m DynamicHttpResponseBody) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicHttpResponseBody) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DynamicHttpResponseBody) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicHttpResponseBody DynamicHttpResponseBody
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDynamicHttpResponseBody
	}{
		"DYNAMIC",
		(MarshalTypeDynamicHttpResponseBody)(m),
	}

	return json.Marshal(&s)
}
