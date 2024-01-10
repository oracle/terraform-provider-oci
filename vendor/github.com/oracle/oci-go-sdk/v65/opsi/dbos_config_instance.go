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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbosConfigInstance Configuration parameters defined for external databases instance level.
type DbosConfigInstance struct {

	// Name of the database instance.
	InstanceName *string `mandatory:"true" json:"instanceName"`

	// Host name of the database instance.
	HostName *string `mandatory:"true" json:"hostName"`

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// Total number of CPUs available.
	NumCPUs *int `mandatory:"false" json:"numCPUs"`

	// Number of CPU cores available (includes subcores of multicore CPUs as well as single-core CPUs).
	NumCPUCores *int `mandatory:"false" json:"numCPUCores"`

	// Number of CPU Sockets available.
	NumCPUSockets *int `mandatory:"false" json:"numCPUSockets"`

	// Total number of bytes of physical memory.
	PhysicalMemoryBytes *float64 `mandatory:"false" json:"physicalMemoryBytes"`
}

// GetTimeCollected returns TimeCollected
func (m DbosConfigInstance) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m DbosConfigInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbosConfigInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbosConfigInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbosConfigInstance DbosConfigInstance
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeDbosConfigInstance
	}{
		"DB_OS_CONFIG_INSTANCE",
		(MarshalTypeDbosConfigInstance)(m),
	}

	return json.Marshal(&s)
}
