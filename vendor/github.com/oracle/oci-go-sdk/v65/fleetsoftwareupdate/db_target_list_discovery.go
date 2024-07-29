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

// DbTargetListDiscovery Collection discovery conformed by the specified list of targets.
type DbTargetListDiscovery struct {

	// OCIDs of target database resources to include.
	Targets []string `mandatory:"true" json:"targets"`
}

func (m DbTargetListDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbTargetListDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbTargetListDiscovery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbTargetListDiscovery DbTargetListDiscovery
	s := struct {
		DiscriminatorParam string `json:"strategy"`
		MarshalTypeDbTargetListDiscovery
	}{
		"TARGET_LIST",
		(MarshalTypeDbTargetListDiscovery)(m),
	}

	return json.Marshal(&s)
}
