// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"strings"
)

// DataSourceTypeEnum Enum with underlying type: string
type DataSourceTypeEnum string

// Set of constants representing the allowable values for DataSourceTypeEnum
const (
	DataSourceTypeOracleObjectStorage DataSourceTypeEnum = "ORACLE_OBJECT_STORAGE"
	DataSourceTypeOracleAtp           DataSourceTypeEnum = "ORACLE_ATP"
	DataSourceTypeInflux              DataSourceTypeEnum = "INFLUX"
)

var mappingDataSourceTypeEnum = map[string]DataSourceTypeEnum{
	"ORACLE_OBJECT_STORAGE": DataSourceTypeOracleObjectStorage,
	"ORACLE_ATP":            DataSourceTypeOracleAtp,
	"INFLUX":                DataSourceTypeInflux,
}

var mappingDataSourceTypeEnumLowerCase = map[string]DataSourceTypeEnum{
	"oracle_object_storage": DataSourceTypeOracleObjectStorage,
	"oracle_atp":            DataSourceTypeOracleAtp,
	"influx":                DataSourceTypeInflux,
}

// GetDataSourceTypeEnumValues Enumerates the set of values for DataSourceTypeEnum
func GetDataSourceTypeEnumValues() []DataSourceTypeEnum {
	values := make([]DataSourceTypeEnum, 0)
	for _, v := range mappingDataSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSourceTypeEnumStringValues Enumerates the set of values in String for DataSourceTypeEnum
func GetDataSourceTypeEnumStringValues() []string {
	return []string{
		"ORACLE_OBJECT_STORAGE",
		"ORACLE_ATP",
		"INFLUX",
	}
}

// GetMappingDataSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSourceTypeEnum(val string) (DataSourceTypeEnum, bool) {
	enum, ok := mappingDataSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
