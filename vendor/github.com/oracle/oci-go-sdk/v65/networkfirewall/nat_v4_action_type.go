// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// NatV4ActionTypeEnum Enum with underlying type: string
type NatV4ActionTypeEnum string

// Set of constants representing the allowable values for NatV4ActionTypeEnum
const (
	NatV4ActionTypeDippSrcNat NatV4ActionTypeEnum = "DIPP_SRC_NAT"
)

var mappingNatV4ActionTypeEnum = map[string]NatV4ActionTypeEnum{
	"DIPP_SRC_NAT": NatV4ActionTypeDippSrcNat,
}

var mappingNatV4ActionTypeEnumLowerCase = map[string]NatV4ActionTypeEnum{
	"dipp_src_nat": NatV4ActionTypeDippSrcNat,
}

// GetNatV4ActionTypeEnumValues Enumerates the set of values for NatV4ActionTypeEnum
func GetNatV4ActionTypeEnumValues() []NatV4ActionTypeEnum {
	values := make([]NatV4ActionTypeEnum, 0)
	for _, v := range mappingNatV4ActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNatV4ActionTypeEnumStringValues Enumerates the set of values in String for NatV4ActionTypeEnum
func GetNatV4ActionTypeEnumStringValues() []string {
	return []string{
		"DIPP_SRC_NAT",
	}
}

// GetMappingNatV4ActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNatV4ActionTypeEnum(val string) (NatV4ActionTypeEnum, bool) {
	enum, ok := mappingNatV4ActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
