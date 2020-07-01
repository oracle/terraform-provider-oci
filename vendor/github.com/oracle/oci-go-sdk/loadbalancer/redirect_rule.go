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
	"github.com/oracle/oci-go-sdk/common"
)

// RedirectRule An object that represents the action of returning a specified response code and a redirect URI. Each RedirectRule
// object is configured for a particular listener and a designated path.
// The default response code is `302 Found`.
// **NOTES:**
// *  This rule applies only to HTTP listeners.
// *  You can specify this rule only with the RuleCondition
//    type `PATH`.
// *  A listener can have only one RedirectRule object for a given original path. The
//   PathMatchCondition `attributeValue` specifies the
//   original path.
type RedirectRule struct {
	Conditions []RuleCondition `mandatory:"true" json:"conditions"`

	// The HTTP status code to return when the incoming request is redirected.
	// The status line returned with the code is mapped from the standard HTTP specification. Valid response
	// codes for redirection are:
	// *  301
	// *  302
	// *  303
	// *  307
	// *  308
	// The default value is `302` (Found).
	// Example: `301`
	ResponseCode *int `mandatory:"false" json:"responseCode"`

	RedirectUri *RedirectUri `mandatory:"false" json:"redirectUri"`
}

func (m RedirectRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RedirectRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRedirectRule RedirectRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeRedirectRule
	}{
		"REDIRECT",
		(MarshalTypeRedirectRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *RedirectRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ResponseCode *int            `json:"responseCode"`
		RedirectUri  *RedirectUri    `json:"redirectUri"`
		Conditions   []rulecondition `json:"conditions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ResponseCode = model.ResponseCode

	m.RedirectUri = model.RedirectUri

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

	return
}
