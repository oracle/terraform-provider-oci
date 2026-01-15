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

// ModelCategoryEnum Enum with underlying type: string
type ModelCategoryEnum string

// Set of constants representing the allowable values for ModelCategoryEnum
const (
	ModelCategoryUser    ModelCategoryEnum = "USER"
	ModelCategoryService ModelCategoryEnum = "SERVICE"
)

var mappingModelCategoryEnum = map[string]ModelCategoryEnum{
	"USER":    ModelCategoryUser,
	"SERVICE": ModelCategoryService,
}

var mappingModelCategoryEnumLowerCase = map[string]ModelCategoryEnum{
	"user":    ModelCategoryUser,
	"service": ModelCategoryService,
}

// GetModelCategoryEnumValues Enumerates the set of values for ModelCategoryEnum
func GetModelCategoryEnumValues() []ModelCategoryEnum {
	values := make([]ModelCategoryEnum, 0)
	for _, v := range mappingModelCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetModelCategoryEnumStringValues Enumerates the set of values in String for ModelCategoryEnum
func GetModelCategoryEnumStringValues() []string {
	return []string{
		"USER",
		"SERVICE",
	}
}

// GetMappingModelCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelCategoryEnum(val string) (ModelCategoryEnum, bool) {
	enum, ok := mappingModelCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
