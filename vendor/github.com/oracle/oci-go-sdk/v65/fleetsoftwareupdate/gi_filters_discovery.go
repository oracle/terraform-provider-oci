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

// GiFiltersDiscovery Collection discovery done from the results of the specified filters.
type GiFiltersDiscovery struct {

	// Filters to perform the target discovery.
	Filters []GiFleetDiscoveryFilter `mandatory:"true" json:"filters"`
}

func (m GiFiltersDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiFiltersDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GiFiltersDiscovery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGiFiltersDiscovery GiFiltersDiscovery
	s := struct {
		DiscriminatorParam string `json:"strategy"`
		MarshalTypeGiFiltersDiscovery
	}{
		"FILTERS",
		(MarshalTypeGiFiltersDiscovery)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GiFiltersDiscovery) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Filters []gifleetdiscoveryfilter `json:"filters"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Filters = make([]GiFleetDiscoveryFilter, len(model.Filters))
	for i, n := range model.Filters {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Filters[i] = nn.(GiFleetDiscoveryFilter)
		} else {
			m.Filters[i] = nil
		}
	}
	return
}
