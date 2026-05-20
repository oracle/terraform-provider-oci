// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// SslModeEnum Enum with underlying type: string
type SslModeEnum string

// Set of constants representing the allowable values for SslModeEnum
const (
	SslModeVerifyIdentity SslModeEnum = "VERIFY_IDENTITY"
	SslModeVerifyCa       SslModeEnum = "VERIFY_CA"
	SslModeRequired       SslModeEnum = "REQUIRED"
	SslModeDisabled       SslModeEnum = "DISABLED"
)

var mappingSslModeEnum = map[string]SslModeEnum{
	"VERIFY_IDENTITY": SslModeVerifyIdentity,
	"VERIFY_CA":       SslModeVerifyCa,
	"REQUIRED":        SslModeRequired,
	"DISABLED":        SslModeDisabled,
}

var mappingSslModeEnumLowerCase = map[string]SslModeEnum{
	"verify_identity": SslModeVerifyIdentity,
	"verify_ca":       SslModeVerifyCa,
	"required":        SslModeRequired,
	"disabled":        SslModeDisabled,
}

// GetSslModeEnumValues Enumerates the set of values for SslModeEnum
func GetSslModeEnumValues() []SslModeEnum {
	values := make([]SslModeEnum, 0)
	for _, v := range mappingSslModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSslModeEnumStringValues Enumerates the set of values in String for SslModeEnum
func GetSslModeEnumStringValues() []string {
	return []string{
		"VERIFY_IDENTITY",
		"VERIFY_CA",
		"REQUIRED",
		"DISABLED",
	}
}

// GetMappingSslModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSslModeEnum(val string) (SslModeEnum, bool) {
	enum, ok := mappingSslModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
