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

// CreateBotKindEnum Enum with underlying type: string
type CreateBotKindEnum string

// Set of constants representing the allowable values for CreateBotKindEnum
const (
	CreateBotKindNew     CreateBotKindEnum = "NEW"
	CreateBotKindClone   CreateBotKindEnum = "CLONE"
	CreateBotKindVersion CreateBotKindEnum = "VERSION"
	CreateBotKindExtend  CreateBotKindEnum = "EXTEND"
)

var mappingCreateBotKindEnum = map[string]CreateBotKindEnum{
	"NEW":     CreateBotKindNew,
	"CLONE":   CreateBotKindClone,
	"VERSION": CreateBotKindVersion,
	"EXTEND":  CreateBotKindExtend,
}

var mappingCreateBotKindEnumLowerCase = map[string]CreateBotKindEnum{
	"new":     CreateBotKindNew,
	"clone":   CreateBotKindClone,
	"version": CreateBotKindVersion,
	"extend":  CreateBotKindExtend,
}

// GetCreateBotKindEnumValues Enumerates the set of values for CreateBotKindEnum
func GetCreateBotKindEnumValues() []CreateBotKindEnum {
	values := make([]CreateBotKindEnum, 0)
	for _, v := range mappingCreateBotKindEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBotKindEnumStringValues Enumerates the set of values in String for CreateBotKindEnum
func GetCreateBotKindEnumStringValues() []string {
	return []string{
		"NEW",
		"CLONE",
		"VERSION",
		"EXTEND",
	}
}

// GetMappingCreateBotKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBotKindEnum(val string) (CreateBotKindEnum, bool) {
	enum, ok := mappingCreateBotKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
