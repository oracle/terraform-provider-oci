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

// ContainerResourceConfig The resource configuration for a container. The resource configuration determines
// the amount of resources allocated to the container and the maximum allowed resources for a container.
type ContainerResourceConfig struct {

	// The maximum amount of CPUs that can be consumed by the container's process.
	// If you do not set a value, then the process
	// may use all available CPU resources on the container instance.
	// CPU usage is defined in terms of logical CPUs. This means that the
	// maximum possible value on an E3 ContainerInstance with 1 OCPU is 2.0.
	VcpusLimit *float32 `mandatory:"false" json:"vcpusLimit"`

	// The maximum amount of memory that can be consumed by the container's
	// process. If you do not set a value, then the process
	// may use all available memory on the instance.
	MemoryLimitInGBs *float32 `mandatory:"false" json:"memoryLimitInGBs"`
}

func (m ContainerResourceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerResourceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
