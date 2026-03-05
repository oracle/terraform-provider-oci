// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PrecheckStatusEnum Enum with underlying type: string
type PrecheckStatusEnum string

// Set of constants representing the allowable values for PrecheckStatusEnum
const (
	PrecheckStatusSucceeded      PrecheckStatusEnum = "SUCCEEDED"
	PrecheckStatusInProgress     PrecheckStatusEnum = "IN_PROGRESS"
	PrecheckStatusFailed         PrecheckStatusEnum = "FAILED"
	PrecheckStatusNeedsAttention PrecheckStatusEnum = "NEEDS_ATTENTION"
)

var mappingPrecheckStatusEnum = map[string]PrecheckStatusEnum{
	"SUCCEEDED":       PrecheckStatusSucceeded,
	"IN_PROGRESS":     PrecheckStatusInProgress,
	"FAILED":          PrecheckStatusFailed,
	"NEEDS_ATTENTION": PrecheckStatusNeedsAttention,
}

var mappingPrecheckStatusEnumLowerCase = map[string]PrecheckStatusEnum{
	"succeeded":       PrecheckStatusSucceeded,
	"in_progress":     PrecheckStatusInProgress,
	"failed":          PrecheckStatusFailed,
	"needs_attention": PrecheckStatusNeedsAttention,
}

// GetPrecheckStatusEnumValues Enumerates the set of values for PrecheckStatusEnum
func GetPrecheckStatusEnumValues() []PrecheckStatusEnum {
	values := make([]PrecheckStatusEnum, 0)
	for _, v := range mappingPrecheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPrecheckStatusEnumStringValues Enumerates the set of values in String for PrecheckStatusEnum
func GetPrecheckStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"IN_PROGRESS",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingPrecheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrecheckStatusEnum(val string) (PrecheckStatusEnum, bool) {
	enum, ok := mappingPrecheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
