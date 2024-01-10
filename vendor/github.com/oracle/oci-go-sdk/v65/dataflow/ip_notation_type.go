// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// IpNotationTypeEnum Enum with underlying type: string
type IpNotationTypeEnum string

// Set of constants representing the allowable values for IpNotationTypeEnum
const (
	IpNotationTypeIpAddress IpNotationTypeEnum = "IP_ADDRESS"
	IpNotationTypeCidr      IpNotationTypeEnum = "CIDR"
	IpNotationTypeVcn       IpNotationTypeEnum = "VCN"
	IpNotationTypeVcnOcid   IpNotationTypeEnum = "VCN_OCID"
)

var mappingIpNotationTypeEnum = map[string]IpNotationTypeEnum{
	"IP_ADDRESS": IpNotationTypeIpAddress,
	"CIDR":       IpNotationTypeCidr,
	"VCN":        IpNotationTypeVcn,
	"VCN_OCID":   IpNotationTypeVcnOcid,
}

var mappingIpNotationTypeEnumLowerCase = map[string]IpNotationTypeEnum{
	"ip_address": IpNotationTypeIpAddress,
	"cidr":       IpNotationTypeCidr,
	"vcn":        IpNotationTypeVcn,
	"vcn_ocid":   IpNotationTypeVcnOcid,
}

// GetIpNotationTypeEnumValues Enumerates the set of values for IpNotationTypeEnum
func GetIpNotationTypeEnumValues() []IpNotationTypeEnum {
	values := make([]IpNotationTypeEnum, 0)
	for _, v := range mappingIpNotationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIpNotationTypeEnumStringValues Enumerates the set of values in String for IpNotationTypeEnum
func GetIpNotationTypeEnumStringValues() []string {
	return []string{
		"IP_ADDRESS",
		"CIDR",
		"VCN",
		"VCN_OCID",
	}
}

// GetMappingIpNotationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpNotationTypeEnum(val string) (IpNotationTypeEnum, bool) {
	enum, ok := mappingIpNotationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
