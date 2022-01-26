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

// MetricDataPoint The metric values with dimension details.
type MetricDataPoint struct {

	// The value of the metric.
	Value *float64 `mandatory:"false" json:"value"`

	// The unit of the metric value.
	Unit *string `mandatory:"false" json:"unit"`

	// The dimensions of the metric.
	Dimensions []MetricDimensionDefinition `mandatory:"false" json:"dimensions"`
}

func (m MetricDataPoint) String() string {
	return common.PointerString(m)
}
