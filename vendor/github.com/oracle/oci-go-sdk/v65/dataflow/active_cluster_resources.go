// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActiveClusterResources Specifies the currently allocated resources for the driver and executors, including autoscaled workers.
type ActiveClusterResources struct {

	// The number of active executors currently running in the cluster.
	ActiveExecutorCount *float64 `mandatory:"false" json:"activeExecutorCount"`

	// The total number of CPU cores currently allocated to the driver and executors.
	ActiveCores *float64 `mandatory:"false" json:"activeCores"`

	// The total number of GPU cores currently allocated to the driver and executors.
	ActiveGpuCores *float64 `mandatory:"false" json:"activeGpuCores"`

	// The total amount of system memory (in GB) currently allocated to the driver and executors.
	ActiveMemoryInGB *float64 `mandatory:"false" json:"activeMemoryInGB"`

	// The total amount of GPU memory (in GB) currently allocated to the driver and executors.
	ActiveGpuMemoryInGB *float64 `mandatory:"false" json:"activeGpuMemoryInGB"`
}

func (m ActiveClusterResources) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActiveClusterResources) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
