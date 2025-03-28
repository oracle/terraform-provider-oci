// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DeploymentUpgradeLifecycleStateEnum Enum with underlying type: string
type DeploymentUpgradeLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentUpgradeLifecycleStateEnum
const (
	DeploymentUpgradeLifecycleStateWaiting        DeploymentUpgradeLifecycleStateEnum = "WAITING"
	DeploymentUpgradeLifecycleStateInProgress     DeploymentUpgradeLifecycleStateEnum = "IN_PROGRESS"
	DeploymentUpgradeLifecycleStateFailed         DeploymentUpgradeLifecycleStateEnum = "FAILED"
	DeploymentUpgradeLifecycleStateSucceeded      DeploymentUpgradeLifecycleStateEnum = "SUCCEEDED"
	DeploymentUpgradeLifecycleStateCanceling      DeploymentUpgradeLifecycleStateEnum = "CANCELING"
	DeploymentUpgradeLifecycleStateCanceled       DeploymentUpgradeLifecycleStateEnum = "CANCELED"
	DeploymentUpgradeLifecycleStateNeedsAttention DeploymentUpgradeLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDeploymentUpgradeLifecycleStateEnum = map[string]DeploymentUpgradeLifecycleStateEnum{
	"WAITING":         DeploymentUpgradeLifecycleStateWaiting,
	"IN_PROGRESS":     DeploymentUpgradeLifecycleStateInProgress,
	"FAILED":          DeploymentUpgradeLifecycleStateFailed,
	"SUCCEEDED":       DeploymentUpgradeLifecycleStateSucceeded,
	"CANCELING":       DeploymentUpgradeLifecycleStateCanceling,
	"CANCELED":        DeploymentUpgradeLifecycleStateCanceled,
	"NEEDS_ATTENTION": DeploymentUpgradeLifecycleStateNeedsAttention,
}

var mappingDeploymentUpgradeLifecycleStateEnumLowerCase = map[string]DeploymentUpgradeLifecycleStateEnum{
	"waiting":         DeploymentUpgradeLifecycleStateWaiting,
	"in_progress":     DeploymentUpgradeLifecycleStateInProgress,
	"failed":          DeploymentUpgradeLifecycleStateFailed,
	"succeeded":       DeploymentUpgradeLifecycleStateSucceeded,
	"canceling":       DeploymentUpgradeLifecycleStateCanceling,
	"canceled":        DeploymentUpgradeLifecycleStateCanceled,
	"needs_attention": DeploymentUpgradeLifecycleStateNeedsAttention,
}

// GetDeploymentUpgradeLifecycleStateEnumValues Enumerates the set of values for DeploymentUpgradeLifecycleStateEnum
func GetDeploymentUpgradeLifecycleStateEnumValues() []DeploymentUpgradeLifecycleStateEnum {
	values := make([]DeploymentUpgradeLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentUpgradeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentUpgradeLifecycleStateEnumStringValues Enumerates the set of values in String for DeploymentUpgradeLifecycleStateEnum
func GetDeploymentUpgradeLifecycleStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDeploymentUpgradeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentUpgradeLifecycleStateEnum(val string) (DeploymentUpgradeLifecycleStateEnum, bool) {
	enum, ok := mappingDeploymentUpgradeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
