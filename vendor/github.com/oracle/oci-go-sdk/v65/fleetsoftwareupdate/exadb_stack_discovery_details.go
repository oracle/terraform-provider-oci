// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ExadbStackDiscoveryDetails Details to create an 'EXADB_STACK' type Exadata Fleet Update Discovery.
// Currently, components allowed in an Exadata software stack are 'GUEST_OS' and 'GI'.
// At least two distinct component types are required for an Exadata software stack.
type ExadbStackDiscoveryDetails struct {

	// Discovery filter details of components in an Exadata software stack.
	Components []SoftwareComponentDiscoveryDetails `mandatory:"true" json:"components"`

	// Exadata service type for the target resource members.
	ServiceType DiscoveryServiceTypesEnum `mandatory:"true" json:"serviceType"`
}

// GetServiceType returns ServiceType
func (m ExadbStackDiscoveryDetails) GetServiceType() DiscoveryServiceTypesEnum {
	return m.ServiceType
}

func (m ExadbStackDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbStackDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveryServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetDiscoveryServiceTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExadbStackDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadbStackDiscoveryDetails ExadbStackDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExadbStackDiscoveryDetails
	}{
		"EXADB_STACK",
		(MarshalTypeExadbStackDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExadbStackDiscoveryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ServiceType DiscoveryServiceTypesEnum           `json:"serviceType"`
		Components  []softwarecomponentdiscoverydetails `json:"components"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ServiceType = model.ServiceType

	m.Components = make([]SoftwareComponentDiscoveryDetails, len(model.Components))
	for i, n := range model.Components {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Components[i] = nn.(SoftwareComponentDiscoveryDetails)
		} else {
			m.Components[i] = nil
		}
	}
	return
}
