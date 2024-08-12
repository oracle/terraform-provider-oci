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

// ProtectionLevelEnum Enum with underlying type: string
type ProtectionLevelEnum string

// Set of constants representing the allowable values for ProtectionLevelEnum
const (
	ProtectionLevelReadOnly             ProtectionLevelEnum = "READ_ONLY"
	ProtectionLevelPullRequestMergeOnly ProtectionLevelEnum = "PULL_REQUEST_MERGE_ONLY"
)

var mappingProtectionLevelEnum = map[string]ProtectionLevelEnum{
	"READ_ONLY":               ProtectionLevelReadOnly,
	"PULL_REQUEST_MERGE_ONLY": ProtectionLevelPullRequestMergeOnly,
}

var mappingProtectionLevelEnumLowerCase = map[string]ProtectionLevelEnum{
	"read_only":               ProtectionLevelReadOnly,
	"pull_request_merge_only": ProtectionLevelPullRequestMergeOnly,
}

// GetProtectionLevelEnumValues Enumerates the set of values for ProtectionLevelEnum
func GetProtectionLevelEnumValues() []ProtectionLevelEnum {
	values := make([]ProtectionLevelEnum, 0)
	for _, v := range mappingProtectionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectionLevelEnumStringValues Enumerates the set of values in String for ProtectionLevelEnum
func GetProtectionLevelEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"PULL_REQUEST_MERGE_ONLY",
	}
}

// GetMappingProtectionLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectionLevelEnum(val string) (ProtectionLevelEnum, bool) {
	enum, ok := mappingProtectionLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
