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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsCollectionAggregationSummary This provides optimizer statistics collection summary which includes aggregated counts that are grouped by task status.
type OptimizerStatisticsCollectionAggregationSummary struct {

	// As the statistics are aggregated per hour this date time indicates the start of the hour.
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The optimizer statistics tasks group by type.
	GroupBy OptimizerStatisticsGroupByTypesEnum `mandatory:"false" json:"groupBy,omitempty"`

	// As the statistics are aggregated per hour this date time indicates the end of the hour.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Count of tasks/objects for which statistics have yet to be gathered.
	Pending *int `mandatory:"false" json:"pending"`

	// Count of tasks/objects for which statistics gathering is in progress.
	InProgress *int `mandatory:"false" json:"inProgress"`

	// Count of tasks/objects for which statistics gathering completed.
	Completed *int `mandatory:"false" json:"completed"`

	// Count of tasks/objects for which statistics gathering failed.
	Failed *int `mandatory:"false" json:"failed"`

	// Count of tasks/objects for which statistics gathering skipped.
	Skipped *int `mandatory:"false" json:"skipped"`

	// Count of tasks/objects for which statistics gathering timed out.
	TimedOut *int `mandatory:"false" json:"timedOut"`

	// Count of tasks/objects for which statistics gathering is unknown.
	Unknown *int `mandatory:"false" json:"unknown"`

	// Total count of tasks/objects for which statistics collection finished. This count is sum of pending, inProgress, completed,
	// failed, skipped, timedOut and unknown status.
	Total *int `mandatory:"false" json:"total"`
}

func (m OptimizerStatisticsCollectionAggregationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsCollectionAggregationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOptimizerStatisticsGroupByTypesEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetOptimizerStatisticsGroupByTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
