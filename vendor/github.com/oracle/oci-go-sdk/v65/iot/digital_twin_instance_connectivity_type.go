// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"strings"
)

// DigitalTwinInstanceConnectivityTypeEnum Enum with underlying type: string
type DigitalTwinInstanceConnectivityTypeEnum string

// Set of constants representing the allowable values for DigitalTwinInstanceConnectivityTypeEnum
const (
	DigitalTwinInstanceConnectivityTypeDirect   DigitalTwinInstanceConnectivityTypeEnum = "DIRECT"
	DigitalTwinInstanceConnectivityTypeIndirect DigitalTwinInstanceConnectivityTypeEnum = "INDIRECT"
	DigitalTwinInstanceConnectivityTypeGateway  DigitalTwinInstanceConnectivityTypeEnum = "GATEWAY"
	DigitalTwinInstanceConnectivityTypeNone     DigitalTwinInstanceConnectivityTypeEnum = "NONE"
)

var mappingDigitalTwinInstanceConnectivityTypeEnum = map[string]DigitalTwinInstanceConnectivityTypeEnum{
	"DIRECT":   DigitalTwinInstanceConnectivityTypeDirect,
	"INDIRECT": DigitalTwinInstanceConnectivityTypeIndirect,
	"GATEWAY":  DigitalTwinInstanceConnectivityTypeGateway,
	"NONE":     DigitalTwinInstanceConnectivityTypeNone,
}

var mappingDigitalTwinInstanceConnectivityTypeEnumLowerCase = map[string]DigitalTwinInstanceConnectivityTypeEnum{
	"direct":   DigitalTwinInstanceConnectivityTypeDirect,
	"indirect": DigitalTwinInstanceConnectivityTypeIndirect,
	"gateway":  DigitalTwinInstanceConnectivityTypeGateway,
	"none":     DigitalTwinInstanceConnectivityTypeNone,
}

// GetDigitalTwinInstanceConnectivityTypeEnumValues Enumerates the set of values for DigitalTwinInstanceConnectivityTypeEnum
func GetDigitalTwinInstanceConnectivityTypeEnumValues() []DigitalTwinInstanceConnectivityTypeEnum {
	values := make([]DigitalTwinInstanceConnectivityTypeEnum, 0)
	for _, v := range mappingDigitalTwinInstanceConnectivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDigitalTwinInstanceConnectivityTypeEnumStringValues Enumerates the set of values in String for DigitalTwinInstanceConnectivityTypeEnum
func GetDigitalTwinInstanceConnectivityTypeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"INDIRECT",
		"GATEWAY",
		"NONE",
	}
}

// GetMappingDigitalTwinInstanceConnectivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDigitalTwinInstanceConnectivityTypeEnum(val string) (DigitalTwinInstanceConnectivityTypeEnum, bool) {
	enum, ok := mappingDigitalTwinInstanceConnectivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
