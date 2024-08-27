// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// DatastoreClusterTypesEnum Enum with underlying type: string
type DatastoreClusterTypesEnum string

// Set of constants representing the allowable values for DatastoreClusterTypesEnum
const (
	DatastoreClusterTypesManagement DatastoreClusterTypesEnum = "MANAGEMENT"
	DatastoreClusterTypesWorkload   DatastoreClusterTypesEnum = "WORKLOAD"
)

var mappingDatastoreClusterTypesEnum = map[string]DatastoreClusterTypesEnum{
	"MANAGEMENT": DatastoreClusterTypesManagement,
	"WORKLOAD":   DatastoreClusterTypesWorkload,
}

var mappingDatastoreClusterTypesEnumLowerCase = map[string]DatastoreClusterTypesEnum{
	"management": DatastoreClusterTypesManagement,
	"workload":   DatastoreClusterTypesWorkload,
}

// GetDatastoreClusterTypesEnumValues Enumerates the set of values for DatastoreClusterTypesEnum
func GetDatastoreClusterTypesEnumValues() []DatastoreClusterTypesEnum {
	values := make([]DatastoreClusterTypesEnum, 0)
	for _, v := range mappingDatastoreClusterTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatastoreClusterTypesEnumStringValues Enumerates the set of values in String for DatastoreClusterTypesEnum
func GetDatastoreClusterTypesEnumStringValues() []string {
	return []string{
		"MANAGEMENT",
		"WORKLOAD",
	}
}

// GetMappingDatastoreClusterTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatastoreClusterTypesEnum(val string) (DatastoreClusterTypesEnum, bool) {
	enum, ok := mappingDatastoreClusterTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
