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

// AwrQueryResult The AWR query result.
type AwrQueryResult interface {

	// The name of the query result.
	GetName() *string

	// The version of the query result.
	GetVersion() *string

	// The time taken to query the database tier (in seconds).
	GetDbQueryTimeInSecs() *float64
}

type awrqueryresult struct {
	JsonData          []byte
	Version           *string  `mandatory:"false" json:"version"`
	DbQueryTimeInSecs *float64 `mandatory:"false" json:"dbQueryTimeInSecs"`
	Name              *string  `mandatory:"true" json:"name"`
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
	case "AWRDB_ASH_CPU_USAGE_SET":
		mm := AwrDatabaseCpuUsageCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_DB_PARAMETER_SET":
		mm := AwrDatabaseParameterCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_EVENT_HISTOGRAM_SET":
		mm := AwrDatabaseWaitEventBucketCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SNAPSHOT_RANGE_SET":
		mm := AwrDatabaseSnapshotRangeCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SNAPSHOT_SET":
		mm := AwrDatabaseSnapshotCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SYSSTAT_SET":
		mm := AwrDatabaseSysstatCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_METRICS_SET":
		mm := AwrDatabaseMetricCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_EVENT_SET":
		mm := AwrDatabaseWaitEventCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SET":
		mm := AwrDatabaseCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_TOP_EVENT_SET":
		mm := AwrDatabaseTopWaitEventCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_DB_PARAMETER_CHANGE":
		mm := AwrDatabaseParameterChangeCollection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_DB_REPORT":
		mm := AwrDatabaseReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWRDB_SQL_REPORT":
		mm := AwrDatabaseSqlReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AwrQueryResult: %s.", m.AwrResultType)
		return *m, nil
	}
}

// GetVersion returns Version
func (m awrqueryresult) GetVersion() *string {
	return m.Version
}

// GetDbQueryTimeInSecs returns DbQueryTimeInSecs
func (m awrqueryresult) GetDbQueryTimeInSecs() *float64 {
	return m.DbQueryTimeInSecs
}

// GetName returns Name
func (m awrqueryresult) GetName() *string {
	return m.Name
}

func (m awrqueryresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m awrqueryresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingAwrQueryResultAwrResultTypeEnum = map[string]AwrQueryResultAwrResultTypeEnum{
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

var mappingAwrQueryResultAwrResultTypeEnumLowerCase = map[string]AwrQueryResultAwrResultTypeEnum{
	"awrdb_set":                 AwrQueryResultAwrResultTypeSet,
	"awrdb_snapshot_range_set":  AwrQueryResultAwrResultTypeSnapshotRangeSet,
	"awrdb_snapshot_set":        AwrQueryResultAwrResultTypeSnapshotSet,
	"awrdb_metrics_set":         AwrQueryResultAwrResultTypeMetricsSet,
	"awrdb_sysstat_set":         AwrQueryResultAwrResultTypeSysstatSet,
	"awrdb_top_event_set":       AwrQueryResultAwrResultTypeTopEventSet,
	"awrdb_event_set":           AwrQueryResultAwrResultTypeEventSet,
	"awrdb_event_histogram":     AwrQueryResultAwrResultTypeEventHistogram,
	"awrdb_db_parameter_set":    AwrQueryResultAwrResultTypeDbParameterSet,
	"awrdb_db_parameter_change": AwrQueryResultAwrResultTypeDbParameterChange,
	"awrdb_ash_cpu_usage_set":   AwrQueryResultAwrResultTypeAshCpuUsageSet,
	"awrdb_db_report":           AwrQueryResultAwrResultTypeDbReport,
	"awrdb_sql_report":          AwrQueryResultAwrResultTypeSqlReport,
}

// GetAwrQueryResultAwrResultTypeEnumValues Enumerates the set of values for AwrQueryResultAwrResultTypeEnum
func GetAwrQueryResultAwrResultTypeEnumValues() []AwrQueryResultAwrResultTypeEnum {
	values := make([]AwrQueryResultAwrResultTypeEnum, 0)
	for _, v := range mappingAwrQueryResultAwrResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrQueryResultAwrResultTypeEnumStringValues Enumerates the set of values in String for AwrQueryResultAwrResultTypeEnum
func GetAwrQueryResultAwrResultTypeEnumStringValues() []string {
	return []string{
		"AWRDB_SET",
		"AWRDB_SNAPSHOT_RANGE_SET",
		"AWRDB_SNAPSHOT_SET",
		"AWRDB_METRICS_SET",
		"AWRDB_SYSSTAT_SET",
		"AWRDB_TOP_EVENT_SET",
		"AWRDB_EVENT_SET",
		"AWRDB_EVENT_HISTOGRAM",
		"AWRDB_DB_PARAMETER_SET",
		"AWRDB_DB_PARAMETER_CHANGE",
		"AWRDB_ASH_CPU_USAGE_SET",
		"AWRDB_DB_REPORT",
		"AWRDB_SQL_REPORT",
	}
}

// GetMappingAwrQueryResultAwrResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrQueryResultAwrResultTypeEnum(val string) (AwrQueryResultAwrResultTypeEnum, bool) {
	enum, ok := mappingAwrQueryResultAwrResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
