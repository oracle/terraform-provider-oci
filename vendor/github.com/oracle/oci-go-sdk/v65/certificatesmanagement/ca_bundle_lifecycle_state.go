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

// CaBundleLifecycleStateEnum Enum with underlying type: string
type CaBundleLifecycleStateEnum string

// Set of constants representing the allowable values for CaBundleLifecycleStateEnum
const (
	CaBundleLifecycleStateCreating CaBundleLifecycleStateEnum = "CREATING"
	CaBundleLifecycleStateActive   CaBundleLifecycleStateEnum = "ACTIVE"
	CaBundleLifecycleStateUpdating CaBundleLifecycleStateEnum = "UPDATING"
	CaBundleLifecycleStateDeleting CaBundleLifecycleStateEnum = "DELETING"
	CaBundleLifecycleStateDeleted  CaBundleLifecycleStateEnum = "DELETED"
	CaBundleLifecycleStateFailed   CaBundleLifecycleStateEnum = "FAILED"
)

var mappingCaBundleLifecycleStateEnum = map[string]CaBundleLifecycleStateEnum{
	"CREATING": CaBundleLifecycleStateCreating,
	"ACTIVE":   CaBundleLifecycleStateActive,
	"UPDATING": CaBundleLifecycleStateUpdating,
	"DELETING": CaBundleLifecycleStateDeleting,
	"DELETED":  CaBundleLifecycleStateDeleted,
	"FAILED":   CaBundleLifecycleStateFailed,
}

var mappingCaBundleLifecycleStateEnumLowerCase = map[string]CaBundleLifecycleStateEnum{
	"creating": CaBundleLifecycleStateCreating,
	"active":   CaBundleLifecycleStateActive,
	"updating": CaBundleLifecycleStateUpdating,
	"deleting": CaBundleLifecycleStateDeleting,
	"deleted":  CaBundleLifecycleStateDeleted,
	"failed":   CaBundleLifecycleStateFailed,
}

// GetCaBundleLifecycleStateEnumValues Enumerates the set of values for CaBundleLifecycleStateEnum
func GetCaBundleLifecycleStateEnumValues() []CaBundleLifecycleStateEnum {
	values := make([]CaBundleLifecycleStateEnum, 0)
	for _, v := range mappingCaBundleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCaBundleLifecycleStateEnumStringValues Enumerates the set of values in String for CaBundleLifecycleStateEnum
func GetCaBundleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCaBundleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCaBundleLifecycleStateEnum(val string) (CaBundleLifecycleStateEnum, bool) {
	enum, ok := mappingCaBundleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
