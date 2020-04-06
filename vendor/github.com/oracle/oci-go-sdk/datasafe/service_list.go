// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// ServiceListEnum Enum with underlying type: string
type ServiceListEnum string

// Set of constants representing the allowable values for ServiceListEnum
const (
	ServiceListDataSafe      ServiceListEnum = "DataSafe"
	ServiceListDataSafeDev   ServiceListEnum = "DataSafe-dev"
	ServiceListDataSafeDev1  ServiceListEnum = "DataSafe-dev1"
	ServiceListDataSafeDev2  ServiceListEnum = "DataSafe-dev2"
	ServiceListDataSafeDev3  ServiceListEnum = "DataSafe-dev3"
	ServiceListDataSafeLrg1  ServiceListEnum = "DataSafe-lrg1"
	ServiceListDataSafeLrg2  ServiceListEnum = "DataSafe-lrg2"
	ServiceListDataSafeLrg3  ServiceListEnum = "DataSafe-lrg3"
	ServiceListDataSafeLrg4  ServiceListEnum = "DataSafe-lrg4"
	ServiceListDataSafePtest ServiceListEnum = "DataSafe-ptest"
	ServiceListDataSafeStest ServiceListEnum = "DataSafe-stest"
	ServiceListDataSafeStage ServiceListEnum = "DataSafe-stage"
)

var mappingServiceList = map[string]ServiceListEnum{
	"DataSafe":       ServiceListDataSafe,
	"DataSafe-dev":   ServiceListDataSafeDev,
	"DataSafe-dev1":  ServiceListDataSafeDev1,
	"DataSafe-dev2":  ServiceListDataSafeDev2,
	"DataSafe-dev3":  ServiceListDataSafeDev3,
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
	for _, v := range mappingServiceList {
		values = append(values, v)
	}
	return values
}
