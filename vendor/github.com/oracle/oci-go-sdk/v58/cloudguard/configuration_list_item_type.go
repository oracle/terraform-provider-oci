// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// ConfigurationListItemTypeEnum Enum with underlying type: string
type ConfigurationListItemTypeEnum string

// Set of constants representing the allowable values for ConfigurationListItemTypeEnum
const (
	ConfigurationListItemTypeManaged ConfigurationListItemTypeEnum = "MANAGED"
	ConfigurationListItemTypeCustom  ConfigurationListItemTypeEnum = "CUSTOM"
)

var mappingConfigurationListItemTypeEnum = map[string]ConfigurationListItemTypeEnum{
	"MANAGED": ConfigurationListItemTypeManaged,
	"CUSTOM":  ConfigurationListItemTypeCustom,
}

// GetConfigurationListItemTypeEnumValues Enumerates the set of values for ConfigurationListItemTypeEnum
func GetConfigurationListItemTypeEnumValues() []ConfigurationListItemTypeEnum {
	values := make([]ConfigurationListItemTypeEnum, 0)
	for _, v := range mappingConfigurationListItemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationListItemTypeEnumStringValues Enumerates the set of values in String for ConfigurationListItemTypeEnum
func GetConfigurationListItemTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingConfigurationListItemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationListItemTypeEnum(val string) (ConfigurationListItemTypeEnum, bool) {
	mappingConfigurationListItemTypeEnumIgnoreCase := make(map[string]ConfigurationListItemTypeEnum)
	for k, v := range mappingConfigurationListItemTypeEnum {
		mappingConfigurationListItemTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConfigurationListItemTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
