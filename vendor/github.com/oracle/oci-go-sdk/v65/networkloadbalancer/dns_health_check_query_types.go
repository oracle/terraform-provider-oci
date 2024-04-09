// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DnsHealthCheckQueryTypesEnum Enum with underlying type: string
type DnsHealthCheckQueryTypesEnum string

// Set of constants representing the allowable values for DnsHealthCheckQueryTypesEnum
const (
	DnsHealthCheckQueryTypesA    DnsHealthCheckQueryTypesEnum = "A"
	DnsHealthCheckQueryTypesTxt  DnsHealthCheckQueryTypesEnum = "TXT"
	DnsHealthCheckQueryTypesAaaa DnsHealthCheckQueryTypesEnum = "AAAA"
)

var mappingDnsHealthCheckQueryTypesEnum = map[string]DnsHealthCheckQueryTypesEnum{
	"A":    DnsHealthCheckQueryTypesA,
	"TXT":  DnsHealthCheckQueryTypesTxt,
	"AAAA": DnsHealthCheckQueryTypesAaaa,
}

var mappingDnsHealthCheckQueryTypesEnumLowerCase = map[string]DnsHealthCheckQueryTypesEnum{
	"a":    DnsHealthCheckQueryTypesA,
	"txt":  DnsHealthCheckQueryTypesTxt,
	"aaaa": DnsHealthCheckQueryTypesAaaa,
}

// GetDnsHealthCheckQueryTypesEnumValues Enumerates the set of values for DnsHealthCheckQueryTypesEnum
func GetDnsHealthCheckQueryTypesEnumValues() []DnsHealthCheckQueryTypesEnum {
	values := make([]DnsHealthCheckQueryTypesEnum, 0)
	for _, v := range mappingDnsHealthCheckQueryTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsHealthCheckQueryTypesEnumStringValues Enumerates the set of values in String for DnsHealthCheckQueryTypesEnum
func GetDnsHealthCheckQueryTypesEnumStringValues() []string {
	return []string{
		"A",
		"TXT",
		"AAAA",
	}
}

// GetMappingDnsHealthCheckQueryTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsHealthCheckQueryTypesEnum(val string) (DnsHealthCheckQueryTypesEnum, bool) {
	enum, ok := mappingDnsHealthCheckQueryTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
