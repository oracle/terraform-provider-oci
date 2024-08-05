// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"strings"
)

// MergeStrategyEnum Enum with underlying type: string
type MergeStrategyEnum string

// Set of constants representing the allowable values for MergeStrategyEnum
const (
	MergeStrategyMergeCommit           MergeStrategyEnum = "MERGE_COMMIT"
	MergeStrategyFastForward           MergeStrategyEnum = "FAST_FORWARD"
	MergeStrategyFastForwardOnly       MergeStrategyEnum = "FAST_FORWARD_ONLY"
	MergeStrategyRebaseAndMerge        MergeStrategyEnum = "REBASE_AND_MERGE"
	MergeStrategyRebaseAndFastForward  MergeStrategyEnum = "REBASE_AND_FAST_FORWARD"
	MergeStrategySquash                MergeStrategyEnum = "SQUASH"
	MergeStrategySquashFastForwardOnly MergeStrategyEnum = "SQUASH_FAST_FORWARD_ONLY"
)

var mappingMergeStrategyEnum = map[string]MergeStrategyEnum{
	"MERGE_COMMIT":             MergeStrategyMergeCommit,
	"FAST_FORWARD":             MergeStrategyFastForward,
	"FAST_FORWARD_ONLY":        MergeStrategyFastForwardOnly,
	"REBASE_AND_MERGE":         MergeStrategyRebaseAndMerge,
	"REBASE_AND_FAST_FORWARD":  MergeStrategyRebaseAndFastForward,
	"SQUASH":                   MergeStrategySquash,
	"SQUASH_FAST_FORWARD_ONLY": MergeStrategySquashFastForwardOnly,
}

var mappingMergeStrategyEnumLowerCase = map[string]MergeStrategyEnum{
	"merge_commit":             MergeStrategyMergeCommit,
	"fast_forward":             MergeStrategyFastForward,
	"fast_forward_only":        MergeStrategyFastForwardOnly,
	"rebase_and_merge":         MergeStrategyRebaseAndMerge,
	"rebase_and_fast_forward":  MergeStrategyRebaseAndFastForward,
	"squash":                   MergeStrategySquash,
	"squash_fast_forward_only": MergeStrategySquashFastForwardOnly,
}

// GetMergeStrategyEnumValues Enumerates the set of values for MergeStrategyEnum
func GetMergeStrategyEnumValues() []MergeStrategyEnum {
	values := make([]MergeStrategyEnum, 0)
	for _, v := range mappingMergeStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetMergeStrategyEnumStringValues Enumerates the set of values in String for MergeStrategyEnum
func GetMergeStrategyEnumStringValues() []string {
	return []string{
		"MERGE_COMMIT",
		"FAST_FORWARD",
		"FAST_FORWARD_ONLY",
		"REBASE_AND_MERGE",
		"REBASE_AND_FAST_FORWARD",
		"SQUASH",
		"SQUASH_FAST_FORWARD_ONLY",
	}
}

// GetMappingMergeStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMergeStrategyEnum(val string) (MergeStrategyEnum, bool) {
	enum, ok := mappingMergeStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
