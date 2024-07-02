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
	"strings"
)

// DeploymentTypeEnum Enum with underlying type: string
type DeploymentTypeEnum string

// Set of constants representing the allowable values for DeploymentTypeEnum
const (
	DeploymentTypeOnpremise  DeploymentTypeEnum = "ONPREMISE"
	DeploymentTypeBm         DeploymentTypeEnum = "BM"
	DeploymentTypeVm         DeploymentTypeEnum = "VM"
	DeploymentTypeExadata    DeploymentTypeEnum = "EXADATA"
	DeploymentTypeExadataCc  DeploymentTypeEnum = "EXADATA_CC"
	DeploymentTypeAutonomous DeploymentTypeEnum = "AUTONOMOUS"
	DeploymentTypeExadataXs  DeploymentTypeEnum = "EXADATA_XS"
)

var mappingDeploymentTypeEnum = map[string]DeploymentTypeEnum{
	"ONPREMISE":  DeploymentTypeOnpremise,
	"BM":         DeploymentTypeBm,
	"VM":         DeploymentTypeVm,
	"EXADATA":    DeploymentTypeExadata,
	"EXADATA_CC": DeploymentTypeExadataCc,
	"AUTONOMOUS": DeploymentTypeAutonomous,
	"EXADATA_XS": DeploymentTypeExadataXs,
}

var mappingDeploymentTypeEnumLowerCase = map[string]DeploymentTypeEnum{
	"onpremise":  DeploymentTypeOnpremise,
	"bm":         DeploymentTypeBm,
	"vm":         DeploymentTypeVm,
	"exadata":    DeploymentTypeExadata,
	"exadata_cc": DeploymentTypeExadataCc,
	"autonomous": DeploymentTypeAutonomous,
	"exadata_xs": DeploymentTypeExadataXs,
}

// GetDeploymentTypeEnumValues Enumerates the set of values for DeploymentTypeEnum
func GetDeploymentTypeEnumValues() []DeploymentTypeEnum {
	values := make([]DeploymentTypeEnum, 0)
	for _, v := range mappingDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentTypeEnumStringValues Enumerates the set of values in String for DeploymentTypeEnum
func GetDeploymentTypeEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"BM",
		"VM",
		"EXADATA",
		"EXADATA_CC",
		"AUTONOMOUS",
		"EXADATA_XS",
	}
}

// GetMappingDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentTypeEnum(val string) (DeploymentTypeEnum, bool) {
	enum, ok := mappingDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
