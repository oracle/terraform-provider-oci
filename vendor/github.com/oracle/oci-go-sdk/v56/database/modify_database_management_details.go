// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// ModifyDatabaseManagementDetailsManagementTypeEnum Enum with underlying type: string
type ModifyDatabaseManagementDetailsManagementTypeEnum string

// Set of constants representing the allowable values for ModifyDatabaseManagementDetailsManagementTypeEnum
const (
	ModifyDatabaseManagementDetailsManagementTypeBasic    ModifyDatabaseManagementDetailsManagementTypeEnum = "BASIC"
	ModifyDatabaseManagementDetailsManagementTypeAdvanced ModifyDatabaseManagementDetailsManagementTypeEnum = "ADVANCED"
)

var mappingModifyDatabaseManagementDetailsManagementType = map[string]ModifyDatabaseManagementDetailsManagementTypeEnum{
	"BASIC":    ModifyDatabaseManagementDetailsManagementTypeBasic,
	"ADVANCED": ModifyDatabaseManagementDetailsManagementTypeAdvanced,
}

// GetModifyDatabaseManagementDetailsManagementTypeEnumValues Enumerates the set of values for ModifyDatabaseManagementDetailsManagementTypeEnum
func GetModifyDatabaseManagementDetailsManagementTypeEnumValues() []ModifyDatabaseManagementDetailsManagementTypeEnum {
	values := make([]ModifyDatabaseManagementDetailsManagementTypeEnum, 0)
	for _, v := range mappingModifyDatabaseManagementDetailsManagementType {
		values = append(values, v)
	}
	return values
}
