// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseInstanceHomeMetricsDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
