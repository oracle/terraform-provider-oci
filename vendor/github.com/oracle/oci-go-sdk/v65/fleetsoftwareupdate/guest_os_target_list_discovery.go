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

// GuestOsTargetListDiscovery Discover specified list of Exadata VM Cluster targets for a 'GUEST_OS' collection.
type GuestOsTargetListDiscovery struct {

	// The OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata VM Cluster targets.
	// Only Exadata VM Cluster targets associated with the specified 'serviceType' are allowed.
	Targets []string `mandatory:"true" json:"targets"`
}

func (m GuestOsTargetListDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuestOsTargetListDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GuestOsTargetListDiscovery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGuestOsTargetListDiscovery GuestOsTargetListDiscovery
	s := struct {
		DiscriminatorParam string `json:"strategy"`
		MarshalTypeGuestOsTargetListDiscovery
	}{
		"TARGET_LIST",
		(MarshalTypeGuestOsTargetListDiscovery)(m),
	}

	return json.Marshal(&s)
}
