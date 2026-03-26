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

// CloudExadataInfrastructureDeploymentTypeEnum Enum with underlying type: string
type CloudExadataInfrastructureDeploymentTypeEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureDeploymentTypeEnum
const (
	CloudExadataInfrastructureDeploymentTypeExadata   CloudExadataInfrastructureDeploymentTypeEnum = "EXADATA"
	CloudExadataInfrastructureDeploymentTypeExadataCc CloudExadataInfrastructureDeploymentTypeEnum = "EXADATA_CC"
)

var mappingCloudExadataInfrastructureDeploymentTypeEnum = map[string]CloudExadataInfrastructureDeploymentTypeEnum{
	"EXADATA":    CloudExadataInfrastructureDeploymentTypeExadata,
	"EXADATA_CC": CloudExadataInfrastructureDeploymentTypeExadataCc,
}

var mappingCloudExadataInfrastructureDeploymentTypeEnumLowerCase = map[string]CloudExadataInfrastructureDeploymentTypeEnum{
	"exadata":    CloudExadataInfrastructureDeploymentTypeExadata,
	"exadata_cc": CloudExadataInfrastructureDeploymentTypeExadataCc,
}

// GetCloudExadataInfrastructureDeploymentTypeEnumValues Enumerates the set of values for CloudExadataInfrastructureDeploymentTypeEnum
func GetCloudExadataInfrastructureDeploymentTypeEnumValues() []CloudExadataInfrastructureDeploymentTypeEnum {
	values := make([]CloudExadataInfrastructureDeploymentTypeEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureDeploymentTypeEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureDeploymentTypeEnum
func GetCloudExadataInfrastructureDeploymentTypeEnumStringValues() []string {
	return []string{
		"EXADATA",
		"EXADATA_CC",
	}
}

// GetMappingCloudExadataInfrastructureDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureDeploymentTypeEnum(val string) (CloudExadataInfrastructureDeploymentTypeEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
