// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"strings"
)

// DemandSignalNamespaceEnum Enum with underlying type: string
type DemandSignalNamespaceEnum string

// Set of constants representing the allowable values for DemandSignalNamespaceEnum
const (
	DemandSignalNamespaceCompute DemandSignalNamespaceEnum = "COMPUTE"
	DemandSignalNamespaceNetwork DemandSignalNamespaceEnum = "NETWORK"
	DemandSignalNamespaceGpu     DemandSignalNamespaceEnum = "GPU"
	DemandSignalNamespaceStorage DemandSignalNamespaceEnum = "STORAGE"
)

var mappingDemandSignalNamespaceEnum = map[string]DemandSignalNamespaceEnum{
	"COMPUTE": DemandSignalNamespaceCompute,
	"NETWORK": DemandSignalNamespaceNetwork,
	"GPU":     DemandSignalNamespaceGpu,
	"STORAGE": DemandSignalNamespaceStorage,
}

var mappingDemandSignalNamespaceEnumLowerCase = map[string]DemandSignalNamespaceEnum{
	"compute": DemandSignalNamespaceCompute,
	"network": DemandSignalNamespaceNetwork,
	"gpu":     DemandSignalNamespaceGpu,
	"storage": DemandSignalNamespaceStorage,
}

// GetDemandSignalNamespaceEnumValues Enumerates the set of values for DemandSignalNamespaceEnum
func GetDemandSignalNamespaceEnumValues() []DemandSignalNamespaceEnum {
	values := make([]DemandSignalNamespaceEnum, 0)
	for _, v := range mappingDemandSignalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetDemandSignalNamespaceEnumStringValues Enumerates the set of values in String for DemandSignalNamespaceEnum
func GetDemandSignalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"NETWORK",
		"GPU",
		"STORAGE",
	}
}

// GetMappingDemandSignalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDemandSignalNamespaceEnum(val string) (DemandSignalNamespaceEnum, bool) {
	enum, ok := mappingDemandSignalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
