// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"strings"
)

// LifeCycleStateEnum Enum with underlying type: string
type LifeCycleStateEnum string

// Set of constants representing the allowable values for LifeCycleStateEnum
const (
	LifeCycleStateActive   LifeCycleStateEnum = "ACTIVE"
	LifeCycleStateInactive LifeCycleStateEnum = "INACTIVE"
	LifeCycleStateDeleted  LifeCycleStateEnum = "DELETED"
)

var mappingLifeCycleStateEnum = map[string]LifeCycleStateEnum{
	"ACTIVE":   LifeCycleStateActive,
	"INACTIVE": LifeCycleStateInactive,
	"DELETED":  LifeCycleStateDeleted,
}

var mappingLifeCycleStateEnumLowerCase = map[string]LifeCycleStateEnum{
	"active":   LifeCycleStateActive,
	"inactive": LifeCycleStateInactive,
	"deleted":  LifeCycleStateDeleted,
}

// GetLifeCycleStateEnumValues Enumerates the set of values for LifeCycleStateEnum
func GetLifeCycleStateEnumValues() []LifeCycleStateEnum {
	values := make([]LifeCycleStateEnum, 0)
	for _, v := range mappingLifeCycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifeCycleStateEnumStringValues Enumerates the set of values in String for LifeCycleStateEnum
func GetLifeCycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingLifeCycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifeCycleStateEnum(val string) (LifeCycleStateEnum, bool) {
	enum, ok := mappingLifeCycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
