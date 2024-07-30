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

// UpdateServiceAvailabilityFactorBatchingStrategyDetails Service Availability Factor batching strategy details to use during PRECHECK and APPLY Cycle Actions.
type UpdateServiceAvailabilityFactorBatchingStrategyDetails struct {

	// Percentage of availability in the service during the Patch operation.
	Percentage *int `mandatory:"false" json:"percentage"`

	// True to force rolling patching.
	IsForceRolling *bool `mandatory:"false" json:"isForceRolling"`
}

func (m UpdateServiceAvailabilityFactorBatchingStrategyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateServiceAvailabilityFactorBatchingStrategyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateServiceAvailabilityFactorBatchingStrategyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateServiceAvailabilityFactorBatchingStrategyDetails UpdateServiceAvailabilityFactorBatchingStrategyDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateServiceAvailabilityFactorBatchingStrategyDetails
	}{
		"SERVICE_AVAILABILITY_FACTOR",
		(MarshalTypeUpdateServiceAvailabilityFactorBatchingStrategyDetails)(m),
	}

	return json.Marshal(&s)
}
