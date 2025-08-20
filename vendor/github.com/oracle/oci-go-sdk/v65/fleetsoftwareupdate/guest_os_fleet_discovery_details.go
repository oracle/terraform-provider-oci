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

// GuestOsFleetDiscoveryDetails Fleet discovery strategies for a 'GUEST_OS' collection of Exadata VM Clusters.
// If specified for an UpdateCollection request, discovery for Exadata VM Clusters will be rerun.
type GuestOsFleetDiscoveryDetails interface {
}

type guestosfleetdiscoverydetails struct {
	JsonData []byte
	Strategy string `json:"strategy"`
}

// UnmarshalJSON unmarshals json
func (m *guestosfleetdiscoverydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerguestosfleetdiscoverydetails guestosfleetdiscoverydetails
	s := struct {
		Model Unmarshalerguestosfleetdiscoverydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Strategy = s.Model.Strategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *guestosfleetdiscoverydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Strategy {
	case "TARGET_LIST":
		mm := GuestOsTargetListDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH_QUERY":
		mm := GuestOsSearchQueryDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISCOVERY_RESULTS":
		mm := GuestOsDiscoveryResults{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILTERS":
		mm := GuestOsFiltersDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GuestOsFleetDiscoveryDetails: %s.", m.Strategy)
		return *m, nil
	}
}

func (m guestosfleetdiscoverydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m guestosfleetdiscoverydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
