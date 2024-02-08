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

// AwrDatabaseCpuUsageCollection The AWR CPU usage data.
type AwrDatabaseCpuUsageCollection struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// The number of available CPU cores, which include subcores of multicore and single-core CPUs.
	NumCpuCores *int `mandatory:"false" json:"numCpuCores"`

	// The number of CPUs available for the database to use.
	DatabaseCpuCount *int `mandatory:"false" json:"databaseCpuCount"`

	// The number of available CPUs or processors.
	HostCpuCount *float64 `mandatory:"false" json:"hostCpuCount"`

	// A list of AWR CPU usage summary data.
	Items []AwrDatabaseCpuUsageSummary `mandatory:"false" json:"items"`
}

// GetName returns Name
func (m AwrDatabaseCpuUsageCollection) GetName() *string {
	return m.Name
}

// GetVersion returns Version
func (m AwrDatabaseCpuUsageCollection) GetVersion() *string {
	return m.Version
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDatabaseCpuUsageCollection) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDatabaseCpuUsageCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseCpuUsageCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwrDatabaseCpuUsageCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDatabaseCpuUsageCollection AwrDatabaseCpuUsageCollection
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDatabaseCpuUsageCollection
	}{
		"AWRDB_ASH_CPU_USAGE_SET",
		(MarshalTypeAwrDatabaseCpuUsageCollection)(m),
	}

	return json.Marshal(&s)
}
