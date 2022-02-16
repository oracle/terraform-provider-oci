// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StaticTextHttpResponseBody Allows returning static text as HTTP response body.
// Example:
// {
//   "type": "STATIC_TEXT",
//   "text": "{\n  \"code\": 403,\n  \"message\":\"Unauthorised\"\n}"
// }
type StaticTextHttpResponseBody struct {

	// Static response body text.
	Text *string `mandatory:"true" json:"text"`
}

func (m StaticTextHttpResponseBody) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StaticTextHttpResponseBody) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StaticTextHttpResponseBody) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStaticTextHttpResponseBody StaticTextHttpResponseBody
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStaticTextHttpResponseBody
	}{
		"STATIC_TEXT",
		(MarshalTypeStaticTextHttpResponseBody)(m),
	}

	return json.Marshal(&s)
}
