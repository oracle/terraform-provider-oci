// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// SchedulingPolicyEnum Enum with underlying type: string
type SchedulingPolicyEnum string

// Set of constants representing the allowable values for SchedulingPolicyEnum
const (
	SchedulingPolicyAll               SchedulingPolicyEnum = "ALL"
	SchedulingPolicyRoundRobin        SchedulingPolicyEnum = "ROUND_ROBIN"
	SchedulingPolicyBatchedRoundRobin SchedulingPolicyEnum = "BATCHED_ROUND_ROBIN"
)

var mappingSchedulingPolicyEnum = map[string]SchedulingPolicyEnum{
	"ALL":                 SchedulingPolicyAll,
	"ROUND_ROBIN":         SchedulingPolicyRoundRobin,
	"BATCHED_ROUND_ROBIN": SchedulingPolicyBatchedRoundRobin,
}

var mappingSchedulingPolicyEnumLowerCase = map[string]SchedulingPolicyEnum{
	"all":                 SchedulingPolicyAll,
	"round_robin":         SchedulingPolicyRoundRobin,
	"batched_round_robin": SchedulingPolicyBatchedRoundRobin,
}

// GetSchedulingPolicyEnumValues Enumerates the set of values for SchedulingPolicyEnum
func GetSchedulingPolicyEnumValues() []SchedulingPolicyEnum {
	values := make([]SchedulingPolicyEnum, 0)
	for _, v := range mappingSchedulingPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulingPolicyEnumStringValues Enumerates the set of values in String for SchedulingPolicyEnum
func GetSchedulingPolicyEnumStringValues() []string {
	return []string{
		"ALL",
		"ROUND_ROBIN",
		"BATCHED_ROUND_ROBIN",
	}
}

// GetMappingSchedulingPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulingPolicyEnum(val string) (SchedulingPolicyEnum, bool) {
	enum, ok := mappingSchedulingPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
