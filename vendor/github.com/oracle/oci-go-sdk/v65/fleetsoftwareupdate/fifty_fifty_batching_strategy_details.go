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

// FiftyFiftyBatchingStrategyDetails Fifty-Fifty batching strategy details to use during PRECHECK and APPLY Cycle Actions.
type FiftyFiftyBatchingStrategyDetails struct {

	// True to wait for customer to resume the Apply Action once the first half is done.
	// False to automatically patch the second half.
	IsWaitForBatchResume *bool `mandatory:"false" json:"isWaitForBatchResume"`

	// True to force rolling patching.
	IsForceRolling *bool `mandatory:"false" json:"isForceRolling"`
}

func (m FiftyFiftyBatchingStrategyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FiftyFiftyBatchingStrategyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FiftyFiftyBatchingStrategyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFiftyFiftyBatchingStrategyDetails FiftyFiftyBatchingStrategyDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFiftyFiftyBatchingStrategyDetails
	}{
		"FIFTY_FIFTY",
		(MarshalTypeFiftyFiftyBatchingStrategyDetails)(m),
	}

	return json.Marshal(&s)
}
