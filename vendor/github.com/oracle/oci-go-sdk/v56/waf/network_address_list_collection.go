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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// NetworkAddressListCollection Contains NetworkAddressListSummary items.
type NetworkAddressListCollection struct {

	// A list of NetworkAddressListSummary objects.
	Items []NetworkAddressListSummary `mandatory:"true" json:"items"`
}

func (m NetworkAddressListCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *NetworkAddressListCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []networkaddresslistsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]NetworkAddressListSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(NetworkAddressListSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
