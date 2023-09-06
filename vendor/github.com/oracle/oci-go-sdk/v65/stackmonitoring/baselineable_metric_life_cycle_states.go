// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// BaselineableMetricLifeCycleStatesEnum Enum with underlying type: string
type BaselineableMetricLifeCycleStatesEnum string

// Set of constants representing the allowable values for BaselineableMetricLifeCycleStatesEnum
const (
	BaselineableMetricLifeCycleStatesActive  BaselineableMetricLifeCycleStatesEnum = "ACTIVE"
	BaselineableMetricLifeCycleStatesDeleted BaselineableMetricLifeCycleStatesEnum = "DELETED"
)

var mappingBaselineableMetricLifeCycleStatesEnum = map[string]BaselineableMetricLifeCycleStatesEnum{
	"ACTIVE":  BaselineableMetricLifeCycleStatesActive,
	"DELETED": BaselineableMetricLifeCycleStatesDeleted,
}

var mappingBaselineableMetricLifeCycleStatesEnumLowerCase = map[string]BaselineableMetricLifeCycleStatesEnum{
	"active":  BaselineableMetricLifeCycleStatesActive,
	"deleted": BaselineableMetricLifeCycleStatesDeleted,
}

// GetBaselineableMetricLifeCycleStatesEnumValues Enumerates the set of values for BaselineableMetricLifeCycleStatesEnum
func GetBaselineableMetricLifeCycleStatesEnumValues() []BaselineableMetricLifeCycleStatesEnum {
	values := make([]BaselineableMetricLifeCycleStatesEnum, 0)
	for _, v := range mappingBaselineableMetricLifeCycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetBaselineableMetricLifeCycleStatesEnumStringValues Enumerates the set of values in String for BaselineableMetricLifeCycleStatesEnum
func GetBaselineableMetricLifeCycleStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingBaselineableMetricLifeCycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaselineableMetricLifeCycleStatesEnum(val string) (BaselineableMetricLifeCycleStatesEnum, bool) {
	enum, ok := mappingBaselineableMetricLifeCycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
