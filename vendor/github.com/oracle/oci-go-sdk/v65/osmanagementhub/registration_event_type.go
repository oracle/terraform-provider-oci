// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// RegistrationEventTypeEnum Enum with underlying type: string
type RegistrationEventTypeEnum string

// Set of constants representing the allowable values for RegistrationEventTypeEnum
const (
	RegistrationEventTypeRegisterManagedInstance RegistrationEventTypeEnum = "REGISTER_MANAGED_INSTANCE"
	RegistrationEventTypeReposMatched            RegistrationEventTypeEnum = "REPOS_MATCHED"
)

var mappingRegistrationEventTypeEnum = map[string]RegistrationEventTypeEnum{
	"REGISTER_MANAGED_INSTANCE": RegistrationEventTypeRegisterManagedInstance,
	"REPOS_MATCHED":             RegistrationEventTypeReposMatched,
}

var mappingRegistrationEventTypeEnumLowerCase = map[string]RegistrationEventTypeEnum{
	"register_managed_instance": RegistrationEventTypeRegisterManagedInstance,
	"repos_matched":             RegistrationEventTypeReposMatched,
}

// GetRegistrationEventTypeEnumValues Enumerates the set of values for RegistrationEventTypeEnum
func GetRegistrationEventTypeEnumValues() []RegistrationEventTypeEnum {
	values := make([]RegistrationEventTypeEnum, 0)
	for _, v := range mappingRegistrationEventTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationEventTypeEnumStringValues Enumerates the set of values in String for RegistrationEventTypeEnum
func GetRegistrationEventTypeEnumStringValues() []string {
	return []string{
		"REGISTER_MANAGED_INSTANCE",
		"REPOS_MATCHED",
	}
}

// GetMappingRegistrationEventTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationEventTypeEnum(val string) (RegistrationEventTypeEnum, bool) {
	enum, ok := mappingRegistrationEventTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
