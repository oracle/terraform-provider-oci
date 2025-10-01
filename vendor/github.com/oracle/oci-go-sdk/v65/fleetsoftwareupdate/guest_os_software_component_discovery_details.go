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

// GuestOsSoftwareComponentDiscoveryDetails Discovery filter details of 'GUEST_OS' component in an Exadata software stack.
type GuestOsSoftwareComponentDiscoveryDetails struct {
	Criteria GuestOsFleetDiscoveryDetails `mandatory:"true" json:"criteria"`

	// Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in the results for an Exadata Fleet Update Discovery.
	// Major Versions of Exadata Software are demarcated by the underlying Oracle Linux OS version.
	// For more details, refer to Oracle document 2075007.1 (https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html)
	SourceMajorVersion GuestOsSourceMajorVersionsEnum `mandatory:"true" json:"sourceMajorVersion"`
}

func (m GuestOsSoftwareComponentDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsSoftwareComponentDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGuestOsSourceMajorVersionsEnum(string(m.SourceMajorVersion)); !ok && m.SourceMajorVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceMajorVersion: %s. Supported values are: %s.", m.SourceMajorVersion, strings.Join(GetGuestOsSourceMajorVersionsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsSoftwareComponentDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsSoftwareComponentDiscoveryDetails GuestOsSoftwareComponentDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeGuestOsSoftwareComponentDiscoveryDetails
	}{
		"GUEST_OS",
		(MarshalTypeGuestOsSoftwareComponentDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GuestOsSoftwareComponentDiscoveryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SourceMajorVersion GuestOsSourceMajorVersionsEnum `json:"sourceMajorVersion"`
		Criteria           guestosfleetdiscoverydetails   `json:"criteria"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SourceMajorVersion = model.SourceMajorVersion

	nn, e = model.Criteria.UnmarshalPolymorphicJSON(model.Criteria.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Criteria = nn.(GuestOsFleetDiscoveryDetails)
	} else {
		m.Criteria = nil
	}

	return
}
