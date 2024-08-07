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

// DeploymentCategoryEnum Enum with underlying type: string
type DeploymentCategoryEnum string

// Set of constants representing the allowable values for DeploymentCategoryEnum
const (
	DeploymentCategoryDataReplication DeploymentCategoryEnum = "DATA_REPLICATION"
	DeploymentCategoryStreamAnalytics DeploymentCategoryEnum = "STREAM_ANALYTICS"
	DeploymentCategoryDataTransforms  DeploymentCategoryEnum = "DATA_TRANSFORMS"
)

var mappingDeploymentCategoryEnum = map[string]DeploymentCategoryEnum{
	"DATA_REPLICATION": DeploymentCategoryDataReplication,
	"STREAM_ANALYTICS": DeploymentCategoryStreamAnalytics,
	"DATA_TRANSFORMS":  DeploymentCategoryDataTransforms,
}

var mappingDeploymentCategoryEnumLowerCase = map[string]DeploymentCategoryEnum{
	"data_replication": DeploymentCategoryDataReplication,
	"stream_analytics": DeploymentCategoryStreamAnalytics,
	"data_transforms":  DeploymentCategoryDataTransforms,
}

// GetDeploymentCategoryEnumValues Enumerates the set of values for DeploymentCategoryEnum
func GetDeploymentCategoryEnumValues() []DeploymentCategoryEnum {
	values := make([]DeploymentCategoryEnum, 0)
	for _, v := range mappingDeploymentCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentCategoryEnumStringValues Enumerates the set of values in String for DeploymentCategoryEnum
func GetDeploymentCategoryEnumStringValues() []string {
	return []string{
		"DATA_REPLICATION",
		"STREAM_ANALYTICS",
		"DATA_TRANSFORMS",
	}
}

// GetMappingDeploymentCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentCategoryEnum(val string) (DeploymentCategoryEnum, bool) {
	enum, ok := mappingDeploymentCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
