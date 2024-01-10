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

// HostStorageStatistics Contains storage statistics.
type HostStorageStatistics struct {

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE) for a set of databases.
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Resource utilization in percentage.
	UtilizationPercent *float64 `mandatory:"true" json:"utilizationPercent"`

	// Change in resource utilization in percentage
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`

	FilesystemAvailableInPercent *float64 `mandatory:"false" json:"filesystemAvailableInPercent"`
}

// GetUsage returns Usage
func (m HostStorageStatistics) GetUsage() *float64 {
	return m.Usage
}

// GetCapacity returns Capacity
func (m HostStorageStatistics) GetCapacity() *float64 {
	return m.Capacity
}

// GetUtilizationPercent returns UtilizationPercent
func (m HostStorageStatistics) GetUtilizationPercent() *float64 {
	return m.UtilizationPercent
}

// GetUsageChangePercent returns UsageChangePercent
func (m HostStorageStatistics) GetUsageChangePercent() *float64 {
	return m.UsageChangePercent
}

func (m HostStorageStatistics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostStorageStatistics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostStorageStatistics) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostStorageStatistics HostStorageStatistics
	s := struct {
		DiscriminatorParam string `json:"resourceName"`
		MarshalTypeHostStorageStatistics
	}{
		"HOST_STORAGE_STATISTICS",
		(MarshalTypeHostStorageStatistics)(m),
	}

	return json.Marshal(&s)
}
