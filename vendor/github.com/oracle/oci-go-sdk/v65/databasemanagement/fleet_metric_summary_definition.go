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

// FleetMetricSummaryDefinition A summary of the fleet metrics, which provides the metric aggregated value of the databases in the fleet.
type FleetMetricSummaryDefinition struct {

	// The name of the metric.
	MetricName *string `mandatory:"false" json:"metricName"`

	// The metric aggregated value at the baseline date and time.
	BaselineValue *float64 `mandatory:"false" json:"baselineValue"`

	// The metric aggregated value at the target date and time.
	TargetValue *float64 `mandatory:"false" json:"targetValue"`

	// The unit of the value.
	Unit *string `mandatory:"false" json:"unit"`

	// The percentage change in the metric aggregated value compared to the baseline value.
	PercentageChange *float64 `mandatory:"false" json:"percentageChange"`

	// The unique dimension key and values of the baseline metric.
	Dimensions []MetricDimensionDefinition `mandatory:"false" json:"dimensions"`
}

func (m FleetMetricSummaryDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetMetricSummaryDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
