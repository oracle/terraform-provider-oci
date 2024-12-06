// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNetworkEndpointDetails Base representation of a network endpoint.
// In input payload to update an Visual Builder instance endpoint details, an empty payload will clear out any existing configuration for Public Visual Builder instance.
type UpdateNetworkEndpointDetails interface {
}

type updatenetworkendpointdetails struct {
	JsonData            []byte
	NetworkEndpointType string `json:"networkEndpointType"`
}

// UnmarshalJSON unmarshals json
func (m *updatenetworkendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatenetworkendpointdetails updatenetworkendpointdetails
	s := struct {
		Model Unmarshalerupdatenetworkendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkEndpointType = s.Model.NetworkEndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatenetworkendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkEndpointType {
	case "PRIVATE":
		mm := UpdatePrivateEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateNetworkEndpointDetails: %s.", m.NetworkEndpointType)
		return *m, nil
	}
}

func (m updatenetworkendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatenetworkendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
