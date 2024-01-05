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

// CreateContainerResourceConfigDetails The size and amount of resources available to the container.
type CreateContainerResourceConfigDetails struct {

	// The maximum amount of CPUs that can be consumed by the container's process.
	// If you do not set a value, then the process
	// can use all available CPU resources on the instance.
	// CPU usage is defined in terms of logical CPUs. This means that the maximum possible value on
	// an E3 ContainerInstance with 1 OCPU is 2.0.
	// A container with a 2.0 vcpusLimit could consume up to 100% of the CPU resources available on the container instance.
	// Values can be fractional. A value of "1.5" means that the container
	// can consume at most the equivalent of 1 and a half logical CPUs worth of CPU capacity.
	VcpusLimit *float32 `mandatory:"false" json:"vcpusLimit"`

	// The maximum amount of memory that can be consumed by the container's
	// process.
	// If you do not set a value, then the process
	// may use all available memory on the instance.
	MemoryLimitInGBs *float32 `mandatory:"false" json:"memoryLimitInGBs"`
}

func (m CreateContainerResourceConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerResourceConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
