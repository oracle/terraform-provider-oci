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

// SqlPlanBaselineJob The details of the database job used for loading and evolving SQL plan baselines.
type SqlPlanBaselineJob struct {

	// The job name.
	Name *string `mandatory:"true" json:"name"`

	// The job type.
	Type SqlPlanBaselineJobTypeEnum `mandatory:"true" json:"type"`

	// The job status.
	Status SqlPlanBaselineJobStatusEnum `mandatory:"true" json:"status"`

	// The date and time the job was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m SqlPlanBaselineJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaselineJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlPlanBaselineJobTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSqlPlanBaselineJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlPlanBaselineJobStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlPlanBaselineJobTypeEnum Enum with underlying type: string
type SqlPlanBaselineJobTypeEnum string

// Set of constants representing the allowable values for SqlPlanBaselineJobTypeEnum
const (
	SqlPlanBaselineJobTypeLoad SqlPlanBaselineJobTypeEnum = "LOAD"
)

var mappingSqlPlanBaselineJobTypeEnum = map[string]SqlPlanBaselineJobTypeEnum{
	"LOAD": SqlPlanBaselineJobTypeLoad,
}

var mappingSqlPlanBaselineJobTypeEnumLowerCase = map[string]SqlPlanBaselineJobTypeEnum{
	"load": SqlPlanBaselineJobTypeLoad,
}

// GetSqlPlanBaselineJobTypeEnumValues Enumerates the set of values for SqlPlanBaselineJobTypeEnum
func GetSqlPlanBaselineJobTypeEnumValues() []SqlPlanBaselineJobTypeEnum {
	values := make([]SqlPlanBaselineJobTypeEnum, 0)
	for _, v := range mappingSqlPlanBaselineJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineJobTypeEnumStringValues Enumerates the set of values in String for SqlPlanBaselineJobTypeEnum
func GetSqlPlanBaselineJobTypeEnumStringValues() []string {
	return []string{
		"LOAD",
	}
}

// GetMappingSqlPlanBaselineJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineJobTypeEnum(val string) (SqlPlanBaselineJobTypeEnum, bool) {
	enum, ok := mappingSqlPlanBaselineJobTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineJobStatusEnum Enum with underlying type: string
type SqlPlanBaselineJobStatusEnum string

// Set of constants representing the allowable values for SqlPlanBaselineJobStatusEnum
const (
	SqlPlanBaselineJobStatusSucceeded SqlPlanBaselineJobStatusEnum = "SUCCEEDED"
	SqlPlanBaselineJobStatusScheduled SqlPlanBaselineJobStatusEnum = "SCHEDULED"
	SqlPlanBaselineJobStatusFailed    SqlPlanBaselineJobStatusEnum = "FAILED"
)

var mappingSqlPlanBaselineJobStatusEnum = map[string]SqlPlanBaselineJobStatusEnum{
	"SUCCEEDED": SqlPlanBaselineJobStatusSucceeded,
	"SCHEDULED": SqlPlanBaselineJobStatusScheduled,
	"FAILED":    SqlPlanBaselineJobStatusFailed,
}

var mappingSqlPlanBaselineJobStatusEnumLowerCase = map[string]SqlPlanBaselineJobStatusEnum{
	"succeeded": SqlPlanBaselineJobStatusSucceeded,
	"scheduled": SqlPlanBaselineJobStatusScheduled,
	"failed":    SqlPlanBaselineJobStatusFailed,
}

// GetSqlPlanBaselineJobStatusEnumValues Enumerates the set of values for SqlPlanBaselineJobStatusEnum
func GetSqlPlanBaselineJobStatusEnumValues() []SqlPlanBaselineJobStatusEnum {
	values := make([]SqlPlanBaselineJobStatusEnum, 0)
	for _, v := range mappingSqlPlanBaselineJobStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineJobStatusEnumStringValues Enumerates the set of values in String for SqlPlanBaselineJobStatusEnum
func GetSqlPlanBaselineJobStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"SCHEDULED",
		"FAILED",
	}
}

// GetMappingSqlPlanBaselineJobStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineJobStatusEnum(val string) (SqlPlanBaselineJobStatusEnum, bool) {
	enum, ok := mappingSqlPlanBaselineJobStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
