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
	"strings"
)

// CloudDbSystemDeploymentTypeEnum Enum with underlying type: string
type CloudDbSystemDeploymentTypeEnum string

// Set of constants representing the allowable values for CloudDbSystemDeploymentTypeEnum
const (
	CloudDbSystemDeploymentTypeVm    CloudDbSystemDeploymentTypeEnum = "VM"
	CloudDbSystemDeploymentTypeExacs CloudDbSystemDeploymentTypeEnum = "EXACS"
	CloudDbSystemDeploymentTypeExacc CloudDbSystemDeploymentTypeEnum = "EXACC"
)

var mappingCloudDbSystemDeploymentTypeEnum = map[string]CloudDbSystemDeploymentTypeEnum{
	"VM":    CloudDbSystemDeploymentTypeVm,
	"EXACS": CloudDbSystemDeploymentTypeExacs,
	"EXACC": CloudDbSystemDeploymentTypeExacc,
}

var mappingCloudDbSystemDeploymentTypeEnumLowerCase = map[string]CloudDbSystemDeploymentTypeEnum{
	"vm":    CloudDbSystemDeploymentTypeVm,
	"exacs": CloudDbSystemDeploymentTypeExacs,
	"exacc": CloudDbSystemDeploymentTypeExacc,
}

// GetCloudDbSystemDeploymentTypeEnumValues Enumerates the set of values for CloudDbSystemDeploymentTypeEnum
func GetCloudDbSystemDeploymentTypeEnumValues() []CloudDbSystemDeploymentTypeEnum {
	values := make([]CloudDbSystemDeploymentTypeEnum, 0)
	for _, v := range mappingCloudDbSystemDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemDeploymentTypeEnumStringValues Enumerates the set of values in String for CloudDbSystemDeploymentTypeEnum
func GetCloudDbSystemDeploymentTypeEnumStringValues() []string {
	return []string{
		"VM",
		"EXACS",
		"EXACC",
	}
}

// GetMappingCloudDbSystemDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemDeploymentTypeEnum(val string) (CloudDbSystemDeploymentTypeEnum, bool) {
	enum, ok := mappingCloudDbSystemDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
