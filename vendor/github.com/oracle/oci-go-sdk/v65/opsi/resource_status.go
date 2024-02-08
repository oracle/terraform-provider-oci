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

// ResourceStatusEnum Enum with underlying type: string
type ResourceStatusEnum string

// Set of constants representing the allowable values for ResourceStatusEnum
const (
	ResourceStatusDisabled   ResourceStatusEnum = "DISABLED"
	ResourceStatusEnabled    ResourceStatusEnum = "ENABLED"
	ResourceStatusTerminated ResourceStatusEnum = "TERMINATED"
)

var mappingResourceStatusEnum = map[string]ResourceStatusEnum{
	"DISABLED":   ResourceStatusDisabled,
	"ENABLED":    ResourceStatusEnabled,
	"TERMINATED": ResourceStatusTerminated,
}

var mappingResourceStatusEnumLowerCase = map[string]ResourceStatusEnum{
	"disabled":   ResourceStatusDisabled,
	"enabled":    ResourceStatusEnabled,
	"terminated": ResourceStatusTerminated,
}

// GetResourceStatusEnumValues Enumerates the set of values for ResourceStatusEnum
func GetResourceStatusEnumValues() []ResourceStatusEnum {
	values := make([]ResourceStatusEnum, 0)
	for _, v := range mappingResourceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceStatusEnumStringValues Enumerates the set of values in String for ResourceStatusEnum
func GetResourceStatusEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENABLED",
		"TERMINATED",
	}
}

// GetMappingResourceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceStatusEnum(val string) (ResourceStatusEnum, bool) {
	enum, ok := mappingResourceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
