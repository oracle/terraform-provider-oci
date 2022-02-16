// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ConnectionResultEnum Enum with underlying type: string
type ConnectionResultEnum string

// Set of constants representing the allowable values for ConnectionResultEnum
const (
	ConnectionResultSucceeded ConnectionResultEnum = "SUCCEEDED"
	ConnectionResultFailed    ConnectionResultEnum = "FAILED"
)

var mappingConnectionResultEnum = map[string]ConnectionResultEnum{
	"SUCCEEDED": ConnectionResultSucceeded,
	"FAILED":    ConnectionResultFailed,
}

// GetConnectionResultEnumValues Enumerates the set of values for ConnectionResultEnum
func GetConnectionResultEnumValues() []ConnectionResultEnum {
	values := make([]ConnectionResultEnum, 0)
	for _, v := range mappingConnectionResultEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionResultEnumStringValues Enumerates the set of values in String for ConnectionResultEnum
func GetConnectionResultEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingConnectionResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionResultEnum(val string) (ConnectionResultEnum, bool) {
	mappingConnectionResultEnumIgnoreCase := make(map[string]ConnectionResultEnum)
	for k, v := range mappingConnectionResultEnum {
		mappingConnectionResultEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConnectionResultEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
