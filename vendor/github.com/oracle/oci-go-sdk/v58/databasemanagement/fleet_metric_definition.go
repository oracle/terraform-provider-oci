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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FleetMetricDefinition The database metric details.
type FleetMetricDefinition struct {

	// The name of the metric.
	MetricName *string `mandatory:"false" json:"metricName"`

	// The baseline value of the metric.
	BaselineValue *float64 `mandatory:"false" json:"baselineValue"`

	// The target value of the metric.
	TargetValue *float64 `mandatory:"false" json:"targetValue"`

	// The unit of the value.
	Unit *string `mandatory:"false" json:"unit"`

	// The data point date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The percentage change in the metric aggregated value compared to the baseline value.
	PercentageChange *float64 `mandatory:"false" json:"percentageChange"`

	// The dimensions of the metric.
	Dimensions []MetricDimensionDefinition `mandatory:"false" json:"dimensions"`
}

func (m FleetMetricDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetMetricDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
