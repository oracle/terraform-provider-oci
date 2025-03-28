// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ExadataResourceTypeEnum Enum with underlying type: string
type ExadataResourceTypeEnum string

// Set of constants representing the allowable values for ExadataResourceTypeEnum
const (
	ExadataResourceTypeCloudExadataInfrastructure ExadataResourceTypeEnum = "cloudExadataInfrastructure"
	ExadataResourceTypeExadataInfrastructure      ExadataResourceTypeEnum = "exadataInfrastructure"
)

var mappingExadataResourceTypeEnum = map[string]ExadataResourceTypeEnum{
	"cloudExadataInfrastructure": ExadataResourceTypeCloudExadataInfrastructure,
	"exadataInfrastructure":      ExadataResourceTypeExadataInfrastructure,
}

var mappingExadataResourceTypeEnumLowerCase = map[string]ExadataResourceTypeEnum{
	"cloudexadatainfrastructure": ExadataResourceTypeCloudExadataInfrastructure,
	"exadatainfrastructure":      ExadataResourceTypeExadataInfrastructure,
}

// GetExadataResourceTypeEnumValues Enumerates the set of values for ExadataResourceTypeEnum
func GetExadataResourceTypeEnumValues() []ExadataResourceTypeEnum {
	values := make([]ExadataResourceTypeEnum, 0)
	for _, v := range mappingExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataResourceTypeEnumStringValues Enumerates the set of values in String for ExadataResourceTypeEnum
func GetExadataResourceTypeEnumStringValues() []string {
	return []string{
		"cloudExadataInfrastructure",
		"exadataInfrastructure",
	}
}

// GetMappingExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataResourceTypeEnum(val string) (ExadataResourceTypeEnum, bool) {
	enum, ok := mappingExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
