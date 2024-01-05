// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// AssociationLifecycleStateEnum Enum with underlying type: string
type AssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AssociationLifecycleStateEnum
const (
	AssociationLifecycleStateCreating AssociationLifecycleStateEnum = "CREATING"
	AssociationLifecycleStateActive   AssociationLifecycleStateEnum = "ACTIVE"
	AssociationLifecycleStateUpdating AssociationLifecycleStateEnum = "UPDATING"
	AssociationLifecycleStateDeleting AssociationLifecycleStateEnum = "DELETING"
	AssociationLifecycleStateFailed   AssociationLifecycleStateEnum = "FAILED"
)

var mappingAssociationLifecycleStateEnum = map[string]AssociationLifecycleStateEnum{
	"CREATING": AssociationLifecycleStateCreating,
	"ACTIVE":   AssociationLifecycleStateActive,
	"UPDATING": AssociationLifecycleStateUpdating,
	"DELETING": AssociationLifecycleStateDeleting,
	"FAILED":   AssociationLifecycleStateFailed,
}

var mappingAssociationLifecycleStateEnumLowerCase = map[string]AssociationLifecycleStateEnum{
	"creating": AssociationLifecycleStateCreating,
	"active":   AssociationLifecycleStateActive,
	"updating": AssociationLifecycleStateUpdating,
	"deleting": AssociationLifecycleStateDeleting,
	"failed":   AssociationLifecycleStateFailed,
}

// GetAssociationLifecycleStateEnumValues Enumerates the set of values for AssociationLifecycleStateEnum
func GetAssociationLifecycleStateEnumValues() []AssociationLifecycleStateEnum {
	values := make([]AssociationLifecycleStateEnum, 0)
	for _, v := range mappingAssociationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociationLifecycleStateEnumStringValues Enumerates the set of values in String for AssociationLifecycleStateEnum
func GetAssociationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

// GetMappingAssociationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociationLifecycleStateEnum(val string) (AssociationLifecycleStateEnum, bool) {
	enum, ok := mappingAssociationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
