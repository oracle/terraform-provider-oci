// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// MaintenanceActionTypeEnum Enum with underlying type: string
type MaintenanceActionTypeEnum string

// Set of constants representing the allowable values for MaintenanceActionTypeEnum
const (
	MaintenanceActionTypeUpgrade MaintenanceActionTypeEnum = "UPGRADE"
)

var mappingMaintenanceActionTypeEnum = map[string]MaintenanceActionTypeEnum{
	"UPGRADE": MaintenanceActionTypeUpgrade,
}

var mappingMaintenanceActionTypeEnumLowerCase = map[string]MaintenanceActionTypeEnum{
	"upgrade": MaintenanceActionTypeUpgrade,
}

// GetMaintenanceActionTypeEnumValues Enumerates the set of values for MaintenanceActionTypeEnum
func GetMaintenanceActionTypeEnumValues() []MaintenanceActionTypeEnum {
	values := make([]MaintenanceActionTypeEnum, 0)
	for _, v := range mappingMaintenanceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceActionTypeEnumStringValues Enumerates the set of values in String for MaintenanceActionTypeEnum
func GetMaintenanceActionTypeEnumStringValues() []string {
	return []string{
		"UPGRADE",
	}
}

// GetMappingMaintenanceActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceActionTypeEnum(val string) (MaintenanceActionTypeEnum, bool) {
	enum, ok := mappingMaintenanceActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
