// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// WorkRequestActionTypeEnum Enum with underlying type: string
type WorkRequestActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestActionTypeEnum
const (
	WorkRequestActionTypeAccepted       WorkRequestActionTypeEnum = "ACCEPTED"
	WorkRequestActionTypeInProgress     WorkRequestActionTypeEnum = "IN_PROGRESS"
	WorkRequestActionTypeWaiting        WorkRequestActionTypeEnum = "WAITING"
	WorkRequestActionTypeNeedsAttention WorkRequestActionTypeEnum = "NEEDS_ATTENTION"
	WorkRequestActionTypeFailed         WorkRequestActionTypeEnum = "FAILED"
	WorkRequestActionTypeSucceeded      WorkRequestActionTypeEnum = "SUCCEEDED"
	WorkRequestActionTypeCanceled       WorkRequestActionTypeEnum = "CANCELED"
)

var mappingWorkRequestActionTypeEnum = map[string]WorkRequestActionTypeEnum{
	"ACCEPTED":        WorkRequestActionTypeAccepted,
	"IN_PROGRESS":     WorkRequestActionTypeInProgress,
	"WAITING":         WorkRequestActionTypeWaiting,
	"NEEDS_ATTENTION": WorkRequestActionTypeNeedsAttention,
	"FAILED":          WorkRequestActionTypeFailed,
	"SUCCEEDED":       WorkRequestActionTypeSucceeded,
	"CANCELED":        WorkRequestActionTypeCanceled,
}

var mappingWorkRequestActionTypeEnumLowerCase = map[string]WorkRequestActionTypeEnum{
	"accepted":        WorkRequestActionTypeAccepted,
	"in_progress":     WorkRequestActionTypeInProgress,
	"waiting":         WorkRequestActionTypeWaiting,
	"needs_attention": WorkRequestActionTypeNeedsAttention,
	"failed":          WorkRequestActionTypeFailed,
	"succeeded":       WorkRequestActionTypeSucceeded,
	"canceled":        WorkRequestActionTypeCanceled,
}

// GetWorkRequestActionTypeEnumValues Enumerates the set of values for WorkRequestActionTypeEnum
func GetWorkRequestActionTypeEnumValues() []WorkRequestActionTypeEnum {
	values := make([]WorkRequestActionTypeEnum, 0)
	for _, v := range mappingWorkRequestActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestActionTypeEnum
func GetWorkRequestActionTypeEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
	}
}

// GetMappingWorkRequestActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestActionTypeEnum(val string) (WorkRequestActionTypeEnum, bool) {
	enum, ok := mappingWorkRequestActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
