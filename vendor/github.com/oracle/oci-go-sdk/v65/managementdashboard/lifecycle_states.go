// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

import (
	"strings"
)

// LifecycleStatesEnum Enum with underlying type: string
type LifecycleStatesEnum string

// Set of constants representing the allowable values for LifecycleStatesEnum
const (
	LifecycleStatesActive LifecycleStatesEnum = "ACTIVE"
)

var mappingLifecycleStatesEnum = map[string]LifecycleStatesEnum{
	"ACTIVE": LifecycleStatesActive,
}

var mappingLifecycleStatesEnumLowerCase = map[string]LifecycleStatesEnum{
	"active": LifecycleStatesActive,
}

// GetLifecycleStatesEnumValues Enumerates the set of values for LifecycleStatesEnum
func GetLifecycleStatesEnumValues() []LifecycleStatesEnum {
	values := make([]LifecycleStatesEnum, 0)
	for _, v := range mappingLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStatesEnumStringValues Enumerates the set of values in String for LifecycleStatesEnum
func GetLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStatesEnum(val string) (LifecycleStatesEnum, bool) {
	enum, ok := mappingLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
