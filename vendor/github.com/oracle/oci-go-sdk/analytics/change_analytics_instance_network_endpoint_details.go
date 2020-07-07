// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeAnalyticsInstanceNetworkEndpointDetails Input payload to update an Analytics instance endpoint details.
type ChangeAnalyticsInstanceNetworkEndpointDetails struct {
	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"true" json:"networkEndpointDetails"`
}

func (m ChangeAnalyticsInstanceNetworkEndpointDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ChangeAnalyticsInstanceNetworkEndpointDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		NetworkEndpointDetails networkendpointdetails `json:"networkEndpointDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.NetworkEndpointDetails.UnmarshalPolymorphicJSON(model.NetworkEndpointDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkEndpointDetails = nn.(NetworkEndpointDetails)
	} else {
		m.NetworkEndpointDetails = nil
	}

	return
}
