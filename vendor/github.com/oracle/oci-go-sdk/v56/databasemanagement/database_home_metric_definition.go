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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseHomeMetricDefinition The response containing the CPU, Storage, Wait, DB Time, and Memory metrics for a specific database.
type DatabaseHomeMetricDefinition struct {

	// A list of the active session metrics for CPU and Wait time for a specific database.
	ActivityTimeSeriesMetrics []ActivityTimeSeriesMetrics `mandatory:"true" json:"activityTimeSeriesMetrics"`

	DbTimeAggregateMetrics *DatabaseTimeAggregateMetrics `mandatory:"true" json:"dbTimeAggregateMetrics"`

	IoAggregateMetrics *DatabaseIoAggregateMetrics `mandatory:"true" json:"ioAggregateMetrics"`

	MemoryAggregateMetrics *MemoryAggregateMetrics `mandatory:"true" json:"memoryAggregateMetrics"`

	DbStorageAggregateMetrics *DatabaseStorageAggregateMetrics `mandatory:"true" json:"dbStorageAggregateMetrics"`

	CpuUtilizationAggregateMetrics *CpuUtilizationAggregateMetrics `mandatory:"false" json:"cpuUtilizationAggregateMetrics"`

	StatementsAggregateMetrics *StatementsAggregateMetrics `mandatory:"false" json:"statementsAggregateMetrics"`

	FailedConnectionsAggregateMetrics *FailedConnectionsAggregateMetrics `mandatory:"false" json:"failedConnectionsAggregateMetrics"`
}

func (m DatabaseHomeMetricDefinition) String() string {
	return common.PointerString(m)
}
