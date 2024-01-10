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

// SqlTuningAdvisorTaskRecommendationSummary A recommendation for a given object in a SQL Tuning Task.
type SqlTuningAdvisorTaskRecommendationSummary struct {

	// The unique identifier of the task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskId"`

	// The key of the object to which these recommendations apply. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskObjectId *int64 `mandatory:"true" json:"sqlTuningAdvisorTaskObjectId"`

	// The unique identifier of the recommendation in the scope of the task.
	RecommendationKey *int `mandatory:"true" json:"recommendationKey"`

	// Type of recommendation.
	RecommendationType SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum `mandatory:"true" json:"recommendationType"`

	// Summary of the issue found in the SQL statement.
	Finding *string `mandatory:"false" json:"finding"`

	// The recommendation for a specific finding.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// Describes the reasoning behind the recommendation and how it relates to the finding.
	Rationale *string `mandatory:"false" json:"rationale"`

	// The percentage benefit of this implementation.
	Benefit *float32 `mandatory:"false" json:"benefit"`

	// Action sql to be implemented based on the recommendation result.
	ImplementActionSql *string `mandatory:"false" json:"implementActionSql"`

	// Indicates whether a SQL Profile recommendation uses parallel execution.
	IsParallelExecution *bool `mandatory:"false" json:"isParallelExecution"`
}

func (m SqlTuningAdvisorTaskRecommendationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskRecommendationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum(string(m.RecommendationType)); !ok && m.RecommendationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecommendationType: %s. Supported values are: %s.", m.RecommendationType, strings.Join(GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum = map[string]SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum{
	"STATISTICS":        SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeStatistics,
	"INDEX":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeIndex,
	"SQL_PROFILE":       SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeSqlProfile,
	"RESTRUCTURE_SQL":   SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeRestructureSql,
	"ALTERNATIVE_PLANS": SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeAlternativePlans,
	"ERROR":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeError,
	"MISCELLANEOUS":     SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeMiscellaneous,
}

var mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumLowerCase = map[string]SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum{
	"statistics":        SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeStatistics,
	"index":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeIndex,
	"sql_profile":       SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeSqlProfile,
	"restructure_sql":   SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeRestructureSql,
	"alternative_plans": SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeAlternativePlans,
	"error":             SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeError,
	"miscellaneous":     SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeMiscellaneous,
}

// GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumValues Enumerates the set of values for SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum
func GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumValues() []SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum {
	values := make([]SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum, 0)
	for _, v := range mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumStringValues Enumerates the set of values in String for SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum
func GetSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumStringValues() []string {
	return []string{
		"STATISTICS",
		"INDEX",
		"SQL_PROFILE",
		"RESTRUCTURE_SQL",
		"ALTERNATIVE_PLANS",
		"ERROR",
		"MISCELLANEOUS",
	}
}

// GetMappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum(val string) (SqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnum, bool) {
	enum, ok := mappingSqlTuningAdvisorTaskRecommendationSummaryRecommendationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
