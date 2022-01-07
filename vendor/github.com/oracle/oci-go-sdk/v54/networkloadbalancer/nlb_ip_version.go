// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

// NlbIpVersionEnum Enum with underlying type: string
type NlbIpVersionEnum string

// Set of constants representing the allowable values for NlbIpVersionEnum
const (
	NlbIpVersionIpv4        NlbIpVersionEnum = "IPV4"
	NlbIpVersionIpv4AndIpv6 NlbIpVersionEnum = "IPV4_AND_IPV6"
)

var mappingNlbIpVersion = map[string]NlbIpVersionEnum{
	"IPV4":          NlbIpVersionIpv4,
	"IPV4_AND_IPV6": NlbIpVersionIpv4AndIpv6,
}

// GetNlbIpVersionEnumValues Enumerates the set of values for NlbIpVersionEnum
func GetNlbIpVersionEnumValues() []NlbIpVersionEnum {
	values := make([]NlbIpVersionEnum, 0)
	for _, v := range mappingNlbIpVersion {
		values = append(values, v)
	}
	return values
}
