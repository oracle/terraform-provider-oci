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

// HcsInfraIpVersionEnum Enum with underlying type: string
type HcsInfraIpVersionEnum string

// Set of constants representing the allowable values for HcsInfraIpVersionEnum
const (
	HcsInfraIpVersionIpv4        HcsInfraIpVersionEnum = "IPV4"
	HcsInfraIpVersionIpv4AndIpv6 HcsInfraIpVersionEnum = "IPV4_AND_IPV6"
	HcsInfraIpVersionIpv6        HcsInfraIpVersionEnum = "IPV6"
)

var mappingHcsInfraIpVersionEnum = map[string]HcsInfraIpVersionEnum{
	"IPV4":          HcsInfraIpVersionIpv4,
	"IPV4_AND_IPV6": HcsInfraIpVersionIpv4AndIpv6,
	"IPV6":          HcsInfraIpVersionIpv6,
}

var mappingHcsInfraIpVersionEnumLowerCase = map[string]HcsInfraIpVersionEnum{
	"ipv4":          HcsInfraIpVersionIpv4,
	"ipv4_and_ipv6": HcsInfraIpVersionIpv4AndIpv6,
	"ipv6":          HcsInfraIpVersionIpv6,
}

// GetHcsInfraIpVersionEnumValues Enumerates the set of values for HcsInfraIpVersionEnum
func GetHcsInfraIpVersionEnumValues() []HcsInfraIpVersionEnum {
	values := make([]HcsInfraIpVersionEnum, 0)
	for _, v := range mappingHcsInfraIpVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetHcsInfraIpVersionEnumStringValues Enumerates the set of values in String for HcsInfraIpVersionEnum
func GetHcsInfraIpVersionEnumStringValues() []string {
	return []string{
		"IPV4",
		"IPV4_AND_IPV6",
		"IPV6",
	}
}

// GetMappingHcsInfraIpVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHcsInfraIpVersionEnum(val string) (HcsInfraIpVersionEnum, bool) {
	enum, ok := mappingHcsInfraIpVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
