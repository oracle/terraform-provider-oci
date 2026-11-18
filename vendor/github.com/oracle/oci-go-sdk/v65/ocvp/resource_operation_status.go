// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// ResourceOperationStatusEnum Enum with underlying type: string
type ResourceOperationStatusEnum string

// Set of constants representing the allowable values for ResourceOperationStatusEnum
const (
	ResourceOperationStatusInProgress     ResourceOperationStatusEnum = "IN_PROGRESS"
	ResourceOperationStatusFailed         ResourceOperationStatusEnum = "FAILED"
	ResourceOperationStatusPartialSucceed ResourceOperationStatusEnum = "PARTIAL_SUCCEED"
	ResourceOperationStatusSucceeded      ResourceOperationStatusEnum = "SUCCEEDED"
)

var mappingResourceOperationStatusEnum = map[string]ResourceOperationStatusEnum{
	"IN_PROGRESS":     ResourceOperationStatusInProgress,
	"FAILED":          ResourceOperationStatusFailed,
	"PARTIAL_SUCCEED": ResourceOperationStatusPartialSucceed,
	"SUCCEEDED":       ResourceOperationStatusSucceeded,
}

var mappingResourceOperationStatusEnumLowerCase = map[string]ResourceOperationStatusEnum{
	"in_progress":     ResourceOperationStatusInProgress,
	"failed":          ResourceOperationStatusFailed,
	"partial_succeed": ResourceOperationStatusPartialSucceed,
	"succeeded":       ResourceOperationStatusSucceeded,
}

// GetResourceOperationStatusEnumValues Enumerates the set of values for ResourceOperationStatusEnum
func GetResourceOperationStatusEnumValues() []ResourceOperationStatusEnum {
	values := make([]ResourceOperationStatusEnum, 0)
	for _, v := range mappingResourceOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceOperationStatusEnumStringValues Enumerates the set of values in String for ResourceOperationStatusEnum
func GetResourceOperationStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"FAILED",
		"PARTIAL_SUCCEED",
		"SUCCEEDED",
	}
}

// GetMappingResourceOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceOperationStatusEnum(val string) (ResourceOperationStatusEnum, bool) {
	enum, ok := mappingResourceOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
