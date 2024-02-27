// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// TrainSkillParameterTypeEnum Enum with underlying type: string
type TrainSkillParameterTypeEnum string

// Set of constants representing the allowable values for TrainSkillParameterTypeEnum
const (
	TrainSkillParameterTypeQueryEntity TrainSkillParameterTypeEnum = "QUERY_ENTITY"
)

var mappingTrainSkillParameterTypeEnum = map[string]TrainSkillParameterTypeEnum{
	"QUERY_ENTITY": TrainSkillParameterTypeQueryEntity,
}

var mappingTrainSkillParameterTypeEnumLowerCase = map[string]TrainSkillParameterTypeEnum{
	"query_entity": TrainSkillParameterTypeQueryEntity,
}

// GetTrainSkillParameterTypeEnumValues Enumerates the set of values for TrainSkillParameterTypeEnum
func GetTrainSkillParameterTypeEnumValues() []TrainSkillParameterTypeEnum {
	values := make([]TrainSkillParameterTypeEnum, 0)
	for _, v := range mappingTrainSkillParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrainSkillParameterTypeEnumStringValues Enumerates the set of values in String for TrainSkillParameterTypeEnum
func GetTrainSkillParameterTypeEnumStringValues() []string {
	return []string{
		"QUERY_ENTITY",
	}
}

// GetMappingTrainSkillParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrainSkillParameterTypeEnum(val string) (TrainSkillParameterTypeEnum, bool) {
	enum, ok := mappingTrainSkillParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
