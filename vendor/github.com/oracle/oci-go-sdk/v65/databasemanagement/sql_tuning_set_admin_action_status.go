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

// SqlTuningSetAdminActionStatus The status of a Sql tuning set admin action.
type SqlTuningSetAdminActionStatus struct {

	// The status of a Sql tuning set admin action.
	Status SqlTuningSetAdminActionStatusStatusEnum `mandatory:"true" json:"status"`

	// The success message of the Sql tuning set admin action. The success message is "null" if the admin action is non successful.
	SuccessMessage *string `mandatory:"false" json:"successMessage"`

	// The error code that denotes failure if the Sql tuning set admin action is not successful. The error code is "null" if the admin action is successful.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// The error message that indicates the reason for failure if the Sql tuning set admin action is not successful. The error message is "null" if the admin action is successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Flag to indicate whether to create the Sql tuning set or just display the plsql used for the selected user action.
	ShowSqlOnly *int `mandatory:"false" json:"showSqlOnly"`

	// When showSqlOnly is set to 1, this attribute displays the plsql generated for the selected user action.
	// When showSqlOnly is set to 0, this attribute will not be returned.
	SqlStatement *string `mandatory:"false" json:"sqlStatement"`
}

func (m SqlTuningSetAdminActionStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningSetAdminActionStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlTuningSetAdminActionStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlTuningSetAdminActionStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningSetAdminActionStatusStatusEnum Enum with underlying type: string
type SqlTuningSetAdminActionStatusStatusEnum string

// Set of constants representing the allowable values for SqlTuningSetAdminActionStatusStatusEnum
const (
	SqlTuningSetAdminActionStatusStatusSucceeded SqlTuningSetAdminActionStatusStatusEnum = "SUCCEEDED"
	SqlTuningSetAdminActionStatusStatusFailed    SqlTuningSetAdminActionStatusStatusEnum = "FAILED"
)

var mappingSqlTuningSetAdminActionStatusStatusEnum = map[string]SqlTuningSetAdminActionStatusStatusEnum{
	"SUCCEEDED": SqlTuningSetAdminActionStatusStatusSucceeded,
	"FAILED":    SqlTuningSetAdminActionStatusStatusFailed,
}

var mappingSqlTuningSetAdminActionStatusStatusEnumLowerCase = map[string]SqlTuningSetAdminActionStatusStatusEnum{
	"succeeded": SqlTuningSetAdminActionStatusStatusSucceeded,
	"failed":    SqlTuningSetAdminActionStatusStatusFailed,
}

// GetSqlTuningSetAdminActionStatusStatusEnumValues Enumerates the set of values for SqlTuningSetAdminActionStatusStatusEnum
func GetSqlTuningSetAdminActionStatusStatusEnumValues() []SqlTuningSetAdminActionStatusStatusEnum {
	values := make([]SqlTuningSetAdminActionStatusStatusEnum, 0)
	for _, v := range mappingSqlTuningSetAdminActionStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningSetAdminActionStatusStatusEnumStringValues Enumerates the set of values in String for SqlTuningSetAdminActionStatusStatusEnum
func GetSqlTuningSetAdminActionStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingSqlTuningSetAdminActionStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningSetAdminActionStatusStatusEnum(val string) (SqlTuningSetAdminActionStatusStatusEnum, bool) {
	enum, ok := mappingSqlTuningSetAdminActionStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
