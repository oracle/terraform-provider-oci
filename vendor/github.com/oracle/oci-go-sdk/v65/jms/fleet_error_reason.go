// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// FleetErrorReasonEnum Enum with underlying type: string
type FleetErrorReasonEnum string

// Set of constants representing the allowable values for FleetErrorReasonEnum
const (
	FleetErrorReasonNoManagedInstances FleetErrorReasonEnum = "NO_MANAGED_INSTANCES"
	FleetErrorReasonInventoryLog       FleetErrorReasonEnum = "INVENTORY_LOG"
)

var mappingFleetErrorReasonEnum = map[string]FleetErrorReasonEnum{
	"NO_MANAGED_INSTANCES": FleetErrorReasonNoManagedInstances,
	"INVENTORY_LOG":        FleetErrorReasonInventoryLog,
}

var mappingFleetErrorReasonEnumLowerCase = map[string]FleetErrorReasonEnum{
	"no_managed_instances": FleetErrorReasonNoManagedInstances,
	"inventory_log":        FleetErrorReasonInventoryLog,
}

// GetFleetErrorReasonEnumValues Enumerates the set of values for FleetErrorReasonEnum
func GetFleetErrorReasonEnumValues() []FleetErrorReasonEnum {
	values := make([]FleetErrorReasonEnum, 0)
	for _, v := range mappingFleetErrorReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetErrorReasonEnumStringValues Enumerates the set of values in String for FleetErrorReasonEnum
func GetFleetErrorReasonEnumStringValues() []string {
	return []string{
		"NO_MANAGED_INSTANCES",
		"INVENTORY_LOG",
	}
}

// GetMappingFleetErrorReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetErrorReasonEnum(val string) (FleetErrorReasonEnum, bool) {
	enum, ok := mappingFleetErrorReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
