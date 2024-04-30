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

// HeatWaveFleetMetricDefinition An aggregated metric for a HeatWave cluster.
type HeatWaveFleetMetricDefinition struct {

	// The name of the metric.
	MetricName *string `mandatory:"true" json:"metricName"`

	// The value of the metric.
	MetricValue *float64 `mandatory:"true" json:"metricValue"`

	// The data point date and time in UTC in ISO-8601 format.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// The dimensions of the metric.
	Dimensions []MetricDimensionDefinition `mandatory:"true" json:"dimensions"`

	// The unit of the metric value.
	Unit *string `mandatory:"true" json:"unit"`
}

func (m HeatWaveFleetMetricDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveFleetMetricDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
