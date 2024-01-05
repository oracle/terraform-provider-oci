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

// DatastoreTypesEnum Enum with underlying type: string
type DatastoreTypesEnum string

// Set of constants representing the allowable values for DatastoreTypesEnum
const (
	DatastoreTypesManagement DatastoreTypesEnum = "MANAGEMENT"
	DatastoreTypesWorkload   DatastoreTypesEnum = "WORKLOAD"
)

var mappingDatastoreTypesEnum = map[string]DatastoreTypesEnum{
	"MANAGEMENT": DatastoreTypesManagement,
	"WORKLOAD":   DatastoreTypesWorkload,
}

var mappingDatastoreTypesEnumLowerCase = map[string]DatastoreTypesEnum{
	"management": DatastoreTypesManagement,
	"workload":   DatastoreTypesWorkload,
}

// GetDatastoreTypesEnumValues Enumerates the set of values for DatastoreTypesEnum
func GetDatastoreTypesEnumValues() []DatastoreTypesEnum {
	values := make([]DatastoreTypesEnum, 0)
	for _, v := range mappingDatastoreTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatastoreTypesEnumStringValues Enumerates the set of values in String for DatastoreTypesEnum
func GetDatastoreTypesEnumStringValues() []string {
	return []string{
		"MANAGEMENT",
		"WORKLOAD",
	}
}

// GetMappingDatastoreTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatastoreTypesEnum(val string) (DatastoreTypesEnum, bool) {
	enum, ok := mappingDatastoreTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
