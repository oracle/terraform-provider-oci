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

// ModelDeployInfrastructureTypeEnum Enum with underlying type: string
type ModelDeployInfrastructureTypeEnum string

// Set of constants representing the allowable values for ModelDeployInfrastructureTypeEnum
const (
	ModelDeployInfrastructureTypeManagedComputeCluster ModelDeployInfrastructureTypeEnum = "MANAGED_COMPUTE_CLUSTER"
)

var mappingModelDeployInfrastructureTypeEnum = map[string]ModelDeployInfrastructureTypeEnum{
	"MANAGED_COMPUTE_CLUSTER": ModelDeployInfrastructureTypeManagedComputeCluster,
}

var mappingModelDeployInfrastructureTypeEnumLowerCase = map[string]ModelDeployInfrastructureTypeEnum{
	"managed_compute_cluster": ModelDeployInfrastructureTypeManagedComputeCluster,
}

// GetModelDeployInfrastructureTypeEnumValues Enumerates the set of values for ModelDeployInfrastructureTypeEnum
func GetModelDeployInfrastructureTypeEnumValues() []ModelDeployInfrastructureTypeEnum {
	values := make([]ModelDeployInfrastructureTypeEnum, 0)
	for _, v := range mappingModelDeployInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDeployInfrastructureTypeEnumStringValues Enumerates the set of values in String for ModelDeployInfrastructureTypeEnum
func GetModelDeployInfrastructureTypeEnumStringValues() []string {
	return []string{
		"MANAGED_COMPUTE_CLUSTER",
	}
}

// GetMappingModelDeployInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDeployInfrastructureTypeEnum(val string) (ModelDeployInfrastructureTypeEnum, bool) {
	enum, ok := mappingModelDeployInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
