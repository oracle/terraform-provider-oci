// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// EntityTypeEnum Enum with underlying type: string
type EntityTypeEnum string

// Set of constants representing the allowable values for EntityTypeEnum
const (
	EntityTypeExternalIp EntityTypeEnum = "EXTERNAL_IP"
	EntityTypeInternalIp EntityTypeEnum = "INTERNAL_IP"
	EntityTypeText       EntityTypeEnum = "TEXT"
	EntityTypeJsonList   EntityTypeEnum = "JSON_LIST"
)

var mappingEntityTypeEnum = map[string]EntityTypeEnum{
	"EXTERNAL_IP": EntityTypeExternalIp,
	"INTERNAL_IP": EntityTypeInternalIp,
	"TEXT":        EntityTypeText,
	"JSON_LIST":   EntityTypeJsonList,
}

var mappingEntityTypeEnumLowerCase = map[string]EntityTypeEnum{
	"external_ip": EntityTypeExternalIp,
	"internal_ip": EntityTypeInternalIp,
	"text":        EntityTypeText,
	"json_list":   EntityTypeJsonList,
}

// GetEntityTypeEnumValues Enumerates the set of values for EntityTypeEnum
func GetEntityTypeEnumValues() []EntityTypeEnum {
	values := make([]EntityTypeEnum, 0)
	for _, v := range mappingEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityTypeEnumStringValues Enumerates the set of values in String for EntityTypeEnum
func GetEntityTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL_IP",
		"INTERNAL_IP",
		"TEXT",
		"JSON_LIST",
	}
}

// GetMappingEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityTypeEnum(val string) (EntityTypeEnum, bool) {
	enum, ok := mappingEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
