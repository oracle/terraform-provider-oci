// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"strings"
)

// DecryptionActionTypeEnum Enum with underlying type: string
type DecryptionActionTypeEnum string

// Set of constants representing the allowable values for DecryptionActionTypeEnum
const (
	DecryptionActionTypeNoDecrypt DecryptionActionTypeEnum = "NO_DECRYPT"
	DecryptionActionTypeDecrypt   DecryptionActionTypeEnum = "DECRYPT"
)

var mappingDecryptionActionTypeEnum = map[string]DecryptionActionTypeEnum{
	"NO_DECRYPT": DecryptionActionTypeNoDecrypt,
	"DECRYPT":    DecryptionActionTypeDecrypt,
}

var mappingDecryptionActionTypeEnumLowerCase = map[string]DecryptionActionTypeEnum{
	"no_decrypt": DecryptionActionTypeNoDecrypt,
	"decrypt":    DecryptionActionTypeDecrypt,
}

// GetDecryptionActionTypeEnumValues Enumerates the set of values for DecryptionActionTypeEnum
func GetDecryptionActionTypeEnumValues() []DecryptionActionTypeEnum {
	values := make([]DecryptionActionTypeEnum, 0)
	for _, v := range mappingDecryptionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDecryptionActionTypeEnumStringValues Enumerates the set of values in String for DecryptionActionTypeEnum
func GetDecryptionActionTypeEnumStringValues() []string {
	return []string{
		"NO_DECRYPT",
		"DECRYPT",
	}
}

// GetMappingDecryptionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDecryptionActionTypeEnum(val string) (DecryptionActionTypeEnum, bool) {
	enum, ok := mappingDecryptionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
