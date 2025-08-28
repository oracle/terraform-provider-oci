// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmClusterSummary Partial information about the VM Cluster which includes name, memory allocated etc.
type VmClusterSummary struct {

	// The name of the VM Cluster.
	VmclusterName *string `mandatory:"true" json:"vmclusterName"`

	// The memory allocated on a VM Cluster.
	MemoryAllocatedInGBs *int `mandatory:"false" json:"memoryAllocatedInGBs"`

	// The CPU allocated on a VM Cluster.
	CpuAllocated *int `mandatory:"false" json:"cpuAllocated"`

	// The number of DB nodes on a VM Cluster.
	DbNodesCount *int `mandatory:"false" json:"dbNodesCount"`

	// The storage allocated on a VM Cluster.
	StorageAllocatedInGBs *int `mandatory:"false" json:"storageAllocatedInGBs"`

	// The OCID of the VM Cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
