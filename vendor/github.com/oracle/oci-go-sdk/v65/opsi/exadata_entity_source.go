// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ExadataEntitySourceEnum Enum with underlying type: string
type ExadataEntitySourceEnum string

// Set of constants representing the allowable values for ExadataEntitySourceEnum
const (
	ExadataEntitySourceEmManagedExternalExadata ExadataEntitySourceEnum = "EM_MANAGED_EXTERNAL_EXADATA"
	ExadataEntitySourcePeComanagedExadata       ExadataEntitySourceEnum = "PE_COMANAGED_EXADATA"
)

var mappingExadataEntitySourceEnum = map[string]ExadataEntitySourceEnum{
	"EM_MANAGED_EXTERNAL_EXADATA": ExadataEntitySourceEmManagedExternalExadata,
	"PE_COMANAGED_EXADATA":        ExadataEntitySourcePeComanagedExadata,
}

var mappingExadataEntitySourceEnumLowerCase = map[string]ExadataEntitySourceEnum{
	"em_managed_external_exadata": ExadataEntitySourceEmManagedExternalExadata,
	"pe_comanaged_exadata":        ExadataEntitySourcePeComanagedExadata,
}

// GetExadataEntitySourceEnumValues Enumerates the set of values for ExadataEntitySourceEnum
func GetExadataEntitySourceEnumValues() []ExadataEntitySourceEnum {
	values := make([]ExadataEntitySourceEnum, 0)
	for _, v := range mappingExadataEntitySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataEntitySourceEnumStringValues Enumerates the set of values in String for ExadataEntitySourceEnum
func GetExadataEntitySourceEnumStringValues() []string {
	return []string{
		"EM_MANAGED_EXTERNAL_EXADATA",
		"PE_COMANAGED_EXADATA",
	}
}

// GetMappingExadataEntitySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataEntitySourceEnum(val string) (ExadataEntitySourceEnum, bool) {
	enum, ok := mappingExadataEntitySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
