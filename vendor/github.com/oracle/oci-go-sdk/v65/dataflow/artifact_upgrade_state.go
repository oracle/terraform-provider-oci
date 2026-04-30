// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// ArtifactUpgradeStateEnum Enum with underlying type: string
type ArtifactUpgradeStateEnum string

// Set of constants representing the allowable values for ArtifactUpgradeStateEnum
const (
	ArtifactUpgradeStateUpgradable               ArtifactUpgradeStateEnum = "UPGRADABLE"
	ArtifactUpgradeStateUpgradableDefaultCluster ArtifactUpgradeStateEnum = "UPGRADABLE_DEFAULT_CLUSTER"
	ArtifactUpgradeStateUpgradableUserCluster    ArtifactUpgradeStateEnum = "UPGRADABLE_USER_CLUSTER"
	ArtifactUpgradeStateUpgrading                ArtifactUpgradeStateEnum = "UPGRADING"
	ArtifactUpgradeStateUpgraded                 ArtifactUpgradeStateEnum = "UPGRADED"
	ArtifactUpgradeStateFailed                   ArtifactUpgradeStateEnum = "FAILED"
	ArtifactUpgradeStateUnknown                  ArtifactUpgradeStateEnum = "UNKNOWN"
)

var mappingArtifactUpgradeStateEnum = map[string]ArtifactUpgradeStateEnum{
	"UPGRADABLE":                 ArtifactUpgradeStateUpgradable,
	"UPGRADABLE_DEFAULT_CLUSTER": ArtifactUpgradeStateUpgradableDefaultCluster,
	"UPGRADABLE_USER_CLUSTER":    ArtifactUpgradeStateUpgradableUserCluster,
	"UPGRADING":                  ArtifactUpgradeStateUpgrading,
	"UPGRADED":                   ArtifactUpgradeStateUpgraded,
	"FAILED":                     ArtifactUpgradeStateFailed,
	"UNKNOWN":                    ArtifactUpgradeStateUnknown,
}

var mappingArtifactUpgradeStateEnumLowerCase = map[string]ArtifactUpgradeStateEnum{
	"upgradable":                 ArtifactUpgradeStateUpgradable,
	"upgradable_default_cluster": ArtifactUpgradeStateUpgradableDefaultCluster,
	"upgradable_user_cluster":    ArtifactUpgradeStateUpgradableUserCluster,
	"upgrading":                  ArtifactUpgradeStateUpgrading,
	"upgraded":                   ArtifactUpgradeStateUpgraded,
	"failed":                     ArtifactUpgradeStateFailed,
	"unknown":                    ArtifactUpgradeStateUnknown,
}

// GetArtifactUpgradeStateEnumValues Enumerates the set of values for ArtifactUpgradeStateEnum
func GetArtifactUpgradeStateEnumValues() []ArtifactUpgradeStateEnum {
	values := make([]ArtifactUpgradeStateEnum, 0)
	for _, v := range mappingArtifactUpgradeStateEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactUpgradeStateEnumStringValues Enumerates the set of values in String for ArtifactUpgradeStateEnum
func GetArtifactUpgradeStateEnumStringValues() []string {
	return []string{
		"UPGRADABLE",
		"UPGRADABLE_DEFAULT_CLUSTER",
		"UPGRADABLE_USER_CLUSTER",
		"UPGRADING",
		"UPGRADED",
		"FAILED",
		"UNKNOWN",
	}
}

// GetMappingArtifactUpgradeStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactUpgradeStateEnum(val string) (ArtifactUpgradeStateEnum, bool) {
	enum, ok := mappingArtifactUpgradeStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
