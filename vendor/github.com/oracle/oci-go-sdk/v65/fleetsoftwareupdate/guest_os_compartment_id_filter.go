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

// GuestOsCompartmentIdFilter List of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartments to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
type GuestOsCompartmentIdFilter struct {

	// List of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Compartments to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	Identifiers []string `mandatory:"true" json:"identifiers"`

	// INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	// Supported only for RESOURCE_ID filter.
	Mode GuestOsFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
}

// GetMode returns Mode
func (m GuestOsCompartmentIdFilter) GetMode() GuestOsFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GuestOsCompartmentIdFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsCompartmentIdFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGuestOsFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGuestOsFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsCompartmentIdFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsCompartmentIdFilter GuestOsCompartmentIdFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGuestOsCompartmentIdFilter
	}{
		"COMPARTMENT_ID",
		(MarshalTypeGuestOsCompartmentIdFilter)(m),
	}

	return json.Marshal(&s)
}
