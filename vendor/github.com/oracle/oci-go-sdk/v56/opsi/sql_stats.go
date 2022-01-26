// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SqlStats Sql Stats type object.
type SqlStats struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Plan hash value for the SQL Execution Plan
	PlanHashValue *int64 `mandatory:"true" json:"planHashValue"`

	// Collection timestamp
	// Example: `"2020-03-31T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Name of Database Instance
	// Example: `"DB10902_1"`
	InstanceName *string `mandatory:"true" json:"instanceName"`

	// last_active_time
	// Example: `"0000000099CCE300"`
	LastActiveTime *string `mandatory:"false" json:"lastActiveTime"`

	// Total integer of parse calls
	//  Example: `60`
	ParseCalls *int64 `mandatory:"false" json:"parseCalls"`

	// Number of disk reads
	DiskReads *int64 `mandatory:"false" json:"diskReads"`

	// Number of direct reads
	DirectReads *int64 `mandatory:"false" json:"directReads"`

	// Number of Direct writes
	DirectWrites *int64 `mandatory:"false" json:"directWrites"`

	// Number of Buffer Gets
	BufferGets *int64 `mandatory:"false" json:"bufferGets"`

	// Number of row processed
	RowsProcessed *int64 `mandatory:"false" json:"rowsProcessed"`

	// Number of serializable aborts
	SerializableAborts *int64 `mandatory:"false" json:"serializableAborts"`

	// Number of fetches
	Fetches *int64 `mandatory:"false" json:"fetches"`

	// Number of executions
	Executions *int64 `mandatory:"false" json:"executions"`

	// Number of executions attempted on this object, but prevented due to the SQL statement being in quarantine
	AvoidedExecutions *int64 `mandatory:"false" json:"avoidedExecutions"`

	// Number of times this cursor was fully executed since the cursor was brought into the library cache
	EndOfFetchCount *int64 `mandatory:"false" json:"endOfFetchCount"`

	// Number of times the object was either loaded or reloaded
	Loads *int64 `mandatory:"false" json:"loads"`

	// Number of cursors present in the cache with this SQL text and plan
	VersionCount *int64 `mandatory:"false" json:"versionCount"`

	// Number of times this child cursor has been invalidated
	Invalidations *int64 `mandatory:"false" json:"invalidations"`

	// Number of times that a parent cursor became obsolete
	ObsoleteCount *int64 `mandatory:"false" json:"obsoleteCount"`

	// Total number of executions performed by parallel execution servers (0 when the statement has never been executed in parallel)
	PxServersExecutions *int64 `mandatory:"false" json:"pxServersExecutions"`

	// CPU time (in microseconds) used by this cursor for parsing, executing, and fetching
	CpuTimeInUs *int64 `mandatory:"false" json:"cpuTimeInUs"`

	// Elapsed time (in microseconds) used by this cursor for parsing, executing, and fetching.
	ElapsedTimeInUs *int64 `mandatory:"false" json:"elapsedTimeInUs"`

	// Average hard parse time (in microseconds) used by this cursor
	AvgHardParseTimeInUs *int64 `mandatory:"false" json:"avgHardParseTimeInUs"`

	// Concurrency wait time (in microseconds)
	ConcurrencyWaitTimeInUs *int64 `mandatory:"false" json:"concurrencyWaitTimeInUs"`

	// Application wait time (in microseconds)
	ApplicationWaitTimeInUs *int64 `mandatory:"false" json:"applicationWaitTimeInUs"`

	// Cluster wait time (in microseconds). This value is specific to Oracle RAC
	ClusterWaitTimeInUs *int64 `mandatory:"false" json:"clusterWaitTimeInUs"`

	// User I/O wait time (in microseconds)
	UserIoWaitTimeInUs *int64 `mandatory:"false" json:"userIoWaitTimeInUs"`

	// PL/SQL execution time (in microseconds)
	PlsqlExecTimeInUs *int64 `mandatory:"false" json:"plsqlExecTimeInUs"`

	// Java execution time (in microseconds)
	JavaExecTimeInUs *int64 `mandatory:"false" json:"javaExecTimeInUs"`

	// Number of sorts that were done for the child cursor
	Sorts *int64 `mandatory:"false" json:"sorts"`

	// Total shared memory (in bytes) currently occupied by all cursors with this SQL text and plan
	SharableMem *int64 `mandatory:"false" json:"sharableMem"`

	// Total shared memory (in bytes) occupied by all cursors with this SQL text and plan if they were to be fully loaded in the shared pool (that is, cursor size)
	TotalSharableMem *int64 `mandatory:"false" json:"totalSharableMem"`

	// Typecheck memory
	TypeCheckMem *int64 `mandatory:"false" json:"typeCheckMem"`

	// Number of I/O bytes which can be filtered by the Exadata storage system
	IoCellOffloadEligibleBytes *int64 `mandatory:"false" json:"ioCellOffloadEligibleBytes"`

	// Number of I/O bytes exchanged between Oracle Database and the storage system. Typically used for Cache Fusion or parallel queries
	IoInterconnectBytes *int64 `mandatory:"false" json:"ioInterconnectBytes"`

	// Number of physical read I/O requests issued by the monitored SQL. The requests may not be disk reads
	PhysicalReadRequests *int64 `mandatory:"false" json:"physicalReadRequests"`

	// Number of bytes read from disks by the monitored SQL
	PhysicalReadBytes *int64 `mandatory:"false" json:"physicalReadBytes"`

	// Number of physical write I/O requests issued by the monitored SQL
	PhysicalWriteRequests *int64 `mandatory:"false" json:"physicalWriteRequests"`

	// Number of bytes written to disks by the monitored SQL
	PhysicalWriteBytes *int64 `mandatory:"false" json:"physicalWriteBytes"`

	// exact_matching_signature
	// Example: `"18067345456756876713"`
	ExactMatchingSignature *string `mandatory:"false" json:"exactMatchingSignature"`

	// force_matching_signature
	// Example: `"18067345456756876713"`
	ForceMatchingSignature *string `mandatory:"false" json:"forceMatchingSignature"`

	// Number of uncompressed bytes (that is, size after decompression) that are offloaded to the Exadata cells
	IoCellUncompressedBytes *int64 `mandatory:"false" json:"ioCellUncompressedBytes"`

	// Number of bytes that are returned by Exadata cell through the regular I/O path
	IoCellOffloadReturnedBytes *int64 `mandatory:"false" json:"ioCellOffloadReturnedBytes"`

	// Number of this child cursor
	ChildNumber *int64 `mandatory:"false" json:"childNumber"`

	// Oracle command type definition
	CommandType *int64 `mandatory:"false" json:"commandType"`

	// Number of users that have any of the child cursors open
	UsersOpening *int64 `mandatory:"false" json:"usersOpening"`

	// Number of users executing the statement
	UsersExecuting *int64 `mandatory:"false" json:"usersExecuting"`

	// Cost of this query given by the optimizer
	OptimizerCost *int64 `mandatory:"false" json:"optimizerCost"`

	// Total Number of rows in SQLStats table
	FullPlanHashValue *string `mandatory:"false" json:"fullPlanHashValue"`

	// Module name
	Module *string `mandatory:"false" json:"module"`

	// Service name
	Service *string `mandatory:"false" json:"service"`

	// Contains the name of the action that was executing when the SQL statement was first parsed, which is set by calling DBMS_APPLICATION_INFO.SET_ACTION
	Action *string `mandatory:"false" json:"action"`

	// SQL profile used for this statement, if any
	SqlProfile *string `mandatory:"false" json:"sqlProfile"`

	// SQL patch used for this statement, if any
	SqlPatch *string `mandatory:"false" json:"sqlPatch"`

	// SQL plan baseline used for this statement, if any
	SqlPlanBaseline *string `mandatory:"false" json:"sqlPlanBaseline"`

	// Number of executions for the cursor since the last AWR snapshot
	DeltaExecutionCount *int64 `mandatory:"false" json:"deltaExecutionCount"`

	// CPU time (in microseconds) for the cursor since the last AWR snapshot
	DeltaCpuTime *int64 `mandatory:"false" json:"deltaCpuTime"`

	// Number of I/O bytes exchanged between the Oracle database and the storage system for the cursor since the last AWR snapshot
	DeltaIoBytes *int64 `mandatory:"false" json:"deltaIoBytes"`

	// Rank based on CPU Consumption
	DeltaCpuRank *int64 `mandatory:"false" json:"deltaCpuRank"`

	// Rank based on number of execution
	DeltaExecsRank *int64 `mandatory:"false" json:"deltaExecsRank"`

	// Rank based on sharable memory
	SharableMemRank *int64 `mandatory:"false" json:"sharableMemRank"`

	// Rank based on I/O Consumption
	DeltaIoRank *int64 `mandatory:"false" json:"deltaIoRank"`

	// Harmonic sum based on ranking parameters
	HarmonicSum *int64 `mandatory:"false" json:"harmonicSum"`

	// Weight based harmonic sum of ranking parameters
	WtHarmonicSum *int64 `mandatory:"false" json:"wtHarmonicSum"`

	// Total number of rows in SQLStats table
	TotalSqlCount *int64 `mandatory:"false" json:"totalSqlCount"`
}

func (m SqlStats) String() string {
	return common.PointerString(m)
}
