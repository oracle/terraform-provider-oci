// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Control Plane API
//
// Use the Application Performance Monitoring Control Plane API to perform operations such as creating, updating,
// deleting and listing APM domains and monitoring the progress of these operations using the work request APIs. For more information, see Application Performance Monitoring (https://docs.cloud.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmcontrolplane

import (
	"strings"
)

// DataKeyTypesEnum Enum with underlying type: string
type DataKeyTypesEnum string

// Set of constants representing the allowable values for DataKeyTypesEnum
const (
	DataKeyTypesPrivate DataKeyTypesEnum = "PRIVATE"
	DataKeyTypesPublic  DataKeyTypesEnum = "PUBLIC"
)

var mappingDataKeyTypesEnum = map[string]DataKeyTypesEnum{
	"PRIVATE": DataKeyTypesPrivate,
	"PUBLIC":  DataKeyTypesPublic,
}

var mappingDataKeyTypesEnumLowerCase = map[string]DataKeyTypesEnum{
	"private": DataKeyTypesPrivate,
	"public":  DataKeyTypesPublic,
}

// GetDataKeyTypesEnumValues Enumerates the set of values for DataKeyTypesEnum
func GetDataKeyTypesEnumValues() []DataKeyTypesEnum {
	values := make([]DataKeyTypesEnum, 0)
	for _, v := range mappingDataKeyTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDataKeyTypesEnumStringValues Enumerates the set of values in String for DataKeyTypesEnum
func GetDataKeyTypesEnumStringValues() []string {
	return []string{
		"PRIVATE",
		"PUBLIC",
	}
}

// GetMappingDataKeyTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataKeyTypesEnum(val string) (DataKeyTypesEnum, bool) {
	enum, ok := mappingDataKeyTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
