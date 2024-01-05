// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmClusterSummary Partial information about the VM Cluster which includes name, memory allocated etc.
type VmClusterSummary struct {

	// The name of the vm cluster.
	VmclusterName *string `mandatory:"true" json:"vmclusterName"`

	// The memory allocated on a vm cluster.
	MemoryAllocatedInGBs *int `mandatory:"false" json:"memoryAllocatedInGBs"`

	// The cpu allocated on a vm cluster.
	CpuAllocated *int `mandatory:"false" json:"cpuAllocated"`

	// The number of DB nodes on a vm cluster.
	DbNodesCount *int `mandatory:"false" json:"dbNodesCount"`
}

func (m VmClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
