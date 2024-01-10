// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MetricStatisticsDefinition The metric statistics values with dimension details.
type MetricStatisticsDefinition struct {

	// The minimum value of the metric.
	Min *float64 `mandatory:"false" json:"min"`

	// The maximum value of the metric.
	Max *float64 `mandatory:"false" json:"max"`

	// The median value of the metric.
	Median *float64 `mandatory:"false" json:"median"`

	// The first quartile value of the metric.
	LowerQuartile *float64 `mandatory:"false" json:"lowerQuartile"`

	// The third quartile value of the metric.
	UpperQuartile *float64 `mandatory:"false" json:"upperQuartile"`

	// The unit of the metric value.
	Unit *string `mandatory:"false" json:"unit"`

	// The dimensions of the metric.
	Dimensions []MetricDimensionDefinition `mandatory:"false" json:"dimensions"`
}

func (m MetricStatisticsDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricStatisticsDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
