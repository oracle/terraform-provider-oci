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

// DatabaseFleetHealthMetrics The details of the fleet health metrics.
type DatabaseFleetHealthMetrics struct {

	// The baseline date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	// This is the date and time against which percentage change is calculated.
	CompareBaselineTime *string `mandatory:"true" json:"compareBaselineTime"`

	// The target date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	// All the metrics are returned for the target date and time and the percentage change
	// is calculated against the baseline date and time.
	CompareTargetTime *string `mandatory:"true" json:"compareTargetTime"`

	// A list of the databases present in the fleet and their usage metrics.
	FleetDatabases []DatabaseUsageMetrics `mandatory:"true" json:"fleetDatabases"`

	// The time window used for metrics comparison.
	CompareType CompareTypeEnum `mandatory:"false" json:"compareType,omitempty"`

	FleetSummary *FleetSummary `mandatory:"false" json:"fleetSummary"`
}

func (m DatabaseFleetHealthMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseFleetHealthMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCompareTypeEnum(string(m.CompareType)); !ok && m.CompareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompareType: %s. Supported values are: %s.", m.CompareType, strings.Join(GetCompareTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
