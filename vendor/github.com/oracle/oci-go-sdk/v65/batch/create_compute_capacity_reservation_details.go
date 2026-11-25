// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateComputeCapacityReservationDetails Compute capacity reservations supplied for initial fleet provisioning.
type CreateComputeCapacityReservationDetails struct {

	// Compute capacity reservation OCIDs used for initial fleet provisioning.
	CapacityReservationIds []string `mandatory:"true" json:"capacityReservationIds"`
}

func (m CreateComputeCapacityReservationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeCapacityReservationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateComputeCapacityReservationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeCapacityReservationDetails CreateComputeCapacityReservationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateComputeCapacityReservationDetails
	}{
		"COMPUTE_CAPACITY_RESERVATION",
		(MarshalTypeCreateComputeCapacityReservationDetails)(m),
	}

	return json.Marshal(&s)
}
