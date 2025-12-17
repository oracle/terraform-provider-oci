// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// FlexFitFleetAssignmentPolicy Similar to best-fit, but with the ability for the system to use larger fleets if the smallest sufficient fleet
// is not available, up to a specified max shaped fleet, based on a user-configured threshold value.
type FlexFitFleetAssignmentPolicy struct {

	// Specifies how much larger a fleet's nodes (shape) can be to still be considered for a task.
	// If threshold is not supplied the task will use any sufficient node available regards to minimum hardware requirements.
	Threshold *float32 `mandatory:"false" json:"threshold"`
}

func (m FlexFitFleetAssignmentPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlexFitFleetAssignmentPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FlexFitFleetAssignmentPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFlexFitFleetAssignmentPolicy FlexFitFleetAssignmentPolicy
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFlexFitFleetAssignmentPolicy
	}{
		"FLEX_FIT",
		(MarshalTypeFlexFitFleetAssignmentPolicy)(m),
	}

	return json.Marshal(&s)
}
