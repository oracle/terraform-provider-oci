// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryDetails Discovery filter details for search.
type DiscoveryDetails interface {

	// Exadata service type for the target resource members.
	GetServiceType() DiscoveryServiceTypesEnum
}

type discoverydetails struct {
	JsonData    []byte
	ServiceType DiscoveryServiceTypesEnum `mandatory:"true" json:"serviceType"`
	Type        string                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *discoverydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdiscoverydetails discoverydetails
	s := struct {
		Model Unmarshalerdiscoverydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ServiceType = s.Model.ServiceType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *discoverydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "GI":
		mm := GiDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB":
		mm := DbDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DiscoveryDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetServiceType returns ServiceType
func (m discoverydetails) GetServiceType() DiscoveryServiceTypesEnum {
	return m.ServiceType
}

func (m discoverydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m discoverydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetDiscoveryServiceTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
