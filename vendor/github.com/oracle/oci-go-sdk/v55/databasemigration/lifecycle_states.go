// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// LifecycleStatesEnum Enum with underlying type: string
type LifecycleStatesEnum string

// Set of constants representing the allowable values for LifecycleStatesEnum
const (
	LifecycleStatesCreating LifecycleStatesEnum = "CREATING"
	LifecycleStatesUpdating LifecycleStatesEnum = "UPDATING"
	LifecycleStatesActive   LifecycleStatesEnum = "ACTIVE"
	LifecycleStatesInactive LifecycleStatesEnum = "INACTIVE"
	LifecycleStatesDeleting LifecycleStatesEnum = "DELETING"
	LifecycleStatesDeleted  LifecycleStatesEnum = "DELETED"
	LifecycleStatesFailed   LifecycleStatesEnum = "FAILED"
)

var mappingLifecycleStates = map[string]LifecycleStatesEnum{
	"CREATING": LifecycleStatesCreating,
	"UPDATING": LifecycleStatesUpdating,
	"ACTIVE":   LifecycleStatesActive,
	"INACTIVE": LifecycleStatesInactive,
	"DELETING": LifecycleStatesDeleting,
	"DELETED":  LifecycleStatesDeleted,
	"FAILED":   LifecycleStatesFailed,
}

// GetLifecycleStatesEnumValues Enumerates the set of values for LifecycleStatesEnum
func GetLifecycleStatesEnumValues() []LifecycleStatesEnum {
	values := make([]LifecycleStatesEnum, 0)
	for _, v := range mappingLifecycleStates {
		values = append(values, v)
	}
	return values
}
