// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// DataScienceResourceTypeEnum Enum with underlying type: string
type DataScienceResourceTypeEnum string

// Set of constants representing the allowable values for DataScienceResourceTypeEnum
const (
	DataScienceResourceTypeNotebookSession DataScienceResourceTypeEnum = "NOTEBOOK_SESSION"
)

var mappingDataScienceResourceTypeEnum = map[string]DataScienceResourceTypeEnum{
	"NOTEBOOK_SESSION": DataScienceResourceTypeNotebookSession,
}

var mappingDataScienceResourceTypeEnumLowerCase = map[string]DataScienceResourceTypeEnum{
	"notebook_session": DataScienceResourceTypeNotebookSession,
}

// GetDataScienceResourceTypeEnumValues Enumerates the set of values for DataScienceResourceTypeEnum
func GetDataScienceResourceTypeEnumValues() []DataScienceResourceTypeEnum {
	values := make([]DataScienceResourceTypeEnum, 0)
	for _, v := range mappingDataScienceResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataScienceResourceTypeEnumStringValues Enumerates the set of values in String for DataScienceResourceTypeEnum
func GetDataScienceResourceTypeEnumStringValues() []string {
	return []string{
		"NOTEBOOK_SESSION",
	}
}

// GetMappingDataScienceResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataScienceResourceTypeEnum(val string) (DataScienceResourceTypeEnum, bool) {
	enum, ok := mappingDataScienceResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
