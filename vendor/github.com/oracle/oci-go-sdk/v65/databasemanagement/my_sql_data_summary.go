// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlDataSummary The SQL performance data record for a specific SQL query.
type MySqlDataSummary struct {

	// The name of the default schema when executing the query. If a schema is not set as the default, then the value is NULL.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The digest information of the normalized query.
	Digest *string `mandatory:"true" json:"digest"`

	// The normalized query.
	DigestText *string `mandatory:"true" json:"digestText"`

	// The number Of times the query has been executed.
	CountStar *float32 `mandatory:"true" json:"countStar"`

	// The total amount of time that has been spent executing the query.
	SumTimerWait *float32 `mandatory:"true" json:"sumTimerWait"`

	// The fastest the query has been executed.
	MinTimerWait *float32 `mandatory:"true" json:"minTimerWait"`

	// The average execution time.
	AvgTimerWait *float32 `mandatory:"true" json:"avgTimerWait"`

	// The slowest the query has been executed.
	MaxTimerWait *float32 `mandatory:"true" json:"maxTimerWait"`

	// The total amount of time that has been spent waiting for table locks.
	SumLockTime *float32 `mandatory:"true" json:"sumLockTime"`

	// The total number of errors that have been encountered executing the query.
	SumErrors *float32 `mandatory:"true" json:"sumErrors"`

	// The total number of warnings that have been encountered executing the query.
	SumWarnings *float32 `mandatory:"true" json:"sumWarnings"`

	// The total number of rows that have been modified by the query.
	SumRowsAffected *float32 `mandatory:"true" json:"sumRowsAffected"`

	// The total number of rows that have been returned (sent) to the client.
	SumRowsSent *float32 `mandatory:"true" json:"sumRowsSent"`

	// The total number of rows that have been examined by the query.
	SumRowsExamined *float32 `mandatory:"true" json:"sumRowsExamined"`

	// The total number of On-Disk internal temporary tables that have been created by the query.
	SumCreatedTempDiskTables *float32 `mandatory:"true" json:"sumCreatedTempDiskTables"`

	// The total number of internal temporary tables (in memory or on disk), which have been created by the query.
	SumCreatedTempTables *float32 `mandatory:"true" json:"sumCreatedTempTables"`

	// The total number of joins that have performed full table scans as there was no join condition or no index for the join condition. This is the same as the select_full_join status variable.
	SumSelectFullJoin *float32 `mandatory:"true" json:"sumSelectFullJoin"`

	// The total number of joins that use a full range search. This is the same as the select_full_range_join status variable.
	SumSelectFullRangeJoin *float32 `mandatory:"true" json:"sumSelectFullRangeJoin"`

	// The total number of times the query has used a range search. This is the same as the select_range status variable.
	SumSelectRange *float32 `mandatory:"true" json:"sumSelectRange"`

	// The total number of joins by the query where the join does not have an index that checks for the index usage after each row. This is the same as the select_range_check status variable.
	SumSelectRangeCheck *float32 `mandatory:"true" json:"sumSelectRangeCheck"`

	// The total number of times the query has performed a full table scan on the first table in the join. This is the same as the select_scan status variable.
	SumSelectScan *float32 `mandatory:"true" json:"sumSelectScan"`

	// The total number of sort merge passes that have been done to sort the result of the query. This is the same as the sort_merge_passes status variable.
	SumSortMergePasses *float32 `mandatory:"true" json:"sumSortMergePasses"`

	// The total number of times a sort was done using ranges. This is the same as the sort_range status variable.
	SumSortRange *float32 `mandatory:"true" json:"sumSortRange"`

	// The total number of rows sorted. This is the same as the sort_rowsStatus variable.
	SumSortRows *float32 `mandatory:"true" json:"sumSortRows"`

	// The total number of times a sort was done by scanning the table. This is the same as the sort_scan status variable.
	SumSortScan *float32 `mandatory:"true" json:"sumSortScan"`

	// The total number of times no index was used to execute the query.
	SumNoIndexUsed *float32 `mandatory:"true" json:"sumNoIndexUsed"`

	// The total number of times no good index was used. This means that the extra column in The EXPLAIN output includes “Range Checked For Each Record.”
	SumNoGoodIndexUsed *float32 `mandatory:"true" json:"sumNoGoodIndexUsed"`

	// The date and time the query was first seen. If the table is truncated, the first seen value is reset.
	FirstSeen *common.SDKTime `mandatory:"true" json:"firstSeen"`

	// The date and time the query was last seen.
	LastSeen *common.SDKTime `mandatory:"true" json:"lastSeen"`

	// The 95th percentile of the query latency. That is, 95% of the queries complete in the time given or in less time.
	Quantile95 *float32 `mandatory:"true" json:"quantile95"`

	// The 99th percentile of the query latency.
	Quantile99 *float32 `mandatory:"true" json:"quantile99"`

	// The 99.9th percentile of the query latency.
	Quantile999 *float32 `mandatory:"true" json:"quantile999"`

	// The number of query executions offloaded to HeatWave.
	HeatWaveOffloaded *float32 `mandatory:"false" json:"heatWaveOffloaded"`

	// The number of query executions with HeatWave out-of-memory errors.
	HeatWaveOutOfMemory *float32 `mandatory:"false" json:"heatWaveOutOfMemory"`
}

func (m MySqlDataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlDataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
