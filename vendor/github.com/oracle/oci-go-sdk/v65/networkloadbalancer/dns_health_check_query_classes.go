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

// DnsHealthCheckQueryClassesEnum Enum with underlying type: string
type DnsHealthCheckQueryClassesEnum string

// Set of constants representing the allowable values for DnsHealthCheckQueryClassesEnum
const (
	DnsHealthCheckQueryClassesIn DnsHealthCheckQueryClassesEnum = "IN"
	DnsHealthCheckQueryClassesCh DnsHealthCheckQueryClassesEnum = "CH"
)

var mappingDnsHealthCheckQueryClassesEnum = map[string]DnsHealthCheckQueryClassesEnum{
	"IN": DnsHealthCheckQueryClassesIn,
	"CH": DnsHealthCheckQueryClassesCh,
}

var mappingDnsHealthCheckQueryClassesEnumLowerCase = map[string]DnsHealthCheckQueryClassesEnum{
	"in": DnsHealthCheckQueryClassesIn,
	"ch": DnsHealthCheckQueryClassesCh,
}

// GetDnsHealthCheckQueryClassesEnumValues Enumerates the set of values for DnsHealthCheckQueryClassesEnum
func GetDnsHealthCheckQueryClassesEnumValues() []DnsHealthCheckQueryClassesEnum {
	values := make([]DnsHealthCheckQueryClassesEnum, 0)
	for _, v := range mappingDnsHealthCheckQueryClassesEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsHealthCheckQueryClassesEnumStringValues Enumerates the set of values in String for DnsHealthCheckQueryClassesEnum
func GetDnsHealthCheckQueryClassesEnumStringValues() []string {
	return []string{
		"IN",
		"CH",
	}
}

// GetMappingDnsHealthCheckQueryClassesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsHealthCheckQueryClassesEnum(val string) (DnsHealthCheckQueryClassesEnum, bool) {
	enum, ok := mappingDnsHealthCheckQueryClassesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
