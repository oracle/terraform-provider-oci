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

// EnableDatabaseManagementDetailsManagementTypeEnum Enum with underlying type: string
type EnableDatabaseManagementDetailsManagementTypeEnum string

// Set of constants representing the allowable values for EnableDatabaseManagementDetailsManagementTypeEnum
const (
	EnableDatabaseManagementDetailsManagementTypeBasic    EnableDatabaseManagementDetailsManagementTypeEnum = "BASIC"
	EnableDatabaseManagementDetailsManagementTypeAdvanced EnableDatabaseManagementDetailsManagementTypeEnum = "ADVANCED"
)

var mappingEnableDatabaseManagementDetailsManagementType = map[string]EnableDatabaseManagementDetailsManagementTypeEnum{
	"BASIC":    EnableDatabaseManagementDetailsManagementTypeBasic,
	"ADVANCED": EnableDatabaseManagementDetailsManagementTypeAdvanced,
}

// GetEnableDatabaseManagementDetailsManagementTypeEnumValues Enumerates the set of values for EnableDatabaseManagementDetailsManagementTypeEnum
func GetEnableDatabaseManagementDetailsManagementTypeEnumValues() []EnableDatabaseManagementDetailsManagementTypeEnum {
	values := make([]EnableDatabaseManagementDetailsManagementTypeEnum, 0)
	for _, v := range mappingEnableDatabaseManagementDetailsManagementType {
		values = append(values, v)
	}
	return values
}
