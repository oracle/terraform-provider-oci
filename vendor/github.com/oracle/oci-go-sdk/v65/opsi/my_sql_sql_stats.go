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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlSqlStats MySql Sql Stats type object.
type MySqlSqlStats struct {

	// Unique SQL ID Digest for a MySql Statement.
	// Example: `"c20fcea11911be36651b7ca7bd3712d4ed9ac1134cee9c6620039e1fb13b5eff"`
	Digest *string `mandatory:"true" json:"digest"`

	// Collection timestamp.
	// Example: `"2020-03-31T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Type of statement such as select, update or delete.
	CommandType *string `mandatory:"false" json:"commandType"`

	// Total number of SQL statements used in collection ranking calculation.
	TotalRows *int64 `mandatory:"false" json:"totalRows"`

	// Percent of SQL statements in the perf schema table relative to max or overflow count set in @@GLOBAL.performance_schema_digests_size.
	PerfSchemaUsedPercent *int64 `mandatory:"false" json:"perfSchemaUsedPercent"`

	// Name of Database Schema.
	// Example: `"performance_schema"`
	SchemaName *string `mandatory:"false" json:"schemaName"`

	// The total number of times the statement has executed.
	ExecCount *int64 `mandatory:"false" json:"execCount"`

	// The total wait time (in picoseconds) of timed occurrences of the statement.
	TotalLatencyInPs *int64 `mandatory:"false" json:"totalLatencyInPs"`

	// The total time waiting (in picoseconds) for locks by timed occurrences of the statement.
	LockLatencyInPs *int64 `mandatory:"false" json:"lockLatencyInPs"`

	// The total number of errors produced by occurrences of the statement.
	ErrCount *int64 `mandatory:"false" json:"errCount"`

	// The total number of warnings produced by occurrences of the statement.
	WarnCount *int64 `mandatory:"false" json:"warnCount"`

	// The total number of rows affected by occurrences of the statement.
	RowsAffected *int64 `mandatory:"false" json:"rowsAffected"`

	// The total number of rows returned by occurrences of the statement.
	RowsSent *int64 `mandatory:"false" json:"rowsSent"`

	// The total number of rows read from storage engines by occurrences of the statement.
	RowsExamined *int64 `mandatory:"false" json:"rowsExamined"`

	// The total number of internal on-disk temporary tables created by occurrences of the statement.
	TmpDiskTables *int64 `mandatory:"false" json:"tmpDiskTables"`

	// The total number of internal in-memory temporary tables created by occurrences of the statement Count
	TmpTables *int64 `mandatory:"false" json:"tmpTables"`

	// The total number of joins that perform table scans because they do not use indexes by occurrences of the statement. If this value is not 0
	SelectFullJoin *int64 `mandatory:"false" json:"selectFullJoin"`

	// The total number of joins that used a range search on a reference table by occurrences of the statement
	SelectFullRangeJoin *int64 `mandatory:"false" json:"selectFullRangeJoin"`

	// The total number of joins that used ranges on the first table by occurrences of the statement. This is normally not a critical issue even if the value is quite large. Count
	SelectRange *int64 `mandatory:"false" json:"selectRange"`

	// The total number of joins without keys that check for key usage after each row by occurrences of the statement. If this is not 0
	SelectRangeCheck *int64 `mandatory:"false" json:"selectRangeCheck"`

	// The total number of joins that did a full scan of the first table by occurrences of the statement Count
	SelectScan *int64 `mandatory:"false" json:"selectScan"`

	// The total number of sort merge passes by occurrences of the statement.
	SortMergePasses *int64 `mandatory:"false" json:"sortMergePasses"`

	// The total number of sorts that were done using ranges by occurrences of the statement.
	SortRange *int64 `mandatory:"false" json:"sortRange"`

	// The total number of rows sorted by occurrences of the statement.
	RowsSorted *int64 `mandatory:"false" json:"rowsSorted"`

	// The total number of sorts that were done by scanning the table by occurrences of the statement.
	SortScan *int64 `mandatory:"false" json:"sortScan"`

	// The number of occurences of the statement which performed a table scan without using an index Count
	NoIndexUsedCount *int64 `mandatory:"false" json:"noIndexUsedCount"`

	// The number of occurences of the statement where the server found no good index to use Count
	NoGoodIndexUsedCount *int64 `mandatory:"false" json:"noGoodIndexUsedCount"`

	// The total time spent on CPU (in picoseconds) for the current thread.
	CpuLatencyInPs *int64 `mandatory:"false" json:"cpuLatencyInPs"`

	// The maximum amount of controlled memory (in bytes) used by the statement.
	MaxControlledMemoryInBytes *int64 `mandatory:"false" json:"maxControlledMemoryInBytes"`

	// The maximum amount of memory (in bytes) used by the statement.
	MaxTotalMemoryInBytes *int64 `mandatory:"false" json:"maxTotalMemoryInBytes"`

	// The total number of times a query was processed on the secondary engine (HEATWAVE) for occurrences of this statement Count.
	ExecCountSecondary *int64 `mandatory:"false" json:"execCountSecondary"`

	// The time at which statement was first seen.
	// Example: `"2023-01-16 08:04:31.533577"`
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The time at which statement was most recently seen for all occurrences of the statement.
	// Example: `"2023-01-30 02:17:08.067961"`
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m MySqlSqlStats) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlSqlStats) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
