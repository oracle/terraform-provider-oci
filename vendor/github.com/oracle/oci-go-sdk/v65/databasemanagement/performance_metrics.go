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

// PerformanceMetrics The Data Guard performance metric details.
type PerformanceMetrics struct {

	// The name of the metric.
	Name *string `mandatory:"false" json:"name"`

	// The dimensions of the Data Guard performance metrics, such as primary database ID, primary database unique name.
	Dimensions *interface{} `mandatory:"false" json:"dimensions"`

	// The metadata of the metric, such as Unit.
	Metadata *interface{} `mandatory:"false" json:"metadata"`

	// The aggregated datapoints of the metric.
	Datapoints []DataPoints `mandatory:"false" json:"datapoints"`
}

func (m PerformanceMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PerformanceMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
