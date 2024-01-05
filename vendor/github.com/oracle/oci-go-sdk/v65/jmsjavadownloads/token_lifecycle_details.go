// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// TokenLifecycleDetailsEnum Enum with underlying type: string
type TokenLifecycleDetailsEnum string

// Set of constants representing the allowable values for TokenLifecycleDetailsEnum
const (
	TokenLifecycleDetailsExpired  TokenLifecycleDetailsEnum = "EXPIRED"
	TokenLifecycleDetailsRevoking TokenLifecycleDetailsEnum = "REVOKING"
	TokenLifecycleDetailsRevoked  TokenLifecycleDetailsEnum = "REVOKED"
)

var mappingTokenLifecycleDetailsEnum = map[string]TokenLifecycleDetailsEnum{
	"EXPIRED":  TokenLifecycleDetailsExpired,
	"REVOKING": TokenLifecycleDetailsRevoking,
	"REVOKED":  TokenLifecycleDetailsRevoked,
}

var mappingTokenLifecycleDetailsEnumLowerCase = map[string]TokenLifecycleDetailsEnum{
	"expired":  TokenLifecycleDetailsExpired,
	"revoking": TokenLifecycleDetailsRevoking,
	"revoked":  TokenLifecycleDetailsRevoked,
}

// GetTokenLifecycleDetailsEnumValues Enumerates the set of values for TokenLifecycleDetailsEnum
func GetTokenLifecycleDetailsEnumValues() []TokenLifecycleDetailsEnum {
	values := make([]TokenLifecycleDetailsEnum, 0)
	for _, v := range mappingTokenLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetTokenLifecycleDetailsEnumStringValues Enumerates the set of values in String for TokenLifecycleDetailsEnum
func GetTokenLifecycleDetailsEnumStringValues() []string {
	return []string{
		"EXPIRED",
		"REVOKING",
		"REVOKED",
	}
}

// GetMappingTokenLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTokenLifecycleDetailsEnum(val string) (TokenLifecycleDetailsEnum, bool) {
	enum, ok := mappingTokenLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
