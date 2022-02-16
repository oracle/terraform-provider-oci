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

// SqlTuningTaskPlanStats The statistics of a SQL execution plan.
type SqlTuningTaskPlanStats struct {

	// The type of the original or modified plan with profile, index, and so on.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningTaskPlanStats) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningTaskPlanStatsPlanStatusEnum(string(m.PlanStatus)); !ok && m.PlanStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanStatus: %s. Supported values are: %s.", m.PlanStatus, strings.Join(GetSqlTuningTaskPlanStatsPlanStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningTaskPlanStatsPlanStatusEnum Enum with underlying type: string
type SqlTuningTaskPlanStatsPlanStatusEnum string

// Set of constants representing the allowable values for SqlTuningTaskPlanStatsPlanStatusEnum
const (
	SqlTuningTaskPlanStatsPlanStatusComplete SqlTuningTaskPlanStatsPlanStatusEnum = "COMPLETE"
	SqlTuningTaskPlanStatsPlanStatusPartial  SqlTuningTaskPlanStatsPlanStatusEnum = "PARTIAL"
)

var mappingSqlTuningTaskPlanStatsPlanStatusEnum = map[string]SqlTuningTaskPlanStatsPlanStatusEnum{
	"COMPLETE": SqlTuningTaskPlanStatsPlanStatusComplete,
	"PARTIAL":  SqlTuningTaskPlanStatsPlanStatusPartial,
}

// GetSqlTuningTaskPlanStatsPlanStatusEnumValues Enumerates the set of values for SqlTuningTaskPlanStatsPlanStatusEnum
func GetSqlTuningTaskPlanStatsPlanStatusEnumValues() []SqlTuningTaskPlanStatsPlanStatusEnum {
	values := make([]SqlTuningTaskPlanStatsPlanStatusEnum, 0)
	for _, v := range mappingSqlTuningTaskPlanStatsPlanStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningTaskPlanStatsPlanStatusEnumStringValues Enumerates the set of values in String for SqlTuningTaskPlanStatsPlanStatusEnum
func GetSqlTuningTaskPlanStatsPlanStatusEnumStringValues() []string {
	return []string{
		"COMPLETE",
		"PARTIAL",
	}
}

// GetMappingSqlTuningTaskPlanStatsPlanStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningTaskPlanStatsPlanStatusEnum(val string) (SqlTuningTaskPlanStatsPlanStatusEnum, bool) {
	mappingSqlTuningTaskPlanStatsPlanStatusEnumIgnoreCase := make(map[string]SqlTuningTaskPlanStatsPlanStatusEnum)
	for k, v := range mappingSqlTuningTaskPlanStatsPlanStatusEnum {
		mappingSqlTuningTaskPlanStatsPlanStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSqlTuningTaskPlanStatsPlanStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
