// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content Management API
//
// Oracle Content Management is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"strings"
)

// LifecycleDetailsEnum Enum with underlying type: string
type LifecycleDetailsEnum string

// Set of constants representing the allowable values for LifecycleDetailsEnum
const (
	LifecycleDetailsStandby  LifecycleDetailsEnum = "STANDBY"
	LifecycleDetailsFailover LifecycleDetailsEnum = "FAILOVER"
	LifecycleDetailsDown     LifecycleDetailsEnum = "DOWN"
	LifecycleDetailsPrimary  LifecycleDetailsEnum = "PRIMARY"
)

var mappingLifecycleDetailsEnum = map[string]LifecycleDetailsEnum{
	"STANDBY":  LifecycleDetailsStandby,
	"FAILOVER": LifecycleDetailsFailover,
	"DOWN":     LifecycleDetailsDown,
	"PRIMARY":  LifecycleDetailsPrimary,
}

var mappingLifecycleDetailsEnumLowerCase = map[string]LifecycleDetailsEnum{
	"standby":  LifecycleDetailsStandby,
	"failover": LifecycleDetailsFailover,
	"down":     LifecycleDetailsDown,
	"primary":  LifecycleDetailsPrimary,
}

// GetLifecycleDetailsEnumValues Enumerates the set of values for LifecycleDetailsEnum
func GetLifecycleDetailsEnumValues() []LifecycleDetailsEnum {
	values := make([]LifecycleDetailsEnum, 0)
	for _, v := range mappingLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleDetailsEnumStringValues Enumerates the set of values in String for LifecycleDetailsEnum
func GetLifecycleDetailsEnumStringValues() []string {
	return []string{
		"STANDBY",
		"FAILOVER",
		"DOWN",
		"PRIMARY",
	}
}

// GetMappingLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleDetailsEnum(val string) (LifecycleDetailsEnum, bool) {
	enum, ok := mappingLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
