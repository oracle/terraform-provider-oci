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

// RecipeTypeEnum Enum with underlying type: string
type RecipeTypeEnum string

// Set of constants representing the allowable values for RecipeTypeEnum
const (
	RecipeTypeZeroEtl RecipeTypeEnum = "ZERO_ETL"
)

var mappingRecipeTypeEnum = map[string]RecipeTypeEnum{
	"ZERO_ETL": RecipeTypeZeroEtl,
}

var mappingRecipeTypeEnumLowerCase = map[string]RecipeTypeEnum{
	"zero_etl": RecipeTypeZeroEtl,
}

// GetRecipeTypeEnumValues Enumerates the set of values for RecipeTypeEnum
func GetRecipeTypeEnumValues() []RecipeTypeEnum {
	values := make([]RecipeTypeEnum, 0)
	for _, v := range mappingRecipeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecipeTypeEnumStringValues Enumerates the set of values in String for RecipeTypeEnum
func GetRecipeTypeEnumStringValues() []string {
	return []string{
		"ZERO_ETL",
	}
}

// GetMappingRecipeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecipeTypeEnum(val string) (RecipeTypeEnum, bool) {
	enum, ok := mappingRecipeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
