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

// SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummary A summary for all the statistic findings of an object in a SQL Tuning Advisor task. Includes the object's hash, name, type, schema, problem type and the object reference count.
type SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummary struct {

	// Numerical representation of the object.
	ObjectHashValue *int64 `mandatory:"true" json:"objectHashValue"`

	// Name of the object.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Type of the object.
	ObjectType *string `mandatory:"true" json:"objectType"`

	// Schema of the object.
	Schema *string `mandatory:"true" json:"schema"`

	// Type of statistics problem related to the object.
	ProblemType SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum `mandatory:"true" json:"problemType"`

	// The number of the times the object is referenced within the SQL Tuning advisor task findings.
	ReferenceCount *int `mandatory:"true" json:"referenceCount"`
}

func (m SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum(string(m.ProblemType)); !ok && m.ProblemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemType: %s. Supported values are: %s.", m.ProblemType, strings.Join(GetSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum Enum with underlying type: string
type SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum string

// Set of constants representing the allowable values for SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum
const (
	SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeMissing SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum = "MISSING"
	SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeStale   SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum = "STALE"
)

var mappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum = map[string]SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum{
	"MISSING": SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeMissing,
	"STALE":   SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeStale,
}

var mappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumLowerCase = map[string]SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum{
	"missing": SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeMissing,
	"stale":   SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeStale,
}

// GetSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumValues Enumerates the set of values for SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum
func GetSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumValues() []SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum {
	values := make([]SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum, 0)
	for _, v := range mappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumStringValues Enumerates the set of values in String for SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum
func GetSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumStringValues() []string {
	return []string{
		"MISSING",
		"STALE",
	}
}

// GetMappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum(val string) (SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnum, bool) {
	enum, ok := mappingSqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryProblemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
