// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DatabaseConnectionTypeEnum Enum with underlying type: string
type DatabaseConnectionTypeEnum string

// Set of constants representing the allowable values for DatabaseConnectionTypeEnum
const (
	DatabaseConnectionTypeCloudWallet DatabaseConnectionTypeEnum = "CLOUD_WALLET"
	DatabaseConnectionTypeCustomJdbc  DatabaseConnectionTypeEnum = "CUSTOM_JDBC"
)

var mappingDatabaseConnectionTypeEnum = map[string]DatabaseConnectionTypeEnum{
	"CLOUD_WALLET": DatabaseConnectionTypeCloudWallet,
	"CUSTOM_JDBC":  DatabaseConnectionTypeCustomJdbc,
}

var mappingDatabaseConnectionTypeEnumLowerCase = map[string]DatabaseConnectionTypeEnum{
	"cloud_wallet": DatabaseConnectionTypeCloudWallet,
	"custom_jdbc":  DatabaseConnectionTypeCustomJdbc,
}

// GetDatabaseConnectionTypeEnumValues Enumerates the set of values for DatabaseConnectionTypeEnum
func GetDatabaseConnectionTypeEnumValues() []DatabaseConnectionTypeEnum {
	values := make([]DatabaseConnectionTypeEnum, 0)
	for _, v := range mappingDatabaseConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionTypeEnumStringValues Enumerates the set of values in String for DatabaseConnectionTypeEnum
func GetDatabaseConnectionTypeEnumStringValues() []string {
	return []string{
		"CLOUD_WALLET",
		"CUSTOM_JDBC",
	}
}

// GetMappingDatabaseConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionTypeEnum(val string) (DatabaseConnectionTypeEnum, bool) {
	enum, ok := mappingDatabaseConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
