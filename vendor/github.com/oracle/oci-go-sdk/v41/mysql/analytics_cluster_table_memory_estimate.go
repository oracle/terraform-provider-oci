// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AnalyticsClusterTableMemoryEstimate DEPRECATED -- please use HeatWave API instead.
// Estimated memory footprint for a MySQL user table
// when loaded to the Analytics Cluster memory.
type AnalyticsClusterTableMemoryEstimate struct {

	// The table name.
	TableName *string `mandatory:"true" json:"tableName"`

	// The number of columns to be loaded to Analytics Cluster memory.
	// These columns contribute to the analytical memory footprint.
	ToLoadColumnCount *int `mandatory:"true" json:"toLoadColumnCount"`

	// The number of variable-length columns to be loaded to Analytics Cluster memory.
	// These columns contribute to the analytical memory footprint.
	VarlenColumnCount *int `mandatory:"true" json:"varlenColumnCount"`

	// The estimated number of rows in the table. This number was used to
	// derive the analytical memory footprint.
	EstimatedRowCount *int64 `mandatory:"true" json:"estimatedRowCount"`

	// The estimated memory footprint of the table in MBs when loaded to
	// Analytics Cluster memory (null if the table cannot be loaded to the
	// Analytics Cluster).
	AnalyticalFootprintInMbs *int64 `mandatory:"true" json:"analyticalFootprintInMbs"`

	// Error comment (empty string if no errors occured).
	ErrorComment *string `mandatory:"true" json:"errorComment"`
}

func (m AnalyticsClusterTableMemoryEstimate) String() string {
	return common.PointerString(m)
}
