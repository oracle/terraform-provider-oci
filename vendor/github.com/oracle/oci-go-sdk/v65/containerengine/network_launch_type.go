// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"strings"
)

// NetworkLaunchTypeEnum Enum with underlying type: string
type NetworkLaunchTypeEnum string

// Set of constants representing the allowable values for NetworkLaunchTypeEnum
const (
	NetworkLaunchTypeVfio            NetworkLaunchTypeEnum = "VFIO"
	NetworkLaunchTypeE1000           NetworkLaunchTypeEnum = "E1000"
	NetworkLaunchTypeParavirtualized NetworkLaunchTypeEnum = "PARAVIRTUALIZED"
)

var mappingNetworkLaunchTypeEnum = map[string]NetworkLaunchTypeEnum{
	"VFIO":            NetworkLaunchTypeVfio,
	"E1000":           NetworkLaunchTypeE1000,
	"PARAVIRTUALIZED": NetworkLaunchTypeParavirtualized,
}

var mappingNetworkLaunchTypeEnumLowerCase = map[string]NetworkLaunchTypeEnum{
	"vfio":            NetworkLaunchTypeVfio,
	"e1000":           NetworkLaunchTypeE1000,
	"paravirtualized": NetworkLaunchTypeParavirtualized,
}

// GetNetworkLaunchTypeEnumValues Enumerates the set of values for NetworkLaunchTypeEnum
func GetNetworkLaunchTypeEnumValues() []NetworkLaunchTypeEnum {
	values := make([]NetworkLaunchTypeEnum, 0)
	for _, v := range mappingNetworkLaunchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkLaunchTypeEnumStringValues Enumerates the set of values in String for NetworkLaunchTypeEnum
func GetNetworkLaunchTypeEnumStringValues() []string {
	return []string{
		"VFIO",
		"E1000",
		"PARAVIRTUALIZED",
	}
}

// GetMappingNetworkLaunchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkLaunchTypeEnum(val string) (NetworkLaunchTypeEnum, bool) {
	enum, ok := mappingNetworkLaunchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
