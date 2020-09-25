// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// AddHttpRequestHeaderRule An object that represents the action of adding a header to a request. Optionally rule
// conditions can be specified to add header conditionally.`SOURCE_IP_ADDRESS` and `REAL_IP_ADDRESS` are the only rule condition supported.
// This rule applies only to HTTP listeners.
// **NOTES:**
// *  If a matching header already exists in the request, the system removes all of its occurrences, and then adds the
//    new header.
// * If a customer adds empty value, it has the same effect as dropping that header.
// *  The system does not distinquish between underscore and dash characters in headers. That is, it treats
//   `example_header_name` and `example-header-name` as identical. Oracle recommends that you do not rely on underscore
//   or dash characters to uniquely distinguish header names.
type AddHttpRequestHeaderRule struct {

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	Header *string `mandatory:"true" json:"header"`

	// A header value that conforms to RFC 7230.
	// Example: `example_value`
	Value *string `mandatory:"true" json:"value"`

	Conditions []RuleCondition `mandatory:"false" json:"conditions"`
}

func (m AddHttpRequestHeaderRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AddHttpRequestHeaderRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddHttpRequestHeaderRule AddHttpRequestHeaderRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAddHttpRequestHeaderRule
	}{
		"ADD_HTTP_REQUEST_HEADER",
		(MarshalTypeAddHttpRequestHeaderRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AddHttpRequestHeaderRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Conditions []rulecondition `json:"conditions"`
		Header     *string         `json:"header"`
		Value      *string         `json:"value"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Conditions = make([]RuleCondition, len(model.Conditions))
	for i, n := range model.Conditions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Conditions[i] = nn.(RuleCondition)
		} else {
			m.Conditions[i] = nil
		}
	}

	m.Header = model.Header

	m.Value = model.Value

	return
}
