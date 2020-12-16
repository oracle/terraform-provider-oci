// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating LifecycleStateEnum = "CREATING"
	LifecycleStateUpdating LifecycleStateEnum = "UPDATING"
	LifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	LifecycleStateDeleting LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted  LifecycleStateEnum = "DELETED"
	LifecycleStateFailed   LifecycleStateEnum = "FAILED"
	LifecycleStateNa       LifecycleStateEnum = "NA"
)

var mappingLifecycleState = map[string]LifecycleStateEnum{
	"CREATING": LifecycleStateCreating,
	"UPDATING": LifecycleStateUpdating,
	"ACTIVE":   LifecycleStateActive,
	"DELETING": LifecycleStateDeleting,
	"DELETED":  LifecycleStateDeleted,
	"FAILED":   LifecycleStateFailed,
	"NA":       LifecycleStateNa,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleState {
		values = append(values, v)
	}
	return values
}
