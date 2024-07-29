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

// GiDefinedTagsFilter Defined tags to include in the discovery.
type GiDefinedTagsFilter struct {

	// Defined tags to include in the discovery.
	Tags []DefinedTagFilterEntry `mandatory:"true" json:"tags"`

	// INCLUDE or EXCLUDE the filter results in the discovery for GI targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	Mode GiFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`

	// Type of join for each element in this filter.
	Operator FleetDiscoveryOperatorsEnum `mandatory:"false" json:"operator,omitempty"`
}

// GetMode returns Mode
func (m GiDefinedTagsFilter) GetMode() GiFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GiDefinedTagsFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiDefinedTagsFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGiFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGiFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetDiscoveryOperatorsEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetFleetDiscoveryOperatorsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiDefinedTagsFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiDefinedTagsFilter GiDefinedTagsFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGiDefinedTagsFilter
	}{
		"DEFINED_TAG",
		(MarshalTypeGiDefinedTagsFilter)(m),
	}

	return json.Marshal(&s)
}
