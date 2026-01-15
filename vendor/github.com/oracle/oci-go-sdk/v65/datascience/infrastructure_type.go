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

// InfrastructureTypeEnum Enum with underlying type: string
type InfrastructureTypeEnum string

// Set of constants representing the allowable values for InfrastructureTypeEnum
const (
	InfrastructureTypeInstancePool InfrastructureTypeEnum = "INSTANCE_POOL"
)

var mappingInfrastructureTypeEnum = map[string]InfrastructureTypeEnum{
	"INSTANCE_POOL": InfrastructureTypeInstancePool,
}

var mappingInfrastructureTypeEnumLowerCase = map[string]InfrastructureTypeEnum{
	"instance_pool": InfrastructureTypeInstancePool,
}

// GetInfrastructureTypeEnumValues Enumerates the set of values for InfrastructureTypeEnum
func GetInfrastructureTypeEnumValues() []InfrastructureTypeEnum {
	values := make([]InfrastructureTypeEnum, 0)
	for _, v := range mappingInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureTypeEnumStringValues Enumerates the set of values in String for InfrastructureTypeEnum
func GetInfrastructureTypeEnumStringValues() []string {
	return []string{
		"INSTANCE_POOL",
	}
}

// GetMappingInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureTypeEnum(val string) (InfrastructureTypeEnum, bool) {
	enum, ok := mappingInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
