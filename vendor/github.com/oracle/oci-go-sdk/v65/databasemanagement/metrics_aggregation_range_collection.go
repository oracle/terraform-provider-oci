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

// MetricsAggregationRangeCollection The collection of metrics.
type MetricsAggregationRangeCollection struct {

	// The metric data.
	Items []MetricsAggregationRange `mandatory:"true" json:"items"`

	// The beginning of the metric data query time range. Expressed in UTC in
	// ISO-8601 format, which is `yyyy-MM-dd'T'hh:mm:ss.sss'Z'`.
	StartTime *string `mandatory:"false" json:"startTime"`

	// The end of the metric data query time range. Expressed in UTC in
	// ISO-8601 format, which is `yyyy-MM-dd'T'hh:mm:ss.sss'Z'`.
	EndTime *string `mandatory:"false" json:"endTime"`
}

func (m MetricsAggregationRangeCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricsAggregationRangeCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
