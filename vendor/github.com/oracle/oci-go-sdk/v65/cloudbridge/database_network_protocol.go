// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// DatabaseNetworkProtocolEnum Enum with underlying type: string
type DatabaseNetworkProtocolEnum string

// Set of constants representing the allowable values for DatabaseNetworkProtocolEnum
const (
	DatabaseNetworkProtocolTcp  DatabaseNetworkProtocolEnum = "TCP"
	DatabaseNetworkProtocolTcps DatabaseNetworkProtocolEnum = "TCPS"
)

var mappingDatabaseNetworkProtocolEnum = map[string]DatabaseNetworkProtocolEnum{
	"TCP":  DatabaseNetworkProtocolTcp,
	"TCPS": DatabaseNetworkProtocolTcps,
}

var mappingDatabaseNetworkProtocolEnumLowerCase = map[string]DatabaseNetworkProtocolEnum{
	"tcp":  DatabaseNetworkProtocolTcp,
	"tcps": DatabaseNetworkProtocolTcps,
}

// GetDatabaseNetworkProtocolEnumValues Enumerates the set of values for DatabaseNetworkProtocolEnum
func GetDatabaseNetworkProtocolEnumValues() []DatabaseNetworkProtocolEnum {
	values := make([]DatabaseNetworkProtocolEnum, 0)
	for _, v := range mappingDatabaseNetworkProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseNetworkProtocolEnumStringValues Enumerates the set of values in String for DatabaseNetworkProtocolEnum
func GetDatabaseNetworkProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingDatabaseNetworkProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseNetworkProtocolEnum(val string) (DatabaseNetworkProtocolEnum, bool) {
	enum, ok := mappingDatabaseNetworkProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
