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

// AwrHubSourceStatusEnum Enum with underlying type: string
type AwrHubSourceStatusEnum string

// Set of constants representing the allowable values for AwrHubSourceStatusEnum
const (
	AwrHubSourceStatusAccepting     AwrHubSourceStatusEnum = "ACCEPTING"
	AwrHubSourceStatusNotAccepting  AwrHubSourceStatusEnum = "NOT_ACCEPTING"
	AwrHubSourceStatusNotRegistered AwrHubSourceStatusEnum = "NOT_REGISTERED"
	AwrHubSourceStatusTerminated    AwrHubSourceStatusEnum = "TERMINATED"
)

var mappingAwrHubSourceStatusEnum = map[string]AwrHubSourceStatusEnum{
	"ACCEPTING":      AwrHubSourceStatusAccepting,
	"NOT_ACCEPTING":  AwrHubSourceStatusNotAccepting,
	"NOT_REGISTERED": AwrHubSourceStatusNotRegistered,
	"TERMINATED":     AwrHubSourceStatusTerminated,
}

var mappingAwrHubSourceStatusEnumLowerCase = map[string]AwrHubSourceStatusEnum{
	"accepting":      AwrHubSourceStatusAccepting,
	"not_accepting":  AwrHubSourceStatusNotAccepting,
	"not_registered": AwrHubSourceStatusNotRegistered,
	"terminated":     AwrHubSourceStatusTerminated,
}

// GetAwrHubSourceStatusEnumValues Enumerates the set of values for AwrHubSourceStatusEnum
func GetAwrHubSourceStatusEnumValues() []AwrHubSourceStatusEnum {
	values := make([]AwrHubSourceStatusEnum, 0)
	for _, v := range mappingAwrHubSourceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrHubSourceStatusEnumStringValues Enumerates the set of values in String for AwrHubSourceStatusEnum
func GetAwrHubSourceStatusEnumStringValues() []string {
	return []string{
		"ACCEPTING",
		"NOT_ACCEPTING",
		"NOT_REGISTERED",
		"TERMINATED",
	}
}

// GetMappingAwrHubSourceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrHubSourceStatusEnum(val string) (AwrHubSourceStatusEnum, bool) {
	enum, ok := mappingAwrHubSourceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
