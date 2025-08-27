// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeMysqlDatabaseManagementTypeDetails Details to change managed MySQL database management type.
type ChangeMysqlDatabaseManagementTypeDetails struct {

	// The type of HeatWave management.
	ManagementType ManagedMySqlDatabaseHeatWaveManagementTypeEnum `mandatory:"true" json:"managementType"`

	// The type of operation to perform: update managementType, enable or disable database management.
	Operation ChangeMysqlDatabaseManagementTypeDetailsOperationEnum `mandatory:"false" json:"operation,omitempty"`
}

func (m ChangeMysqlDatabaseManagementTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeMysqlDatabaseManagementTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedMySqlDatabaseHeatWaveManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetManagedMySqlDatabaseHeatWaveManagementTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingChangeMysqlDatabaseManagementTypeDetailsOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetChangeMysqlDatabaseManagementTypeDetailsOperationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeMysqlDatabaseManagementTypeDetailsOperationEnum Enum with underlying type: string
type ChangeMysqlDatabaseManagementTypeDetailsOperationEnum string

// Set of constants representing the allowable values for ChangeMysqlDatabaseManagementTypeDetailsOperationEnum
const (
	ChangeMysqlDatabaseManagementTypeDetailsOperationEnableDbmgmt     ChangeMysqlDatabaseManagementTypeDetailsOperationEnum = "ENABLE_DBMGMT"
	ChangeMysqlDatabaseManagementTypeDetailsOperationUpdateDbmgmtType ChangeMysqlDatabaseManagementTypeDetailsOperationEnum = "UPDATE_DBMGMT_TYPE"
	ChangeMysqlDatabaseManagementTypeDetailsOperationDisableDbmgmt    ChangeMysqlDatabaseManagementTypeDetailsOperationEnum = "DISABLE_DBMGMT"
)

var mappingChangeMysqlDatabaseManagementTypeDetailsOperationEnum = map[string]ChangeMysqlDatabaseManagementTypeDetailsOperationEnum{
	"ENABLE_DBMGMT":      ChangeMysqlDatabaseManagementTypeDetailsOperationEnableDbmgmt,
	"UPDATE_DBMGMT_TYPE": ChangeMysqlDatabaseManagementTypeDetailsOperationUpdateDbmgmtType,
	"DISABLE_DBMGMT":     ChangeMysqlDatabaseManagementTypeDetailsOperationDisableDbmgmt,
}

var mappingChangeMysqlDatabaseManagementTypeDetailsOperationEnumLowerCase = map[string]ChangeMysqlDatabaseManagementTypeDetailsOperationEnum{
	"enable_dbmgmt":      ChangeMysqlDatabaseManagementTypeDetailsOperationEnableDbmgmt,
	"update_dbmgmt_type": ChangeMysqlDatabaseManagementTypeDetailsOperationUpdateDbmgmtType,
	"disable_dbmgmt":     ChangeMysqlDatabaseManagementTypeDetailsOperationDisableDbmgmt,
}

// GetChangeMysqlDatabaseManagementTypeDetailsOperationEnumValues Enumerates the set of values for ChangeMysqlDatabaseManagementTypeDetailsOperationEnum
func GetChangeMysqlDatabaseManagementTypeDetailsOperationEnumValues() []ChangeMysqlDatabaseManagementTypeDetailsOperationEnum {
	values := make([]ChangeMysqlDatabaseManagementTypeDetailsOperationEnum, 0)
	for _, v := range mappingChangeMysqlDatabaseManagementTypeDetailsOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeMysqlDatabaseManagementTypeDetailsOperationEnumStringValues Enumerates the set of values in String for ChangeMysqlDatabaseManagementTypeDetailsOperationEnum
func GetChangeMysqlDatabaseManagementTypeDetailsOperationEnumStringValues() []string {
	return []string{
		"ENABLE_DBMGMT",
		"UPDATE_DBMGMT_TYPE",
		"DISABLE_DBMGMT",
	}
}

// GetMappingChangeMysqlDatabaseManagementTypeDetailsOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeMysqlDatabaseManagementTypeDetailsOperationEnum(val string) (ChangeMysqlDatabaseManagementTypeDetailsOperationEnum, bool) {
	enum, ok := mappingChangeMysqlDatabaseManagementTypeDetailsOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
