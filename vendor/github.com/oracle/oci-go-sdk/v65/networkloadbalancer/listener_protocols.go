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

// ListenerProtocolsEnum Enum with underlying type: string
type ListenerProtocolsEnum string

// Set of constants representing the allowable values for ListenerProtocolsEnum
const (
	ListenerProtocolsAny       ListenerProtocolsEnum = "ANY"
	ListenerProtocolsTcp       ListenerProtocolsEnum = "TCP"
	ListenerProtocolsUdp       ListenerProtocolsEnum = "UDP"
	ListenerProtocolsTcpAndUdp ListenerProtocolsEnum = "TCP_AND_UDP"
)

var mappingListenerProtocolsEnum = map[string]ListenerProtocolsEnum{
	"ANY":         ListenerProtocolsAny,
	"TCP":         ListenerProtocolsTcp,
	"UDP":         ListenerProtocolsUdp,
	"TCP_AND_UDP": ListenerProtocolsTcpAndUdp,
}

var mappingListenerProtocolsEnumLowerCase = map[string]ListenerProtocolsEnum{
	"any":         ListenerProtocolsAny,
	"tcp":         ListenerProtocolsTcp,
	"udp":         ListenerProtocolsUdp,
	"tcp_and_udp": ListenerProtocolsTcpAndUdp,
}

// GetListenerProtocolsEnumValues Enumerates the set of values for ListenerProtocolsEnum
func GetListenerProtocolsEnumValues() []ListenerProtocolsEnum {
	values := make([]ListenerProtocolsEnum, 0)
	for _, v := range mappingListenerProtocolsEnum {
		values = append(values, v)
	}
	return values
}

// GetListenerProtocolsEnumStringValues Enumerates the set of values in String for ListenerProtocolsEnum
func GetListenerProtocolsEnumStringValues() []string {
	return []string{
		"ANY",
		"TCP",
		"UDP",
		"TCP_AND_UDP",
	}
}

// GetMappingListenerProtocolsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListenerProtocolsEnum(val string) (ListenerProtocolsEnum, bool) {
	enum, ok := mappingListenerProtocolsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
