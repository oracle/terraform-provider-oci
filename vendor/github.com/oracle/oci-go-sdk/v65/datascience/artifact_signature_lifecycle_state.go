// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ArtifactSignatureLifecycleStateEnum Enum with underlying type: string
type ArtifactSignatureLifecycleStateEnum string

// Set of constants representing the allowable values for ArtifactSignatureLifecycleStateEnum
const (
	ArtifactSignatureLifecycleStateActive   ArtifactSignatureLifecycleStateEnum = "ACTIVE"
	ArtifactSignatureLifecycleStateCreating ArtifactSignatureLifecycleStateEnum = "CREATING"
	ArtifactSignatureLifecycleStateDeleted  ArtifactSignatureLifecycleStateEnum = "DELETED"
	ArtifactSignatureLifecycleStateFailed   ArtifactSignatureLifecycleStateEnum = "FAILED"
)

var mappingArtifactSignatureLifecycleStateEnum = map[string]ArtifactSignatureLifecycleStateEnum{
	"ACTIVE":   ArtifactSignatureLifecycleStateActive,
	"CREATING": ArtifactSignatureLifecycleStateCreating,
	"DELETED":  ArtifactSignatureLifecycleStateDeleted,
	"FAILED":   ArtifactSignatureLifecycleStateFailed,
}

var mappingArtifactSignatureLifecycleStateEnumLowerCase = map[string]ArtifactSignatureLifecycleStateEnum{
	"active":   ArtifactSignatureLifecycleStateActive,
	"creating": ArtifactSignatureLifecycleStateCreating,
	"deleted":  ArtifactSignatureLifecycleStateDeleted,
	"failed":   ArtifactSignatureLifecycleStateFailed,
}

// GetArtifactSignatureLifecycleStateEnumValues Enumerates the set of values for ArtifactSignatureLifecycleStateEnum
func GetArtifactSignatureLifecycleStateEnumValues() []ArtifactSignatureLifecycleStateEnum {
	values := make([]ArtifactSignatureLifecycleStateEnum, 0)
	for _, v := range mappingArtifactSignatureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactSignatureLifecycleStateEnumStringValues Enumerates the set of values in String for ArtifactSignatureLifecycleStateEnum
func GetArtifactSignatureLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingArtifactSignatureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactSignatureLifecycleStateEnum(val string) (ArtifactSignatureLifecycleStateEnum, bool) {
	enum, ok := mappingArtifactSignatureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
