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

// GiVersionFilter Versions to include in the discovery. These should be under the Source Major Version of the Collection.
type GiVersionFilter struct {

	// List of Versions strings to include in the discovery.
	Versions []string `mandatory:"true" json:"versions"`

	// INCLUDE or EXCLUDE the filter results in the discovery for GI targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	Mode GiFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
}

// GetMode returns Mode
func (m GiVersionFilter) GetMode() GiFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m GiVersionFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiVersionFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGiFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetGiFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiVersionFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiVersionFilter GiVersionFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGiVersionFilter
	}{
		"VERSION",
		(MarshalTypeGiVersionFilter)(m),
	}

	return json.Marshal(&s)
}
