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

// DatabaseHomeMetricDefinition The response containing the CPU, Storage, Wait, DB Time, and Memory metrics for a specific Managed Database.
type DatabaseHomeMetricDefinition struct {

	// A list of the active session metrics for CPU and Wait time for a specific Managed Database.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseHomeMetricDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
