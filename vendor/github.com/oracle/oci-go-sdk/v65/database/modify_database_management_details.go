// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModifyDatabaseManagementDetails Data to update one or more attributes of the Database Management configuration for the database.
type ModifyDatabaseManagementDetails struct {
	CredentialDetails *DatabaseCredentialDetails `mandatory:"false" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	PrivateEndPointId *string `mandatory:"false" json:"privateEndPointId"`

	// The Database Management type.
	ManagementType ModifyDatabaseManagementDetailsManagementTypeEnum `mandatory:"false" json:"managementType,omitempty"`

	// The name of the Oracle Database service that will be used to connect to the database.
	ServiceName *string `mandatory:"false" json:"serviceName"`
}

func (m ModifyDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModifyDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingModifyDatabaseManagementDetailsManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModifyDatabaseManagementDetailsManagementTypeEnum Enum with underlying type: string
type ModifyDatabaseManagementDetailsManagementTypeEnum string

// Set of constants representing the allowable values for ModifyDatabaseManagementDetailsManagementTypeEnum
const (
	ModifyDatabaseManagementDetailsManagementTypeBasic    ModifyDatabaseManagementDetailsManagementTypeEnum = "BASIC"
	ModifyDatabaseManagementDetailsManagementTypeAdvanced ModifyDatabaseManagementDetailsManagementTypeEnum = "ADVANCED"
)

var mappingModifyDatabaseManagementDetailsManagementTypeEnum = map[string]ModifyDatabaseManagementDetailsManagementTypeEnum{
	"BASIC":    ModifyDatabaseManagementDetailsManagementTypeBasic,
	"ADVANCED": ModifyDatabaseManagementDetailsManagementTypeAdvanced,
}

var mappingModifyDatabaseManagementDetailsManagementTypeEnumLowerCase = map[string]ModifyDatabaseManagementDetailsManagementTypeEnum{
	"basic":    ModifyDatabaseManagementDetailsManagementTypeBasic,
	"advanced": ModifyDatabaseManagementDetailsManagementTypeAdvanced,
}

// GetModifyDatabaseManagementDetailsManagementTypeEnumValues Enumerates the set of values for ModifyDatabaseManagementDetailsManagementTypeEnum
func GetModifyDatabaseManagementDetailsManagementTypeEnumValues() []ModifyDatabaseManagementDetailsManagementTypeEnum {
	values := make([]ModifyDatabaseManagementDetailsManagementTypeEnum, 0)
	for _, v := range mappingModifyDatabaseManagementDetailsManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues Enumerates the set of values in String for ModifyDatabaseManagementDetailsManagementTypeEnum
func GetModifyDatabaseManagementDetailsManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingModifyDatabaseManagementDetailsManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyDatabaseManagementDetailsManagementTypeEnum(val string) (ModifyDatabaseManagementDetailsManagementTypeEnum, bool) {
	enum, ok := mappingModifyDatabaseManagementDetailsManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
