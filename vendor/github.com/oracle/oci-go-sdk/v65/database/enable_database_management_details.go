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

// EnableDatabaseManagementDetails Data to enable the Database Management service for the database.
type EnableDatabaseManagementDetails struct {
	CredentialDetails *DatabaseCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	PrivateEndPointId *string `mandatory:"true" json:"privateEndPointId"`

	// The name of the Oracle Database service that will be used to connect to the database.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The Database Management type.
	ManagementType EnableDatabaseManagementDetailsManagementTypeEnum `mandatory:"false" json:"managementType,omitempty"`
}

func (m EnableDatabaseManagementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableDatabaseManagementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEnableDatabaseManagementDetailsManagementTypeEnum(string(m.ManagementType)); !ok && m.ManagementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementType: %s. Supported values are: %s.", m.ManagementType, strings.Join(GetEnableDatabaseManagementDetailsManagementTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableDatabaseManagementDetailsManagementTypeEnum Enum with underlying type: string
type EnableDatabaseManagementDetailsManagementTypeEnum string

// Set of constants representing the allowable values for EnableDatabaseManagementDetailsManagementTypeEnum
const (
	EnableDatabaseManagementDetailsManagementTypeBasic    EnableDatabaseManagementDetailsManagementTypeEnum = "BASIC"
	EnableDatabaseManagementDetailsManagementTypeAdvanced EnableDatabaseManagementDetailsManagementTypeEnum = "ADVANCED"
)

var mappingEnableDatabaseManagementDetailsManagementTypeEnum = map[string]EnableDatabaseManagementDetailsManagementTypeEnum{
	"BASIC":    EnableDatabaseManagementDetailsManagementTypeBasic,
	"ADVANCED": EnableDatabaseManagementDetailsManagementTypeAdvanced,
}

var mappingEnableDatabaseManagementDetailsManagementTypeEnumLowerCase = map[string]EnableDatabaseManagementDetailsManagementTypeEnum{
	"basic":    EnableDatabaseManagementDetailsManagementTypeBasic,
	"advanced": EnableDatabaseManagementDetailsManagementTypeAdvanced,
}

// GetEnableDatabaseManagementDetailsManagementTypeEnumValues Enumerates the set of values for EnableDatabaseManagementDetailsManagementTypeEnum
func GetEnableDatabaseManagementDetailsManagementTypeEnumValues() []EnableDatabaseManagementDetailsManagementTypeEnum {
	values := make([]EnableDatabaseManagementDetailsManagementTypeEnum, 0)
	for _, v := range mappingEnableDatabaseManagementDetailsManagementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableDatabaseManagementDetailsManagementTypeEnumStringValues Enumerates the set of values in String for EnableDatabaseManagementDetailsManagementTypeEnum
func GetEnableDatabaseManagementDetailsManagementTypeEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingEnableDatabaseManagementDetailsManagementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableDatabaseManagementDetailsManagementTypeEnum(val string) (EnableDatabaseManagementDetailsManagementTypeEnum, bool) {
	enum, ok := mappingEnableDatabaseManagementDetailsManagementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
