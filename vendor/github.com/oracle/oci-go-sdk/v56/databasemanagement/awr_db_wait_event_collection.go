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

// AwrDbWaitEventCollection The AWR wait event data.
type AwrDbWaitEventCollection struct {

	// The name of the query result.
	Name *string `mandatory:"true" json:"name"`

	// The version of the query result.
	Version *string `mandatory:"false" json:"version"`

	// The ID assigned to the query instance.
	QueryKey *string `mandatory:"false" json:"queryKey"`

	// The time taken to query the database tier (in seconds).
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`

	// A list of AWR wait events.
	Items []AwrDbWaitEventSummary `mandatory:"false" json:"items"`
}

//GetName returns Name
func (m AwrDbWaitEventCollection) GetName() *string {
	return m.Name
}

//GetVersion returns Version
func (m AwrDbWaitEventCollection) GetVersion() *string {
	return m.Version
}

//GetQueryKey returns QueryKey
func (m AwrDbWaitEventCollection) GetQueryKey() *string {
	return m.QueryKey
}

//GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m AwrDbWaitEventCollection) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m AwrDbWaitEventCollection) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AwrDbWaitEventCollection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwrDbWaitEventCollection AwrDbWaitEventCollection
	s := struct {
		DiscriminatorParam string `json:"awrResultType"`
		MarshalTypeAwrDbWaitEventCollection
	}{
		"AWRDB_EVENT_SET",
		(MarshalTypeAwrDbWaitEventCollection)(m),
	}

	return json.Marshal(&s)
}
