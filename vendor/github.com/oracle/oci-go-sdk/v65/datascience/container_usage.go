// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerUsageEnum Enum with underlying type: string
type ContainerUsageEnum string

// Set of constants representing the allowable values for ContainerUsageEnum
const (
	ContainerUsageInference      ContainerUsageEnum = "INFERENCE"
	ContainerUsageFineTune       ContainerUsageEnum = "FINE_TUNE"
	ContainerUsageEvaluation     ContainerUsageEnum = "EVALUATION"
	ContainerUsageBatchInference ContainerUsageEnum = "BATCH_INFERENCE"
	ContainerUsageOther          ContainerUsageEnum = "OTHER"
)

var mappingContainerUsageEnum = map[string]ContainerUsageEnum{
	"INFERENCE":       ContainerUsageInference,
	"FINE_TUNE":       ContainerUsageFineTune,
	"EVALUATION":      ContainerUsageEvaluation,
	"BATCH_INFERENCE": ContainerUsageBatchInference,
	"OTHER":           ContainerUsageOther,
}

var mappingContainerUsageEnumLowerCase = map[string]ContainerUsageEnum{
	"inference":       ContainerUsageInference,
	"fine_tune":       ContainerUsageFineTune,
	"evaluation":      ContainerUsageEvaluation,
	"batch_inference": ContainerUsageBatchInference,
	"other":           ContainerUsageOther,
}

// GetContainerUsageEnumValues Enumerates the set of values for ContainerUsageEnum
func GetContainerUsageEnumValues() []ContainerUsageEnum {
	values := make([]ContainerUsageEnum, 0)
	for _, v := range mappingContainerUsageEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerUsageEnumStringValues Enumerates the set of values in String for ContainerUsageEnum
func GetContainerUsageEnumStringValues() []string {
	return []string{
		"INFERENCE",
		"FINE_TUNE",
		"EVALUATION",
		"BATCH_INFERENCE",
		"OTHER",
	}
}

// GetMappingContainerUsageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerUsageEnum(val string) (ContainerUsageEnum, bool) {
	enum, ok := mappingContainerUsageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
