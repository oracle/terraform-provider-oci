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

// TablespaceAdminStatus The status of a tablespace admin action.
type TablespaceAdminStatus struct {

	// The status of a tablespace admin action.
	Status TablespaceAdminStatusStatusEnum `mandatory:"true" json:"status"`

	// The error code that denotes failure if the tablespace admin action is not successful. The error code is "null" if the admin action is successful.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// The error message that indicates the reason for failure if the tablespace admin action is not successful. The error message is "null" if the admin action is successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m TablespaceAdminStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TablespaceAdminStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTablespaceAdminStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTablespaceAdminStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TablespaceAdminStatusStatusEnum Enum with underlying type: string
type TablespaceAdminStatusStatusEnum string

// Set of constants representing the allowable values for TablespaceAdminStatusStatusEnum
const (
	TablespaceAdminStatusStatusSucceeded TablespaceAdminStatusStatusEnum = "SUCCEEDED"
	TablespaceAdminStatusStatusFailed    TablespaceAdminStatusStatusEnum = "FAILED"
)

var mappingTablespaceAdminStatusStatusEnum = map[string]TablespaceAdminStatusStatusEnum{
	"SUCCEEDED": TablespaceAdminStatusStatusSucceeded,
	"FAILED":    TablespaceAdminStatusStatusFailed,
}

var mappingTablespaceAdminStatusStatusEnumLowerCase = map[string]TablespaceAdminStatusStatusEnum{
	"succeeded": TablespaceAdminStatusStatusSucceeded,
	"failed":    TablespaceAdminStatusStatusFailed,
}

// GetTablespaceAdminStatusStatusEnumValues Enumerates the set of values for TablespaceAdminStatusStatusEnum
func GetTablespaceAdminStatusStatusEnumValues() []TablespaceAdminStatusStatusEnum {
	values := make([]TablespaceAdminStatusStatusEnum, 0)
	for _, v := range mappingTablespaceAdminStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceAdminStatusStatusEnumStringValues Enumerates the set of values in String for TablespaceAdminStatusStatusEnum
func GetTablespaceAdminStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingTablespaceAdminStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceAdminStatusStatusEnum(val string) (TablespaceAdminStatusStatusEnum, bool) {
	enum, ok := mappingTablespaceAdminStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
