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

// IpVersionEnum Enum with underlying type: string
type IpVersionEnum string

// Set of constants representing the allowable values for IpVersionEnum
const (
	IpVersionIpv4 IpVersionEnum = "IPV4"
	IpVersionIpv6 IpVersionEnum = "IPV6"
)

var mappingIpVersionEnum = map[string]IpVersionEnum{
	"IPV4": IpVersionIpv4,
	"IPV6": IpVersionIpv6,
}

// GetIpVersionEnumValues Enumerates the set of values for IpVersionEnum
func GetIpVersionEnumValues() []IpVersionEnum {
	values := make([]IpVersionEnum, 0)
	for _, v := range mappingIpVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetIpVersionEnumStringValues Enumerates the set of values in String for IpVersionEnum
func GetIpVersionEnumStringValues() []string {
	return []string{
		"IPV4",
		"IPV6",
	}
}

// GetMappingIpVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpVersionEnum(val string) (IpVersionEnum, bool) {
	mappingIpVersionEnumIgnoreCase := make(map[string]IpVersionEnum)
	for k, v := range mappingIpVersionEnum {
		mappingIpVersionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIpVersionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
