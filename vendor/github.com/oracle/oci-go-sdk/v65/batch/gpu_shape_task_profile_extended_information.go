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

// GpuShapeTaskProfileExtendedInformation Extended information for the GPU-specific task profile.
type GpuShapeTaskProfileExtendedInformation struct {

	// A name of the GPU shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`
}

func (m GpuShapeTaskProfileExtendedInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GpuShapeTaskProfileExtendedInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GpuShapeTaskProfileExtendedInformation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGpuShapeTaskProfileExtendedInformation GpuShapeTaskProfileExtendedInformation
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGpuShapeTaskProfileExtendedInformation
	}{
		"GPU_SHAPE_TASK_PROFILE_EXTENDED_INFORMATION",
		(MarshalTypeGpuShapeTaskProfileExtendedInformation)(m),
	}

	return json.Marshal(&s)
}
