// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// DataAssetImportExportTypeFilterEnum Enum with underlying type: string
type DataAssetImportExportTypeFilterEnum string

// Set of constants representing the allowable values for DataAssetImportExportTypeFilterEnum
const (
	DataAssetImportExportTypeFilterCustomPropertyValues DataAssetImportExportTypeFilterEnum = "CUSTOM_PROPERTY_VALUES"
	DataAssetImportExportTypeFilterAll                  DataAssetImportExportTypeFilterEnum = "ALL"
)

var mappingDataAssetImportExportTypeFilterEnum = map[string]DataAssetImportExportTypeFilterEnum{
	"CUSTOM_PROPERTY_VALUES": DataAssetImportExportTypeFilterCustomPropertyValues,
	"ALL":                    DataAssetImportExportTypeFilterAll,
}

var mappingDataAssetImportExportTypeFilterEnumLowerCase = map[string]DataAssetImportExportTypeFilterEnum{
	"custom_property_values": DataAssetImportExportTypeFilterCustomPropertyValues,
	"all":                    DataAssetImportExportTypeFilterAll,
}

// GetDataAssetImportExportTypeFilterEnumValues Enumerates the set of values for DataAssetImportExportTypeFilterEnum
func GetDataAssetImportExportTypeFilterEnumValues() []DataAssetImportExportTypeFilterEnum {
	values := make([]DataAssetImportExportTypeFilterEnum, 0)
	for _, v := range mappingDataAssetImportExportTypeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetDataAssetImportExportTypeFilterEnumStringValues Enumerates the set of values in String for DataAssetImportExportTypeFilterEnum
func GetDataAssetImportExportTypeFilterEnumStringValues() []string {
	return []string{
		"CUSTOM_PROPERTY_VALUES",
		"ALL",
	}
}

// GetMappingDataAssetImportExportTypeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataAssetImportExportTypeFilterEnum(val string) (DataAssetImportExportTypeFilterEnum, bool) {
	enum, ok := mappingDataAssetImportExportTypeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
