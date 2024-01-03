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

// ContainerInstanceShapeConfig The shape configuration for a container instance. The shape configuration determines
// the resources thats are available to the container instance and its containers.
type ContainerInstanceShapeConfig struct {

	// The total number of OCPUs available to the container instance.
	Ocpus *float32 `mandatory:"true" json:"ocpus"`

	// The total amount of memory available to the container instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"true" json:"memoryInGBs"`

	// A short description of the container instance's processor (CPU).
	ProcessorDescription *string `mandatory:"true" json:"processorDescription"`

	// The networking bandwidth available to the container instance, in gigabits per second.
	NetworkingBandwidthInGbps *float32 `mandatory:"true" json:"networkingBandwidthInGbps"`
}

func (m ContainerInstanceShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerInstanceShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
