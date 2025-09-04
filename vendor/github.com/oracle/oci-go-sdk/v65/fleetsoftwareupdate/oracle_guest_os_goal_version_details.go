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

// OracleGuestOsGoalVersionDetails Details of a goal Oracle released 'GI' software image.
type OracleGuestOsGoalVersionDetails struct {

	// Goal version string matching an Oracle released 'GUEST_OS' software image.
	GoalVersion *string `mandatory:"true" json:"goalVersion"`
}

func (m OracleGuestOsGoalVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleGuestOsGoalVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleGuestOsGoalVersionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleGuestOsGoalVersionDetails OracleGuestOsGoalVersionDetails
	s := struct {
		DiscriminatorParam string `json:"goalType"`
		MarshalTypeOracleGuestOsGoalVersionDetails
	}{
		"GUEST_OS_ORACLE_IMAGE",
		(MarshalTypeOracleGuestOsGoalVersionDetails)(m),
	}

	return json.Marshal(&s)
}
