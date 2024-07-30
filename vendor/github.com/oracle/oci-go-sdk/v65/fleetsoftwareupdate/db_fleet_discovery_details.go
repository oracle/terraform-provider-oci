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

// DbFleetDiscoveryDetails Supported fleet discovery strategies for DB Collections.
// If specified on an Update Collection request, this will re-discover the targets of the Collection.
type DbFleetDiscoveryDetails interface {
}

type dbfleetdiscoverydetails struct {
	JsonData []byte
	Strategy string `json:"strategy"`
}

// UnmarshalJSON unmarshals json
func (m *dbfleetdiscoverydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbfleetdiscoverydetails dbfleetdiscoverydetails
	s := struct {
		Model Unmarshalerdbfleetdiscoverydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Strategy = s.Model.Strategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbfleetdiscoverydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Strategy {
	case "DISCOVERY_RESULTS":
		mm := DbDiscoveryResults{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH_QUERY":
		mm := DbSearchQueryDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET_LIST":
		mm := DbTargetListDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILTERS":
		mm := DbFiltersDiscovery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DbFleetDiscoveryDetails: %s.", m.Strategy)
		return *m, nil
	}
}

func (m dbfleetdiscoverydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbfleetdiscoverydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
