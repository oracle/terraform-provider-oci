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

// CpuFleetShapeExecutionDetails Details about the CPU shape which was used for the task execution.
type CpuFleetShapeExecutionDetails struct {

	// Name of the shape.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// Number of OCPUs provided by the shape.
	Ocpus *int `mandatory:"false" json:"ocpus"`

	// Amount of memory in GBs provided by the shape.
	MemoryInGBs *int `mandatory:"false" json:"memoryInGBs"`

	// Amount of disk space provided by the shape.
	DiskSizeInGBs *int `mandatory:"false" json:"diskSizeInGBs"`
}

func (m CpuFleetShapeExecutionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CpuFleetShapeExecutionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CpuFleetShapeExecutionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCpuFleetShapeExecutionDetails CpuFleetShapeExecutionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCpuFleetShapeExecutionDetails
	}{
		"CPU_FLEET_SHAPE_EXECUTION_DETAILS",
		(MarshalTypeCpuFleetShapeExecutionDetails)(m),
	}

	return json.Marshal(&s)
}
