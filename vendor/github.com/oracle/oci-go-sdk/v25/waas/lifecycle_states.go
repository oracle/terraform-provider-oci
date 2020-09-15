// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

// LifecycleStatesEnum Enum with underlying type: string
type LifecycleStatesEnum string

// Set of constants representing the allowable values for LifecycleStatesEnum
const (
	LifecycleStatesCreating LifecycleStatesEnum = "CREATING"
	LifecycleStatesActive   LifecycleStatesEnum = "ACTIVE"
	LifecycleStatesFailed   LifecycleStatesEnum = "FAILED"
	LifecycleStatesUpdating LifecycleStatesEnum = "UPDATING"
	LifecycleStatesDeleting LifecycleStatesEnum = "DELETING"
	LifecycleStatesDeleted  LifecycleStatesEnum = "DELETED"
)

var mappingLifecycleStates = map[string]LifecycleStatesEnum{
	"CREATING": LifecycleStatesCreating,
	"ACTIVE":   LifecycleStatesActive,
	"FAILED":   LifecycleStatesFailed,
	"UPDATING": LifecycleStatesUpdating,
	"DELETING": LifecycleStatesDeleting,
	"DELETED":  LifecycleStatesDeleted,
}

// GetLifecycleStatesEnumValues Enumerates the set of values for LifecycleStatesEnum
func GetLifecycleStatesEnumValues() []LifecycleStatesEnum {
	values := make([]LifecycleStatesEnum, 0)
	for _, v := range mappingLifecycleStates {
		values = append(values, v)
	}
	return values
}
