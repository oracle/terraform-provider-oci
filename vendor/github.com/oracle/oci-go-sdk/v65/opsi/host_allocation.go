// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostAllocation Resource Allocation metric for the host
type HostAllocation struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// Name of the host resource
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Value of the host resource
	ResourceValue *int64 `mandatory:"false" json:"resourceValue"`
}

// GetTimeCollected returns TimeCollected
func (m HostAllocation) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostAllocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostAllocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostAllocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostAllocation HostAllocation
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostAllocation
	}{
		"HOST_RESOURCE_ALLOCATION",
		(MarshalTypeHostAllocation)(m),
	}

	return json.Marshal(&s)
}
