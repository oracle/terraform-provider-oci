// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ExadataInfrastructureDeploymentTypeEnum Enum with underlying type: string
type ExadataInfrastructureDeploymentTypeEnum string

// Set of constants representing the allowable values for ExadataInfrastructureDeploymentTypeEnum
const (
	ExadataInfrastructureDeploymentTypeOnpremise ExadataInfrastructureDeploymentTypeEnum = "ONPREMISE"
	ExadataInfrastructureDeploymentTypeExadata   ExadataInfrastructureDeploymentTypeEnum = "EXADATA"
	ExadataInfrastructureDeploymentTypeExadataCc ExadataInfrastructureDeploymentTypeEnum = "EXADATA_CC"
)

var mappingExadataInfrastructureDeploymentTypeEnum = map[string]ExadataInfrastructureDeploymentTypeEnum{
	"ONPREMISE":  ExadataInfrastructureDeploymentTypeOnpremise,
	"EXADATA":    ExadataInfrastructureDeploymentTypeExadata,
	"EXADATA_CC": ExadataInfrastructureDeploymentTypeExadataCc,
}

var mappingExadataInfrastructureDeploymentTypeEnumLowerCase = map[string]ExadataInfrastructureDeploymentTypeEnum{
	"onpremise":  ExadataInfrastructureDeploymentTypeOnpremise,
	"exadata":    ExadataInfrastructureDeploymentTypeExadata,
	"exadata_cc": ExadataInfrastructureDeploymentTypeExadataCc,
}

// GetExadataInfrastructureDeploymentTypeEnumValues Enumerates the set of values for ExadataInfrastructureDeploymentTypeEnum
func GetExadataInfrastructureDeploymentTypeEnumValues() []ExadataInfrastructureDeploymentTypeEnum {
	values := make([]ExadataInfrastructureDeploymentTypeEnum, 0)
	for _, v := range mappingExadataInfrastructureDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureDeploymentTypeEnumStringValues Enumerates the set of values in String for ExadataInfrastructureDeploymentTypeEnum
func GetExadataInfrastructureDeploymentTypeEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"EXADATA",
		"EXADATA_CC",
	}
}

// GetMappingExadataInfrastructureDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureDeploymentTypeEnum(val string) (ExadataInfrastructureDeploymentTypeEnum, bool) {
	enum, ok := mappingExadataInfrastructureDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
