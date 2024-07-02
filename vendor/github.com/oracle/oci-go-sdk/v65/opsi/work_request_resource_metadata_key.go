// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// WorkRequestResourceMetadataKeyEnum Enum with underlying type: string
type WorkRequestResourceMetadataKeyEnum string

// Set of constants representing the allowable values for WorkRequestResourceMetadataKeyEnum
const (
	WorkRequestResourceMetadataKeyRelatedResourceIds WorkRequestResourceMetadataKeyEnum = "RELATED_RESOURCE_IDS"
	WorkRequestResourceMetadataKeyResourceName       WorkRequestResourceMetadataKeyEnum = "RESOURCE_NAME"
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"RELATED_RESOURCE_IDS": WorkRequestResourceMetadataKeyRelatedResourceIds,
	"RESOURCE_NAME":        WorkRequestResourceMetadataKeyResourceName,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"related_resource_ids": WorkRequestResourceMetadataKeyRelatedResourceIds,
	"resource_name":        WorkRequestResourceMetadataKeyResourceName,
}

// GetWorkRequestResourceMetadataKeyEnumValues Enumerates the set of values for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumValues() []WorkRequestResourceMetadataKeyEnum {
	values := make([]WorkRequestResourceMetadataKeyEnum, 0)
	for _, v := range mappingWorkRequestResourceMetadataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceMetadataKeyEnumStringValues Enumerates the set of values in String for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumStringValues() []string {
	return []string{
		"RELATED_RESOURCE_IDS",
		"RESOURCE_NAME",
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
