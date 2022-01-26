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

// SqlTuningTaskPlanStats The statistics of an SQL execution plan.
type SqlTuningTaskPlanStats struct {

	// The type of the plan for the original or the new plan with profile/index etc.
	PlanType *string `mandatory:"true" json:"planType"`

	// A map contains the statistics for the SQL execution using the plan.
	// The key of the map is the metric's name. The value of the map is the metric's value.
	PlanStats map[string]float64 `mandatory:"true" json:"planStats"`

	// The status of the execution using the plan.
	PlanStatus SqlTuningTaskPlanStatsPlanStatusEnum `mandatory:"true" json:"planStatus"`
}

func (m SqlTuningTaskPlanStats) String() string {
	return common.PointerString(m)
}

// SqlTuningTaskPlanStatsPlanStatusEnum Enum with underlying type: string
type SqlTuningTaskPlanStatsPlanStatusEnum string

// Set of constants representing the allowable values for SqlTuningTaskPlanStatsPlanStatusEnum
const (
	SqlTuningTaskPlanStatsPlanStatusComplete SqlTuningTaskPlanStatsPlanStatusEnum = "COMPLETE"
	SqlTuningTaskPlanStatsPlanStatusPartial  SqlTuningTaskPlanStatsPlanStatusEnum = "PARTIAL"
)

var mappingSqlTuningTaskPlanStatsPlanStatus = map[string]SqlTuningTaskPlanStatsPlanStatusEnum{
	"COMPLETE": SqlTuningTaskPlanStatsPlanStatusComplete,
	"PARTIAL":  SqlTuningTaskPlanStatsPlanStatusPartial,
}

// GetSqlTuningTaskPlanStatsPlanStatusEnumValues Enumerates the set of values for SqlTuningTaskPlanStatsPlanStatusEnum
func GetSqlTuningTaskPlanStatsPlanStatusEnumValues() []SqlTuningTaskPlanStatsPlanStatusEnum {
	values := make([]SqlTuningTaskPlanStatsPlanStatusEnum, 0)
	for _, v := range mappingSqlTuningTaskPlanStatsPlanStatus {
		values = append(values, v)
	}
	return values
}
