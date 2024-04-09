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

// DnsHealthCheckRCodesEnum Enum with underlying type: string
type DnsHealthCheckRCodesEnum string

// Set of constants representing the allowable values for DnsHealthCheckRCodesEnum
const (
	DnsHealthCheckRCodesNoerror  DnsHealthCheckRCodesEnum = "NOERROR"
	DnsHealthCheckRCodesServfail DnsHealthCheckRCodesEnum = "SERVFAIL"
	DnsHealthCheckRCodesNxdomain DnsHealthCheckRCodesEnum = "NXDOMAIN"
	DnsHealthCheckRCodesRefused  DnsHealthCheckRCodesEnum = "REFUSED"
)

var mappingDnsHealthCheckRCodesEnum = map[string]DnsHealthCheckRCodesEnum{
	"NOERROR":  DnsHealthCheckRCodesNoerror,
	"SERVFAIL": DnsHealthCheckRCodesServfail,
	"NXDOMAIN": DnsHealthCheckRCodesNxdomain,
	"REFUSED":  DnsHealthCheckRCodesRefused,
}

var mappingDnsHealthCheckRCodesEnumLowerCase = map[string]DnsHealthCheckRCodesEnum{
	"noerror":  DnsHealthCheckRCodesNoerror,
	"servfail": DnsHealthCheckRCodesServfail,
	"nxdomain": DnsHealthCheckRCodesNxdomain,
	"refused":  DnsHealthCheckRCodesRefused,
}

// GetDnsHealthCheckRCodesEnumValues Enumerates the set of values for DnsHealthCheckRCodesEnum
func GetDnsHealthCheckRCodesEnumValues() []DnsHealthCheckRCodesEnum {
	values := make([]DnsHealthCheckRCodesEnum, 0)
	for _, v := range mappingDnsHealthCheckRCodesEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsHealthCheckRCodesEnumStringValues Enumerates the set of values in String for DnsHealthCheckRCodesEnum
func GetDnsHealthCheckRCodesEnumStringValues() []string {
	return []string{
		"NOERROR",
		"SERVFAIL",
		"NXDOMAIN",
		"REFUSED",
	}
}

// GetMappingDnsHealthCheckRCodesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsHealthCheckRCodesEnum(val string) (DnsHealthCheckRCodesEnum, bool) {
	enum, ok := mappingDnsHealthCheckRCodesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
