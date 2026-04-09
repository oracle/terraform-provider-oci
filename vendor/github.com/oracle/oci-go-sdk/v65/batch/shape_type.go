// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"strings"
)

// ShapeTypeEnum Enum with underlying type: string
type ShapeTypeEnum string

// Set of constants representing the allowable values for ShapeTypeEnum
const (
	ShapeTypeCpu ShapeTypeEnum = "CPU"
	ShapeTypeGpu ShapeTypeEnum = "GPU"
)

var mappingShapeTypeEnum = map[string]ShapeTypeEnum{
	"CPU": ShapeTypeCpu,
	"GPU": ShapeTypeGpu,
}

var mappingShapeTypeEnumLowerCase = map[string]ShapeTypeEnum{
	"cpu": ShapeTypeCpu,
	"gpu": ShapeTypeGpu,
}

// GetShapeTypeEnumValues Enumerates the set of values for ShapeTypeEnum
func GetShapeTypeEnumValues() []ShapeTypeEnum {
	values := make([]ShapeTypeEnum, 0)
	for _, v := range mappingShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetShapeTypeEnumStringValues Enumerates the set of values in String for ShapeTypeEnum
func GetShapeTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"GPU",
	}
}

// GetMappingShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShapeTypeEnum(val string) (ShapeTypeEnum, bool) {
	enum, ok := mappingShapeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
