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

// DbDiscoveryDetails 'DB' type Exadata Fleet Update Discovery details.
type DbDiscoveryDetails struct {
	Criteria DbFleetDiscoveryDetails `mandatory:"true" json:"criteria"`

	// Exadata service type for the target resource members.
	ServiceType DiscoveryServiceTypesEnum `mandatory:"true" json:"serviceType"`

	// Database Major Version of targets to be included in the Exadata Fleet Update Discovery results.
	// https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions
	// Only Database targets that match the version specified in this value would be added to the Exadata Fleet Update Discovery results.
	SourceMajorVersion DbSourceMajorVersionsEnum `mandatory:"true" json:"sourceMajorVersion"`
}

// GetServiceType returns ServiceType
func (m DbDiscoveryDetails) GetServiceType() DiscoveryServiceTypesEnum {
	return m.ServiceType
}

func (m DbDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveryServiceTypesEnum(string(m.ServiceType)); !ok && m.ServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceType: %s. Supported values are: %s.", m.ServiceType, strings.Join(GetDiscoveryServiceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSourceMajorVersionsEnum(string(m.SourceMajorVersion)); !ok && m.SourceMajorVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceMajorVersion: %s. Supported values are: %s.", m.SourceMajorVersion, strings.Join(GetDbSourceMajorVersionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbDiscoveryDetails DbDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDbDiscoveryDetails
	}{
		"DB",
		(MarshalTypeDbDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DbDiscoveryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ServiceType        DiscoveryServiceTypesEnum `json:"serviceType"`
		SourceMajorVersion DbSourceMajorVersionsEnum `json:"sourceMajorVersion"`
		Criteria           dbfleetdiscoverydetails   `json:"criteria"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ServiceType = model.ServiceType

	m.SourceMajorVersion = model.SourceMajorVersion

	nn, e = model.Criteria.UnmarshalPolymorphicJSON(model.Criteria.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Criteria = nn.(DbFleetDiscoveryDetails)
	} else {
		m.Criteria = nil
	}

	return
}
