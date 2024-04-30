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

// DatabaseParameterUpdateStatus The result of database parameter update.
type DatabaseParameterUpdateStatus struct {

	// The status of the parameter update.
	Status DatabaseParameterUpdateStatusStatusEnum `mandatory:"false" json:"status,omitempty"`

	// An error code that defines the failure or `null` if the parameter
	// was updated successfully.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// The error message indicating the reason for failure or `null` if
	// the parameter was updated successfully.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m DatabaseParameterUpdateStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseParameterUpdateStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseParameterUpdateStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDatabaseParameterUpdateStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseParameterUpdateStatusStatusEnum Enum with underlying type: string
type DatabaseParameterUpdateStatusStatusEnum string

// Set of constants representing the allowable values for DatabaseParameterUpdateStatusStatusEnum
const (
	DatabaseParameterUpdateStatusStatusSucceeded DatabaseParameterUpdateStatusStatusEnum = "SUCCEEDED"
	DatabaseParameterUpdateStatusStatusFailed    DatabaseParameterUpdateStatusStatusEnum = "FAILED"
)

var mappingDatabaseParameterUpdateStatusStatusEnum = map[string]DatabaseParameterUpdateStatusStatusEnum{
	"SUCCEEDED": DatabaseParameterUpdateStatusStatusSucceeded,
	"FAILED":    DatabaseParameterUpdateStatusStatusFailed,
}

var mappingDatabaseParameterUpdateStatusStatusEnumLowerCase = map[string]DatabaseParameterUpdateStatusStatusEnum{
	"succeeded": DatabaseParameterUpdateStatusStatusSucceeded,
	"failed":    DatabaseParameterUpdateStatusStatusFailed,
}

// GetDatabaseParameterUpdateStatusStatusEnumValues Enumerates the set of values for DatabaseParameterUpdateStatusStatusEnum
func GetDatabaseParameterUpdateStatusStatusEnumValues() []DatabaseParameterUpdateStatusStatusEnum {
	values := make([]DatabaseParameterUpdateStatusStatusEnum, 0)
	for _, v := range mappingDatabaseParameterUpdateStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseParameterUpdateStatusStatusEnumStringValues Enumerates the set of values in String for DatabaseParameterUpdateStatusStatusEnum
func GetDatabaseParameterUpdateStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingDatabaseParameterUpdateStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseParameterUpdateStatusStatusEnum(val string) (DatabaseParameterUpdateStatusStatusEnum, bool) {
	enum, ok := mappingDatabaseParameterUpdateStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
