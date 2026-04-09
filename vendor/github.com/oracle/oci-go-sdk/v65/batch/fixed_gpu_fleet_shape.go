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

// FixedGpuFleetShape Fixed shape of the GPU fleet. Describes hardware resources of each node in the fleet.
type FixedGpuFleetShape struct {

	// The name of the fixed GPU fleet shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// Amount of disk space in GBs required for the shape.
	DiskSizeInGBs *int `mandatory:"false" json:"diskSizeInGBs"`
}

func (m FixedGpuFleetShape) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FixedGpuFleetShape) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FixedGpuFleetShape) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFixedGpuFleetShape FixedGpuFleetShape
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFixedGpuFleetShape
	}{
		"FIXED_GPU_FLEET_SHAPE",
		(MarshalTypeFixedGpuFleetShape)(m),
	}

	return json.Marshal(&s)
}
