// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetShape Shape of the fleet. Describes hardware resources of each node in the fleet.
type FleetShape struct {

	// The name of the shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// Number of OCPUs required by the shape.
	Ocpus *int `mandatory:"true" json:"ocpus"`

	// Amount of memory in GBs required by the shape.
	MemoryInGBs *int `mandatory:"true" json:"memoryInGBs"`
}

func (m FleetShape) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetShape) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
