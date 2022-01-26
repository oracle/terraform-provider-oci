// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ChangeIntegrationInstanceNetworkEndpointDetails Input payload to update an Integration instance endpoint details. An empty payload will clear out any existing configuration.
type ChangeIntegrationInstanceNetworkEndpointDetails struct {
	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m ChangeIntegrationInstanceNetworkEndpointDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ChangeIntegrationInstanceNetworkEndpointDetails) UnmarshalJSON(data []byte) (e error) {
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
