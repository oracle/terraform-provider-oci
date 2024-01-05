// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"strings"
)

// BastionDnsProxyStatusEnum Enum with underlying type: string
type BastionDnsProxyStatusEnum string

// Set of constants representing the allowable values for BastionDnsProxyStatusEnum
const (
	BastionDnsProxyStatusDisabled BastionDnsProxyStatusEnum = "DISABLED"
	BastionDnsProxyStatusEnabled  BastionDnsProxyStatusEnum = "ENABLED"
)

var mappingBastionDnsProxyStatusEnum = map[string]BastionDnsProxyStatusEnum{
	"DISABLED": BastionDnsProxyStatusDisabled,
	"ENABLED":  BastionDnsProxyStatusEnabled,
}

var mappingBastionDnsProxyStatusEnumLowerCase = map[string]BastionDnsProxyStatusEnum{
	"disabled": BastionDnsProxyStatusDisabled,
	"enabled":  BastionDnsProxyStatusEnabled,
}

// GetBastionDnsProxyStatusEnumValues Enumerates the set of values for BastionDnsProxyStatusEnum
func GetBastionDnsProxyStatusEnumValues() []BastionDnsProxyStatusEnum {
	values := make([]BastionDnsProxyStatusEnum, 0)
	for _, v := range mappingBastionDnsProxyStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBastionDnsProxyStatusEnumStringValues Enumerates the set of values in String for BastionDnsProxyStatusEnum
func GetBastionDnsProxyStatusEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENABLED",
	}
}

// GetMappingBastionDnsProxyStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBastionDnsProxyStatusEnum(val string) (BastionDnsProxyStatusEnum, bool) {
	enum, ok := mappingBastionDnsProxyStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
