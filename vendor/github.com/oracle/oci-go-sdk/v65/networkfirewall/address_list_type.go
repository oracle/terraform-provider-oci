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

// AddressListTypeEnum Enum with underlying type: string
type AddressListTypeEnum string

// Set of constants representing the allowable values for AddressListTypeEnum
const (
	AddressListTypeFqdn AddressListTypeEnum = "FQDN"
	AddressListTypeIp   AddressListTypeEnum = "IP"
)

var mappingAddressListTypeEnum = map[string]AddressListTypeEnum{
	"FQDN": AddressListTypeFqdn,
	"IP":   AddressListTypeIp,
}

var mappingAddressListTypeEnumLowerCase = map[string]AddressListTypeEnum{
	"fqdn": AddressListTypeFqdn,
	"ip":   AddressListTypeIp,
}

// GetAddressListTypeEnumValues Enumerates the set of values for AddressListTypeEnum
func GetAddressListTypeEnumValues() []AddressListTypeEnum {
	values := make([]AddressListTypeEnum, 0)
	for _, v := range mappingAddressListTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddressListTypeEnumStringValues Enumerates the set of values in String for AddressListTypeEnum
func GetAddressListTypeEnumStringValues() []string {
	return []string{
		"FQDN",
		"IP",
	}
}

// GetMappingAddressListTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddressListTypeEnum(val string) (AddressListTypeEnum, bool) {
	enum, ok := mappingAddressListTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
