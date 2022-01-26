// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AwrDbCpuUsageCollection The AWR CPU usage data.
type AwrDbCpuUsageCollection struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The ID assigned to the query instance.
	QueryKey *string `mandatory:"false" json:"queryKey"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// The number of available CPU cores, which include subcores of multicore and single-core CPUs.
	NumCpuCores *int `mandatory:"false" json:"numCpuCores"`

	// The number of CPUs available for the database to use.
	CpuCount *int `mandatory:"false" json:"cpuCount"`

	// The number of available CPUs or processors.
	NumCpus *float64 `mandatory:"false" json:"numCpus"`

	// A list of AWR CPU usage summary data.
	Items []AwrDbCpuUsageSummary `mandatory:"false" json:"items"`
}

//GetName returns Name
func (m AwrDbCpuUsageCollection) GetName() *string {
	return m.Name
}

//GetVersion returns Version
func (m AwrDbCpuUsageCollection) GetVersion() *string {
	return m.Version
}

//GetQueryKey returns QueryKey
func (m AwrDbCpuUsageCollection) GetQueryKey() *string {
	return m.QueryKey
}

//GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDbCpuUsageCollection) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDbCpuUsageCollection) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AwrDbCpuUsageCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDbCpuUsageCollection AwrDbCpuUsageCollection
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDbCpuUsageCollection
	}{
		"AWRDB_ASH_CPU_USAGE_SET",
		(MarshalTypeAwrDbCpuUsageCollection)(m),
	}

	return json.Marshal(&s)
}
