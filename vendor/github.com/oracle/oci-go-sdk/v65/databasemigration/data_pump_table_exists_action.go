// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// DataPumpTableExistsActionEnum Enum with underlying type: string
type DataPumpTableExistsActionEnum string

// Set of constants representing the allowable values for DataPumpTableExistsActionEnum
const (
	DataPumpTableExistsActionTruncate DataPumpTableExistsActionEnum = "TRUNCATE"
	DataPumpTableExistsActionReplace  DataPumpTableExistsActionEnum = "REPLACE"
	DataPumpTableExistsActionAppend   DataPumpTableExistsActionEnum = "APPEND"
	DataPumpTableExistsActionSkip     DataPumpTableExistsActionEnum = "SKIP"
)

var mappingDataPumpTableExistsActionEnum = map[string]DataPumpTableExistsActionEnum{
	"TRUNCATE": DataPumpTableExistsActionTruncate,
	"REPLACE":  DataPumpTableExistsActionReplace,
	"APPEND":   DataPumpTableExistsActionAppend,
	"SKIP":     DataPumpTableExistsActionSkip,
}

var mappingDataPumpTableExistsActionEnumLowerCase = map[string]DataPumpTableExistsActionEnum{
	"truncate": DataPumpTableExistsActionTruncate,
	"replace":  DataPumpTableExistsActionReplace,
	"append":   DataPumpTableExistsActionAppend,
	"skip":     DataPumpTableExistsActionSkip,
}

// GetDataPumpTableExistsActionEnumValues Enumerates the set of values for DataPumpTableExistsActionEnum
func GetDataPumpTableExistsActionEnumValues() []DataPumpTableExistsActionEnum {
	values := make([]DataPumpTableExistsActionEnum, 0)
	for _, v := range mappingDataPumpTableExistsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDataPumpTableExistsActionEnumStringValues Enumerates the set of values in String for DataPumpTableExistsActionEnum
func GetDataPumpTableExistsActionEnumStringValues() []string {
	return []string{
		"TRUNCATE",
		"REPLACE",
		"APPEND",
		"SKIP",
	}
}

// GetMappingDataPumpTableExistsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataPumpTableExistsActionEnum(val string) (DataPumpTableExistsActionEnum, bool) {
	enum, ok := mappingDataPumpTableExistsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
