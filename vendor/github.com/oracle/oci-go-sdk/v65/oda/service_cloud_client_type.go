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

// ServiceCloudClientTypeEnum Enum with underlying type: string
type ServiceCloudClientTypeEnum string

// Set of constants representing the allowable values for ServiceCloudClientTypeEnum
const (
	ServiceCloudClientTypeWsdl ServiceCloudClientTypeEnum = "WSDL"
	ServiceCloudClientTypeRest ServiceCloudClientTypeEnum = "REST"
)

var mappingServiceCloudClientTypeEnum = map[string]ServiceCloudClientTypeEnum{
	"WSDL": ServiceCloudClientTypeWsdl,
	"REST": ServiceCloudClientTypeRest,
}

var mappingServiceCloudClientTypeEnumLowerCase = map[string]ServiceCloudClientTypeEnum{
	"wsdl": ServiceCloudClientTypeWsdl,
	"rest": ServiceCloudClientTypeRest,
}

// GetServiceCloudClientTypeEnumValues Enumerates the set of values for ServiceCloudClientTypeEnum
func GetServiceCloudClientTypeEnumValues() []ServiceCloudClientTypeEnum {
	values := make([]ServiceCloudClientTypeEnum, 0)
	for _, v := range mappingServiceCloudClientTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceCloudClientTypeEnumStringValues Enumerates the set of values in String for ServiceCloudClientTypeEnum
func GetServiceCloudClientTypeEnumStringValues() []string {
	return []string{
		"WSDL",
		"REST",
	}
}

// GetMappingServiceCloudClientTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceCloudClientTypeEnum(val string) (ServiceCloudClientTypeEnum, bool) {
	enum, ok := mappingServiceCloudClientTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
