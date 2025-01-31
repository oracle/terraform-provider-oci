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

// ContainerTargetWorkloadEnum Enum with underlying type: string
type ContainerTargetWorkloadEnum string

// Set of constants representing the allowable values for ContainerTargetWorkloadEnum
const (
	ContainerTargetWorkloadModelDeployment ContainerTargetWorkloadEnum = "MODEL_DEPLOYMENT"
	ContainerTargetWorkloadJobRun          ContainerTargetWorkloadEnum = "JOB_RUN"
)

var mappingContainerTargetWorkloadEnum = map[string]ContainerTargetWorkloadEnum{
	"MODEL_DEPLOYMENT": ContainerTargetWorkloadModelDeployment,
	"JOB_RUN":          ContainerTargetWorkloadJobRun,
}

var mappingContainerTargetWorkloadEnumLowerCase = map[string]ContainerTargetWorkloadEnum{
	"model_deployment": ContainerTargetWorkloadModelDeployment,
	"job_run":          ContainerTargetWorkloadJobRun,
}

// GetContainerTargetWorkloadEnumValues Enumerates the set of values for ContainerTargetWorkloadEnum
func GetContainerTargetWorkloadEnumValues() []ContainerTargetWorkloadEnum {
	values := make([]ContainerTargetWorkloadEnum, 0)
	for _, v := range mappingContainerTargetWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerTargetWorkloadEnumStringValues Enumerates the set of values in String for ContainerTargetWorkloadEnum
func GetContainerTargetWorkloadEnumStringValues() []string {
	return []string{
		"MODEL_DEPLOYMENT",
		"JOB_RUN",
	}
}

// GetMappingContainerTargetWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerTargetWorkloadEnum(val string) (ContainerTargetWorkloadEnum, bool) {
	enum, ok := mappingContainerTargetWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
