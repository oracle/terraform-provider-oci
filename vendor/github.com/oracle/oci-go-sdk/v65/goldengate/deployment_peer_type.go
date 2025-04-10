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

// DeploymentPeerTypeEnum Enum with underlying type: string
type DeploymentPeerTypeEnum string

// Set of constants representing the allowable values for DeploymentPeerTypeEnum
const (
	DeploymentPeerTypeLocal  DeploymentPeerTypeEnum = "LOCAL"
	DeploymentPeerTypeRemote DeploymentPeerTypeEnum = "REMOTE"
)

var mappingDeploymentPeerTypeEnum = map[string]DeploymentPeerTypeEnum{
	"LOCAL":  DeploymentPeerTypeLocal,
	"REMOTE": DeploymentPeerTypeRemote,
}

var mappingDeploymentPeerTypeEnumLowerCase = map[string]DeploymentPeerTypeEnum{
	"local":  DeploymentPeerTypeLocal,
	"remote": DeploymentPeerTypeRemote,
}

// GetDeploymentPeerTypeEnumValues Enumerates the set of values for DeploymentPeerTypeEnum
func GetDeploymentPeerTypeEnumValues() []DeploymentPeerTypeEnum {
	values := make([]DeploymentPeerTypeEnum, 0)
	for _, v := range mappingDeploymentPeerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentPeerTypeEnumStringValues Enumerates the set of values in String for DeploymentPeerTypeEnum
func GetDeploymentPeerTypeEnumStringValues() []string {
	return []string{
		"LOCAL",
		"REMOTE",
	}
}

// GetMappingDeploymentPeerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentPeerTypeEnum(val string) (DeploymentPeerTypeEnum, bool) {
	enum, ok := mappingDeploymentPeerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
