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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlDataSummary SQL Performance Data Record for a given SQL.
type MySqlDataSummary struct {

	// The Schema That Was The Default Schema When Executing The Query. If No Schema Was The Default, The Value Is NULL.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The Digest Of The Normalized Query.
	Digest *string `mandatory:"true" json:"digest"`

	// The Normalized Query.
	DigestText *string `mandatory:"true" json:"digestText"`

	// The Number Of Times The Query Has Been Executed.
	CountStar *float32 `mandatory:"true" json:"countStar"`

	// The Total Amount Of Time That Has Been Spent Executing The Query.
	SumTimerWait *float32 `mandatory:"true" json:"sumTimerWait"`

	// The Fastest The Query Has Been Executed.
	MinTimerWait *float32 `mandatory:"true" json:"minTimerWait"`

	// The Average Execution Time.
	AvgTimerWait *float32 `mandatory:"true" json:"avgTimerWait"`

	// The Slowest The Query Has Been Executed.
	MaxTimerWait *float32 `mandatory:"true" json:"maxTimerWait"`

	// The Total Amount Of Time That Has Been Spent Waiting For Table Locks.
	SumLockTime *float32 `mandatory:"true" json:"sumLockTime"`

	// The Total Number Of Errors That Have Been Encountered Executing The Query.
	SumErrors *float32 `mandatory:"true" json:"sumErrors"`

	// The Total Number Of Warnings That Have Been Encountered Executing The Query.
	SumWarnings *float32 `mandatory:"true" json:"sumWarnings"`

	// The Total Number Of Rows That Have Been Modified By The Query.
	SumRowsAffected *float32 `mandatory:"true" json:"sumRowsAffected"`

	// The Total Number Of Rows That Have Been Returned (Sent) To The Client.
	SumRowsSent *float32 `mandatory:"true" json:"sumRowsSent"`

	// The Total Number Of Rows That Have Been Examined By The Query.
	SumRowsExamined *float32 `mandatory:"true" json:"sumRowsExamined"`

	// The Total Number Of On-Disk Internal Temporary Tables That Have Been Created By The Query.
	SumCreatedTempDiskTables *float32 `mandatory:"true" json:"sumCreatedTempDiskTables"`

	// The Total Number Of Internal Temporary Tables – Whether Created In Memory Or On Disk – That Have Been Created By The Query.
	SumCreatedTempTables *float32 `mandatory:"true" json:"sumCreatedTempTables"`

	// The Total Number Of Joins That Have Performed Full Table Scans As There Is No Index For The Join Condition Or There Is No Join Condition. This Is The Same That Increments The Select_full_join Status Variable.
	SumSelectFullJoin *float32 `mandatory:"true" json:"sumSelectFullJoin"`

	// The Total Number Of Joins That Use A Full Range Search. This Is The Same That Increments The Select_full_range_join Status Variable.
	SumSelectFullRangeJoin *float32 `mandatory:"true" json:"sumSelectFullRangeJoin"`

	// The Total Number Of Times The Query Has Used A Range Search. This Is The Same That Increments The Select_range Status Variable.
	SumSelectRange *float32 `mandatory:"true" json:"sumSelectRange"`

	// The Total Number Of Joins By The Query Where The Join Does Not Have An Index That Checks For The Index Usage After Each Row. This Is The Same That Increments The Select_range_check Status Variable.
	SumSelectRangeCheck *float32 `mandatory:"true" json:"sumSelectRangeCheck"`

	// The Total Number Of Times The Query Has Performed A Full Table Scan On The First Table In The Join. This Is The Same That Increments The Select_scan Status Variable.
	SumSelectScan *float32 `mandatory:"true" json:"sumSelectScan"`

	// The Total Number Of Sort Merge Passes That Have Been Done To Sort The Result Of The Query. This Is The Same That Increments The Sort_merge_passes Status Variable.
	SumSortMergePasses *float32 `mandatory:"true" json:"sumSortMergePasses"`

	// The Total Number Of Times A Sort Was Done Using Ranges. This Is The Same That Increments The Sort_range Status Variable.
	SumSortRange *float32 `mandatory:"true" json:"sumSortRange"`

	// The Total Number Of Rows Sorted. This Is The Same That Increments The Sort_rowsStatus Variable.
	SumSortRows *float32 `mandatory:"true" json:"sumSortRows"`

	// The Total Number Of Times A Sort Was Done By Scanning The Table. This Is The Same That Increments The Sort_scan Status Variable.
	SumSortScan *float32 `mandatory:"true" json:"sumSortScan"`

	// The Total Number Of Times No Index Was Used To Execute The Query.
	SumNoIndexUsed *float32 `mandatory:"true" json:"sumNoIndexUsed"`

	// The Total Number Of Times No Good Index Was Used. This Means That The ExtraColumn In The EXPLAIN Output Includes “Range Checked For Each Record.”
	SumNoGoodIndexUsed *float32 `mandatory:"true" json:"sumNoGoodIndexUsed"`

	// When The Query Was First Seen. When The Table Is Truncated, The First Seen Value Is Also Reset.
	FirstSeen *common.SDKTime `mandatory:"true" json:"firstSeen"`

	// When The Query Was Seen The Last Time.
	LastSeen *common.SDKTime `mandatory:"true" json:"lastSeen"`

	// The 95th Percentile Of The Query Latency. That Is, 95% Of The Queries Complete In The Time Given Or In Less Time.
	Quantile95 *float32 `mandatory:"true" json:"quantile95"`

	// The 99th Percentile Of The Query Latency.
	Quantile99 *float32 `mandatory:"true" json:"quantile99"`

	// The 99.9th Percentile Of The Query Latency.
	Quantile999 *float32 `mandatory:"true" json:"quantile999"`
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
