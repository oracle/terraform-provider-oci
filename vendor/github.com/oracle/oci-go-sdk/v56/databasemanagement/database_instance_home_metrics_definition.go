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

// DatabaseInstanceHomeMetricsDefinition The response containing the CPU, Wait, DB Time, and Memory metrics
// for a specific Oracle Real Application Clusters (Oracle RAC) database
// instance.
type DatabaseInstanceHomeMetricsDefinition struct {

	// The name of the Oracle Real Application Clusters (Oracle RAC)
	// database instance to which the corresponding metrics belong.
	InstanceName *string `mandatory:"true" json:"instanceName"`

	// The number of Oracle Real Application Clusters (Oracle RAC)
	// database instance to which the corresponding metrics belong.
	InstanceNumber *int `mandatory:"true" json:"instanceNumber"`

	// A list of the active session metrics for CPU and Wait time for
	// a specific Oracle Real Application Clusters (Oracle RAC)
	// database instance.
	ActivityTimeSeriesMetrics []ActivityTimeSeriesMetrics `mandatory:"true" json:"activityTimeSeriesMetrics"`

	DbTimeAggregateMetrics *DatabaseTimeAggregateMetrics `mandatory:"true" json:"dbTimeAggregateMetrics"`

	IoAggregateMetrics *DatabaseIoAggregateMetrics `mandatory:"true" json:"ioAggregateMetrics"`

	MemoryAggregateMetrics *MemoryAggregateMetrics `mandatory:"true" json:"memoryAggregateMetrics"`

	CpuUtilizationAggregateMetrics *CpuUtilizationAggregateMetrics `mandatory:"false" json:"cpuUtilizationAggregateMetrics"`
}

func (m DatabaseInstanceHomeMetricsDefinition) String() string {
	return common.PointerString(m)
}
