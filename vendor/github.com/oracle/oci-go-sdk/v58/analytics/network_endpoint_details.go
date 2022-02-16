// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// NetworkEndpointDetails Base representation of a network endpoint.
type NetworkEndpointDetails interface {
}

type networkendpointdetails struct {
	JsonData            []byte
	NetworkEndpointType string `json:"networkEndpointType"`
}

// UnmarshalJSON unmarshals json
func (m *networkendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernetworkendpointdetails networkendpointdetails
	s := struct {
		Model Unmarshalernetworkendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.NetworkEndpointType = s.Model.NetworkEndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *networkendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.NetworkEndpointType {
	case "PRIVATE":
		mm := PrivateEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PUBLIC":
		mm := PublicEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m networkendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m networkendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
