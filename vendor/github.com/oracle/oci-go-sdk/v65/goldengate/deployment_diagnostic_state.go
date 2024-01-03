// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeploymentDiagnosticStateEnum Enum with underlying type: string
type DeploymentDiagnosticStateEnum string

// Set of constants representing the allowable values for DeploymentDiagnosticStateEnum
const (
	DeploymentDiagnosticStateInProgress DeploymentDiagnosticStateEnum = "IN_PROGRESS"
	DeploymentDiagnosticStateSucceeded  DeploymentDiagnosticStateEnum = "SUCCEEDED"
	DeploymentDiagnosticStateFailed     DeploymentDiagnosticStateEnum = "FAILED"
)

var mappingDeploymentDiagnosticStateEnum = map[string]DeploymentDiagnosticStateEnum{
	"IN_PROGRESS": DeploymentDiagnosticStateInProgress,
	"SUCCEEDED":   DeploymentDiagnosticStateSucceeded,
	"FAILED":      DeploymentDiagnosticStateFailed,
}

var mappingDeploymentDiagnosticStateEnumLowerCase = map[string]DeploymentDiagnosticStateEnum{
	"in_progress": DeploymentDiagnosticStateInProgress,
	"succeeded":   DeploymentDiagnosticStateSucceeded,
	"failed":      DeploymentDiagnosticStateFailed,
}

// GetDeploymentDiagnosticStateEnumValues Enumerates the set of values for DeploymentDiagnosticStateEnum
func GetDeploymentDiagnosticStateEnumValues() []DeploymentDiagnosticStateEnum {
	values := make([]DeploymentDiagnosticStateEnum, 0)
	for _, v := range mappingDeploymentDiagnosticStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentDiagnosticStateEnumStringValues Enumerates the set of values in String for DeploymentDiagnosticStateEnum
func GetDeploymentDiagnosticStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingDeploymentDiagnosticStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentDiagnosticStateEnum(val string) (DeploymentDiagnosticStateEnum, bool) {
	enum, ok := mappingDeploymentDiagnosticStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
