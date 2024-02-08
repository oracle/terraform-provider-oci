// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerInstanceShapeSummary Details about a shape for a container instance.
type ContainerInstanceShapeSummary struct {

	// The name identifying the shape.
	Name *string `mandatory:"true" json:"name"`

	// A short description of the container instance's processor (CPU).
	ProcessorDescription *string `mandatory:"true" json:"processorDescription"`

	OcpuOptions *ShapeOcpuOptions `mandatory:"false" json:"ocpuOptions"`

	MemoryOptions *ShapeMemoryOptions `mandatory:"false" json:"memoryOptions"`

	NetworkingBandwidthOptions *ShapeNetworkingBandwidthOptions `mandatory:"false" json:"networkingBandwidthOptions"`
}

func (m ContainerInstanceShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerInstanceShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
