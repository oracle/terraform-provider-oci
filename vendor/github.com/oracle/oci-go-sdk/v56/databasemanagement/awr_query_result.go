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

// AwrQueryResult The AWR query result.
type AwrQueryResult interface {

	// The name of the query result.
	GetName() *string

	// The version of the query result.
	GetVersion() *string

	// The ID assigned to the query instance.
	GetQueryKey() *string

	// The time taken to query the database tier (in seconds).
	GetDbQueryTimeInSecs() *float64
}

type awrqueryresult struct {
	JsonData          []byte
	Name              *string  `mandatory:"true" json:"name"`
	Version           *string  `mandatory:"false" json:"version"`
	QueryKey          *string  `mandatory:"false" json:"queryKey"`
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`
	AwrResultType     string   `json:"awrResultType"`
}

// UnmarshalJSON unmarshals json
func (m *awrqueryresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerawrqueryresult awrqueryresult
	s := struct {
		Model Unmarshalerawrqueryresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Version = s.Model.Version
	m.QueryKey = s.Model.QueryKey
	m.DbQueryTimeInSecs = s.Model.DbQueryTimeInSecs
	m.AwrResultType = s.Model.AwrResultType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *awrqueryresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AwrResultType {
	case "AWRDB_DB_PARAMETER_CHANGE":
		mm := AwrDbParameterChangeCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_ASH_CPU_USAGE_SET":
		mm := AwrDbCpuUsageCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_EVENT_HISTOGRAM_SET":
		mm := AwrDbWaitEventBucketCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_DB_PARAMETER_SET":
		mm := AwrDbParameterCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SYSSTAT_SET":
		mm := AwrDbSysstatCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_TOP_EVENT_SET":
		mm := AwrDbTopWaitEventCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SNAPSHOT_SET":
		mm := AwrDbSnapshotCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SET":
		mm := AwrDbCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SNAPSHOT_RANGE_SET":
		mm := AwrDbSnapshotRangeCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_DB_REPORT":
		mm := AwrDbReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_METRICS_SET":
		mm := AwrDbMetricCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_EVENT_SET":
		mm := AwrDbWaitEventCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SQL_REPORT":
		mm := AwrDbSqlReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m awrqueryresult) GetName() *string {
	return m.Name
}

//GetVersion returns Version
func (m awrqueryresult) GetVersion() *string {
	return m.Version
}

//GetQueryKey returns QueryKey
func (m awrqueryresult) GetQueryKey() *string {
	return m.QueryKey
}

//GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m awrqueryresult) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

func (m awrqueryresult) String() string {
	return common.PointerString(m)
}

// AwrQueryResultAwrResultTypeEnum Enum with underlying type: string
type AwrQueryResultAwrResultTypeEnum string

// Set of constants representing the allowable values for AwrQueryResultAwrResultTypeEnum
const (
	AwrQueryResultAwrResultTypeSet               AwrQueryResultAwrResultTypeEnum = "AWRDB_SET"
	AwrQueryResultAwrResultTypeSnapshotRangeSet  AwrQueryResultAwrResultTypeEnum = "AWRDB_SNAPSHOT_RANGE_SET"
	AwrQueryResultAwrResultTypeSnapshotSet       AwrQueryResultAwrResultTypeEnum = "AWRDB_SNAPSHOT_SET"
	AwrQueryResultAwrResultTypeMetricsSet        AwrQueryResultAwrResultTypeEnum = "AWRDB_METRICS_SET"
	AwrQueryResultAwrResultTypeSysstatSet        AwrQueryResultAwrResultTypeEnum = "AWRDB_SYSSTAT_SET"
	AwrQueryResultAwrResultTypeTopEventSet       AwrQueryResultAwrResultTypeEnum = "AWRDB_TOP_EVENT_SET"
	AwrQueryResultAwrResultTypeEventSet          AwrQueryResultAwrResultTypeEnum = "AWRDB_EVENT_SET"
	AwrQueryResultAwrResultTypeEventHistogram    AwrQueryResultAwrResultTypeEnum = "AWRDB_EVENT_HISTOGRAM"
	AwrQueryResultAwrResultTypeDbParameterSet    AwrQueryResultAwrResultTypeEnum = "AWRDB_DB_PARAMETER_SET"
	AwrQueryResultAwrResultTypeDbParameterChange AwrQueryResultAwrResultTypeEnum = "AWRDB_DB_PARAMETER_CHANGE"
	AwrQueryResultAwrResultTypeAshCpuUsageSet    AwrQueryResultAwrResultTypeEnum = "AWRDB_ASH_CPU_USAGE_SET"
	AwrQueryResultAwrResultTypeDbReport          AwrQueryResultAwrResultTypeEnum = "AWRDB_DB_REPORT"
	AwrQueryResultAwrResultTypeSqlReport         AwrQueryResultAwrResultTypeEnum = "AWRDB_SQL_REPORT"
)

var mappingAwrQueryResultAwrResultType = map[string]AwrQueryResultAwrResultTypeEnum{
	"AWRDB_SET":                 AwrQueryResultAwrResultTypeSet,
	"AWRDB_SNAPSHOT_RANGE_SET":  AwrQueryResultAwrResultTypeSnapshotRangeSet,
	"AWRDB_SNAPSHOT_SET":        AwrQueryResultAwrResultTypeSnapshotSet,
	"AWRDB_METRICS_SET":         AwrQueryResultAwrResultTypeMetricsSet,
	"AWRDB_SYSSTAT_SET":         AwrQueryResultAwrResultTypeSysstatSet,
	"AWRDB_TOP_EVENT_SET":       AwrQueryResultAwrResultTypeTopEventSet,
	"AWRDB_EVENT_SET":           AwrQueryResultAwrResultTypeEventSet,
	"AWRDB_EVENT_HISTOGRAM":     AwrQueryResultAwrResultTypeEventHistogram,
	"AWRDB_DB_PARAMETER_SET":    AwrQueryResultAwrResultTypeDbParameterSet,
	"AWRDB_DB_PARAMETER_CHANGE": AwrQueryResultAwrResultTypeDbParameterChange,
	"AWRDB_ASH_CPU_USAGE_SET":   AwrQueryResultAwrResultTypeAshCpuUsageSet,
	"AWRDB_DB_REPORT":           AwrQueryResultAwrResultTypeDbReport,
	"AWRDB_SQL_REPORT":          AwrQueryResultAwrResultTypeSqlReport,
}

// GetAwrQueryResultAwrResultTypeEnumValues Enumerates the set of values for AwrQueryResultAwrResultTypeEnum
func GetAwrQueryResultAwrResultTypeEnumValues() []AwrQueryResultAwrResultTypeEnum {
	values := make([]AwrQueryResultAwrResultTypeEnum, 0)
	for _, v := range mappingAwrQueryResultAwrResultType {
		values = append(values, v)
	}
	return values
}
