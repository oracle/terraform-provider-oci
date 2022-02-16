// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// ServiceListEnum Enum with underlying type: string
type ServiceListEnum string

// Set of constants representing the allowable values for ServiceListEnum
const (
	ServiceListDataSafe      ServiceListEnum = "DataSafe"
	ServiceListDataSafeDev   ServiceListEnum = "DataSafe-dev"
	ServiceListDataSafeDev1  ServiceListEnum = "DataSafe-dev1"
	ServiceListDataSafeDev2  ServiceListEnum = "DataSafe-dev2"
	ServiceListDataSafeDev3  ServiceListEnum = "DataSafe-dev3"
	ServiceListDataSafeDev4  ServiceListEnum = "DataSafe-dev4"
	ServiceListDataSafeDev5  ServiceListEnum = "DataSafe-dev5"
	ServiceListDataSafeDev6  ServiceListEnum = "DataSafe-dev6"
	ServiceListDataSafeDev7  ServiceListEnum = "DataSafe-dev7"
	ServiceListDataSafeDev8  ServiceListEnum = "DataSafe-dev8"
	ServiceListDataSafeDev9  ServiceListEnum = "DataSafe-dev9"
	ServiceListDataSafeLrg1  ServiceListEnum = "DataSafe-lrg1"
	ServiceListDataSafeLrg2  ServiceListEnum = "DataSafe-lrg2"
	ServiceListDataSafeLrg3  ServiceListEnum = "DataSafe-lrg3"
	ServiceListDataSafeLrg4  ServiceListEnum = "DataSafe-lrg4"
	ServiceListDataSafePtest ServiceListEnum = "DataSafe-ptest"
	ServiceListDataSafeStest ServiceListEnum = "DataSafe-stest"
	ServiceListDataSafeStage ServiceListEnum = "DataSafe-stage"
)

var mappingServiceListEnum = map[string]ServiceListEnum{
	"DataSafe":       ServiceListDataSafe,
	"DataSafe-dev":   ServiceListDataSafeDev,
	"DataSafe-dev1":  ServiceListDataSafeDev1,
	"DataSafe-dev2":  ServiceListDataSafeDev2,
	"DataSafe-dev3":  ServiceListDataSafeDev3,
	"DataSafe-dev4":  ServiceListDataSafeDev4,
	"DataSafe-dev5":  ServiceListDataSafeDev5,
	"DataSafe-dev6":  ServiceListDataSafeDev6,
	"DataSafe-dev7":  ServiceListDataSafeDev7,
	"DataSafe-dev8":  ServiceListDataSafeDev8,
	"DataSafe-dev9":  ServiceListDataSafeDev9,
	"DataSafe-lrg1":  ServiceListDataSafeLrg1,
	"DataSafe-lrg2":  ServiceListDataSafeLrg2,
	"DataSafe-lrg3":  ServiceListDataSafeLrg3,
	"DataSafe-lrg4":  ServiceListDataSafeLrg4,
	"DataSafe-ptest": ServiceListDataSafePtest,
	"DataSafe-stest": ServiceListDataSafeStest,
	"DataSafe-stage": ServiceListDataSafeStage,
}

// GetServiceListEnumValues Enumerates the set of values for ServiceListEnum
func GetServiceListEnumValues() []ServiceListEnum {
	values := make([]ServiceListEnum, 0)
	for _, v := range mappingServiceListEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceListEnumStringValues Enumerates the set of values in String for ServiceListEnum
func GetServiceListEnumStringValues() []string {
	return []string{
		"DataSafe",
		"DataSafe-dev",
		"DataSafe-dev1",
		"DataSafe-dev2",
		"DataSafe-dev3",
		"DataSafe-dev4",
		"DataSafe-dev5",
		"DataSafe-dev6",
		"DataSafe-dev7",
		"DataSafe-dev8",
		"DataSafe-dev9",
		"DataSafe-lrg1",
		"DataSafe-lrg2",
		"DataSafe-lrg3",
		"DataSafe-lrg4",
		"DataSafe-ptest",
		"DataSafe-stest",
		"DataSafe-stage",
	}
}

// GetMappingServiceListEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceListEnum(val string) (ServiceListEnum, bool) {
	mappingServiceListEnumIgnoreCase := make(map[string]ServiceListEnum)
	for k, v := range mappingServiceListEnum {
		mappingServiceListEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingServiceListEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
