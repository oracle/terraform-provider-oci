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

// SqlTuningAdvisorTaskRecommendationSummary A recommendation for a given object in a SQL Tuning Task.
type SqlTuningAdvisorTaskRecommendationSummary struct {

	// Unique identifier of the task. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// Key of the object to which these recommendations apply. It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectId"`

	// Unique identifier of the recommendation in the scope of the task.
	RecommendationKey *int `mandatory:"true" json:"recommendationKey"`

	// Type of recommendation
	RecommendationType SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum `mandatory:"true" json:"recommendationType"`

	// Summary of the issue found for the SQL statement.
	Finding *string `mandatory:"false" json:"finding"`

	// Particular recommendation for the finding.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// Describes the reasoning behind the recommendation and how it relates to the finding.
	Rationale *string `mandatory:"false" json:"rationale"`

	// The percentage benefit of this implementation.
	Benefit *float32 `mandatory:"false" json:"benefit"`

	// Action sql to be implemented based on the recommendation result.
	ImplementActionSql *string `mandatory:"false" json:"implementActionSql"`
}

func (m SqlTuningAdvisorTaskRecommendationSummary) String() string {
	return common.PointerString(m)
}

// SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum Enum with underlying type: string
type SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum string

// Set of constants representing the allowable values for SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum
const (
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeStatistics       SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "STATISTICS"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeIndex            SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "INDEX"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeSqlProfile       SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "SQL_PROFILE"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeRestructureSql   SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "RESTRUCTURE_SQL"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeAlternativePlans SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "ALTERNATIVE_PLANS"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeError            SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "ERROR"
	SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeMiscellaneous    SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = "MISCELLANEOUS"
)

var mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationType = map[string]SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum{
	"STATISTICS":        SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeStatistics,
	"INDEX":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeIndex,
	"SQL_PROFILE":       SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeSqlProfile,
	"RESTRUCTURE_SQL":   SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeRestructureSql,
	"ALTERNATIVE_PLANS": SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeAlternativePlans,
	"ERROR":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeError,
	"MISCELLANEOUS":     SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeMiscellaneous,
}

// GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumValues Enumerates the set of values for SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum
func GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumValues() []SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum {
	values := make([]SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum, 0)
	for _, v := range mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationType {
		values = append(values, v)
	}
	return values
}
