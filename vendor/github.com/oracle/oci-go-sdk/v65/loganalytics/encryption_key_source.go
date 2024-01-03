// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// EncryptionKeySourceEnum Enum with underlying type: string
type EncryptionKeySourceEnum string

// Set of constants representing the allowable values for EncryptionKeySourceEnum
const (
	EncryptionKeySourceOracleManaged   EncryptionKeySourceEnum = "ORACLE_MANAGED"
	EncryptionKeySourceCustomerManaged EncryptionKeySourceEnum = "CUSTOMER_MANAGED"
)

var mappingEncryptionKeySourceEnum = map[string]EncryptionKeySourceEnum{
	"ORACLE_MANAGED":   EncryptionKeySourceOracleManaged,
	"CUSTOMER_MANAGED": EncryptionKeySourceCustomerManaged,
}

var mappingEncryptionKeySourceEnumLowerCase = map[string]EncryptionKeySourceEnum{
	"oracle_managed":   EncryptionKeySourceOracleManaged,
	"customer_managed": EncryptionKeySourceCustomerManaged,
}

// GetEncryptionKeySourceEnumValues Enumerates the set of values for EncryptionKeySourceEnum
func GetEncryptionKeySourceEnumValues() []EncryptionKeySourceEnum {
	values := make([]EncryptionKeySourceEnum, 0)
	for _, v := range mappingEncryptionKeySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetEncryptionKeySourceEnumStringValues Enumerates the set of values in String for EncryptionKeySourceEnum
func GetEncryptionKeySourceEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingEncryptionKeySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEncryptionKeySourceEnum(val string) (EncryptionKeySourceEnum, bool) {
	enum, ok := mappingEncryptionKeySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
