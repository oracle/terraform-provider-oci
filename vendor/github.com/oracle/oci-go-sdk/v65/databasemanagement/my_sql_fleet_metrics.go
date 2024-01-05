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

// MySqlFleetMetrics The details of the MySQL Database fleet health metrics.
type MySqlFleetMetrics struct {

	// The beginning of the time range during which metric data is retrieved.
	StartTime *string `mandatory:"true" json:"startTime"`

	// The end of the time range during which metric data is retrieved.
	EndTime *string `mandatory:"true" json:"endTime"`

	// The list of MySQL Databases in the fleet and their usage metrics.
	FleetDatabases []MySqlDatabaseUsageMetrics `mandatory:"true" json:"fleetDatabases"`

	// A summary of the inventory count and the metrics that describe the aggregated usage of CPU, storage, and so on of all the MySQL Databases in the fleet.
	FleetSummary []MySqlFleetSummary `mandatory:"true" json:"fleetSummary"`
}

func (m MySqlFleetMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlFleetMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
