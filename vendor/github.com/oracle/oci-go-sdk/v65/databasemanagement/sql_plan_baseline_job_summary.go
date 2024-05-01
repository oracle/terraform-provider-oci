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

// SqlPlanBaselineJobSummary A summary of the database job used for loading and evolving SQL plan baselines.
type SqlPlanBaselineJobSummary struct {

	// The name of the job.
	Name *string `mandatory:"true" json:"name"`

	// The type of the job.
	Type SqlPlanBaselineJobSummaryTypeEnum `mandatory:"true" json:"type"`

	// The status of the job.
	Status SqlPlanBaselineJobSummaryStatusEnum `mandatory:"true" json:"status"`

	// The date and time the job was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m SqlPlanBaselineJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaselineJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlPlanBaselineJobSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSqlPlanBaselineJobSummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineJobSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlPlanBaselineJobSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlPlanBaselineJobSummaryTypeEnum Enum with underlying type: string
type SqlPlanBaselineJobSummaryTypeEnum string

// Set of constants representing the allowable values for SqlPlanBaselineJobSummaryTypeEnum
const (
	SqlPlanBaselineJobSummaryTypeLoad SqlPlanBaselineJobSummaryTypeEnum = "LOAD"
)

var mappingSqlPlanBaselineJobSummaryTypeEnum = map[string]SqlPlanBaselineJobSummaryTypeEnum{
	"LOAD": SqlPlanBaselineJobSummaryTypeLoad,
}

var mappingSqlPlanBaselineJobSummaryTypeEnumLowerCase = map[string]SqlPlanBaselineJobSummaryTypeEnum{
	"load": SqlPlanBaselineJobSummaryTypeLoad,
}

// GetSqlPlanBaselineJobSummaryTypeEnumValues Enumerates the set of values for SqlPlanBaselineJobSummaryTypeEnum
func GetSqlPlanBaselineJobSummaryTypeEnumValues() []SqlPlanBaselineJobSummaryTypeEnum {
	values := make([]SqlPlanBaselineJobSummaryTypeEnum, 0)
	for _, v := range mappingSqlPlanBaselineJobSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineJobSummaryTypeEnumStringValues Enumerates the set of values in String for SqlPlanBaselineJobSummaryTypeEnum
func GetSqlPlanBaselineJobSummaryTypeEnumStringValues() []string {
	return []string{
		"LOAD",
	}
}

// GetMappingSqlPlanBaselineJobSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineJobSummaryTypeEnum(val string) (SqlPlanBaselineJobSummaryTypeEnum, bool) {
	enum, ok := mappingSqlPlanBaselineJobSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineJobSummaryStatusEnum Enum with underlying type: string
type SqlPlanBaselineJobSummaryStatusEnum string

// Set of constants representing the allowable values for SqlPlanBaselineJobSummaryStatusEnum
const (
	SqlPlanBaselineJobSummaryStatusSucceeded SqlPlanBaselineJobSummaryStatusEnum = "SUCCEEDED"
	SqlPlanBaselineJobSummaryStatusScheduled SqlPlanBaselineJobSummaryStatusEnum = "SCHEDULED"
	SqlPlanBaselineJobSummaryStatusFailed    SqlPlanBaselineJobSummaryStatusEnum = "FAILED"
)

var mappingSqlPlanBaselineJobSummaryStatusEnum = map[string]SqlPlanBaselineJobSummaryStatusEnum{
	"SUCCEEDED": SqlPlanBaselineJobSummaryStatusSucceeded,
	"SCHEDULED": SqlPlanBaselineJobSummaryStatusScheduled,
	"FAILED":    SqlPlanBaselineJobSummaryStatusFailed,
}

var mappingSqlPlanBaselineJobSummaryStatusEnumLowerCase = map[string]SqlPlanBaselineJobSummaryStatusEnum{
	"succeeded": SqlPlanBaselineJobSummaryStatusSucceeded,
	"scheduled": SqlPlanBaselineJobSummaryStatusScheduled,
	"failed":    SqlPlanBaselineJobSummaryStatusFailed,
}

// GetSqlPlanBaselineJobSummaryStatusEnumValues Enumerates the set of values for SqlPlanBaselineJobSummaryStatusEnum
func GetSqlPlanBaselineJobSummaryStatusEnumValues() []SqlPlanBaselineJobSummaryStatusEnum {
	values := make([]SqlPlanBaselineJobSummaryStatusEnum, 0)
	for _, v := range mappingSqlPlanBaselineJobSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineJobSummaryStatusEnumStringValues Enumerates the set of values in String for SqlPlanBaselineJobSummaryStatusEnum
func GetSqlPlanBaselineJobSummaryStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"SCHEDULED",
		"FAILED",
	}
}

// GetMappingSqlPlanBaselineJobSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineJobSummaryStatusEnum(val string) (SqlPlanBaselineJobSummaryStatusEnum, bool) {
	enum, ok := mappingSqlPlanBaselineJobSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
