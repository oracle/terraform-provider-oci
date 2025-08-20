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

// GuestOsDefinedTagsFilter Defined tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
type GuestOsDefinedTagsFilter struct {

	// Defined tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	Tags []DefinedTagFilterEntry `mandatory:"true" json:"tags"`

	// INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.
	// Supported only for RESOURCE_ID filter.
	Mode GuestOsFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`

	// Type of join for each element in this filter.
	Operator FleetDiscoveryOperatorsEnum `mandatory:"false" json:"operator,omitempty"`
}

// GetMode returns Mode
func (m GuestOsDefinedTagsFilter) GetMode() GuestOsFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GuestOsDefinedTagsFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsDefinedTagsFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGuestOsFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGuestOsFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetDiscoveryOperatorsEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetFleetDiscoveryOperatorsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsDefinedTagsFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsDefinedTagsFilter GuestOsDefinedTagsFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGuestOsDefinedTagsFilter
	}{
		"DEFINED_TAG",
		(MarshalTypeGuestOsDefinedTagsFilter)(m),
	}

	return json.Marshal(&s)
}
