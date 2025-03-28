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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostResourceAllocation Resource Allocation metric for the host
type HostResourceAllocation struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Total number of CPUs available
	TotalCpus *int `mandatory:"false" json:"totalCpus"`

	// Total amount of usable physical memory in gibabytes
	TotalMemoryInGB *float64 `mandatory:"false" json:"totalMemoryInGB"`
}

// GetTimeCollected returns TimeCollected
func (m HostResourceAllocation) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostResourceAllocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostResourceAllocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostResourceAllocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostResourceAllocation HostResourceAllocation
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostResourceAllocation
	}{
		"HOST_RESOURCE_ALLOCATION",
		(MarshalTypeHostResourceAllocation)(m),
	}

	return json.Marshal(&s)
}
