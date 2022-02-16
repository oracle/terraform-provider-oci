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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ChangeIntegrationInstanceNetworkEndpointDetails Input payload to update an Integration instance endpoint details. An empty payload will clear out any existing configuration.
type ChangeIntegrationInstanceNetworkEndpointDetails struct {
	NetworkEndpointDetails NetworkEndpointDetails `mandatory:"false" json:"networkEndpointDetails"`
}

func (m ChangeIntegrationInstanceNetworkEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeIntegrationInstanceNetworkEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
