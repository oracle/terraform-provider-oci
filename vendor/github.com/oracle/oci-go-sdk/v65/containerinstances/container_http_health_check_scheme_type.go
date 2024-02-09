// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"strings"
)

// ContainerHttpHealthCheckSchemeTypeEnum Enum with underlying type: string
type ContainerHttpHealthCheckSchemeTypeEnum string

// Set of constants representing the allowable values for ContainerHttpHealthCheckSchemeTypeEnum
const (
	ContainerHttpHealthCheckSchemeTypeHttp  ContainerHttpHealthCheckSchemeTypeEnum = "HTTP"
	ContainerHttpHealthCheckSchemeTypeHttps ContainerHttpHealthCheckSchemeTypeEnum = "HTTPS"
)

var mappingContainerHttpHealthCheckSchemeTypeEnum = map[string]ContainerHttpHealthCheckSchemeTypeEnum{
	"HTTP":  ContainerHttpHealthCheckSchemeTypeHttp,
	"HTTPS": ContainerHttpHealthCheckSchemeTypeHttps,
}

var mappingContainerHttpHealthCheckSchemeTypeEnumLowerCase = map[string]ContainerHttpHealthCheckSchemeTypeEnum{
	"http":  ContainerHttpHealthCheckSchemeTypeHttp,
	"https": ContainerHttpHealthCheckSchemeTypeHttps,
}

// GetContainerHttpHealthCheckSchemeTypeEnumValues Enumerates the set of values for ContainerHttpHealthCheckSchemeTypeEnum
func GetContainerHttpHealthCheckSchemeTypeEnumValues() []ContainerHttpHealthCheckSchemeTypeEnum {
	values := make([]ContainerHttpHealthCheckSchemeTypeEnum, 0)
	for _, v := range mappingContainerHttpHealthCheckSchemeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerHttpHealthCheckSchemeTypeEnumStringValues Enumerates the set of values in String for ContainerHttpHealthCheckSchemeTypeEnum
func GetContainerHttpHealthCheckSchemeTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"HTTPS",
	}
}

// GetMappingContainerHttpHealthCheckSchemeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerHttpHealthCheckSchemeTypeEnum(val string) (ContainerHttpHealthCheckSchemeTypeEnum, bool) {
	enum, ok := mappingContainerHttpHealthCheckSchemeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
