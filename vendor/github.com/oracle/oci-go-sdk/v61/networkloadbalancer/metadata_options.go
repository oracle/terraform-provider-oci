// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"strings"
)

// MetadataOptionsEnum Enum with underlying type: string
type MetadataOptionsEnum string

// Set of constants representing the allowable values for MetadataOptionsEnum
const (
	MetadataOptionsVcnId      MetadataOptionsEnum = "VCN_ID"
	MetadataOptionsSgwPeNatIp MetadataOptionsEnum = "SGW_PE_NAT_IP"
	MetadataOptionsVcnOcid    MetadataOptionsEnum = "VCN_OCID"
	MetadataOptionsPeOcid     MetadataOptionsEnum = "PE_OCID"
)

var mappingMetadataOptionsEnum = map[string]MetadataOptionsEnum{
	"VCN_ID":        MetadataOptionsVcnId,
	"SGW_PE_NAT_IP": MetadataOptionsSgwPeNatIp,
	"VCN_OCID":      MetadataOptionsVcnOcid,
	"PE_OCID":       MetadataOptionsPeOcid,
}

var mappingMetadataOptionsEnumLowerCase = map[string]MetadataOptionsEnum{
	"vcn_id":        MetadataOptionsVcnId,
	"sgw_pe_nat_ip": MetadataOptionsSgwPeNatIp,
	"vcn_ocid":      MetadataOptionsVcnOcid,
	"pe_ocid":       MetadataOptionsPeOcid,
}

// GetMetadataOptionsEnumValues Enumerates the set of values for MetadataOptionsEnum
func GetMetadataOptionsEnumValues() []MetadataOptionsEnum {
	values := make([]MetadataOptionsEnum, 0)
	for _, v := range mappingMetadataOptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetMetadataOptionsEnumStringValues Enumerates the set of values in String for MetadataOptionsEnum
func GetMetadataOptionsEnumStringValues() []string {
	return []string{
		"VCN_ID",
		"SGW_PE_NAT_IP",
		"VCN_OCID",
		"PE_OCID",
	}
}

// GetMappingMetadataOptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetadataOptionsEnum(val string) (MetadataOptionsEnum, bool) {
	enum, ok := mappingMetadataOptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
