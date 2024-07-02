// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// DataObjectTypeEnum Enum with underlying type: string
type DataObjectTypeEnum string

// Set of constants representing the allowable values for DataObjectTypeEnum
const (
	DataObjectTypeView  DataObjectTypeEnum = "VIEW"
	DataObjectTypeTable DataObjectTypeEnum = "TABLE"
)

var mappingDataObjectTypeEnum = map[string]DataObjectTypeEnum{
	"VIEW":  DataObjectTypeView,
	"TABLE": DataObjectTypeTable,
}

var mappingDataObjectTypeEnumLowerCase = map[string]DataObjectTypeEnum{
	"view":  DataObjectTypeView,
	"table": DataObjectTypeTable,
}

// GetDataObjectTypeEnumValues Enumerates the set of values for DataObjectTypeEnum
func GetDataObjectTypeEnumValues() []DataObjectTypeEnum {
	values := make([]DataObjectTypeEnum, 0)
	for _, v := range mappingDataObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataObjectTypeEnumStringValues Enumerates the set of values in String for DataObjectTypeEnum
func GetDataObjectTypeEnumStringValues() []string {
	return []string{
		"VIEW",
		"TABLE",
	}
}

// GetMappingDataObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataObjectTypeEnum(val string) (DataObjectTypeEnum, bool) {
	enum, ok := mappingDataObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
