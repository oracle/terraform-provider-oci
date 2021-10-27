// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

// ListenerProtocolsEnum Enum with underlying type: string
type ListenerProtocolsEnum string

// Set of constants representing the allowable values for ListenerProtocolsEnum
const (
	ListenerProtocolsAny ListenerProtocolsEnum = "ANY"
	ListenerProtocolsTcp ListenerProtocolsEnum = "TCP"
	ListenerProtocolsUdp ListenerProtocolsEnum = "UDP"
)

var mappingListenerProtocols = map[string]ListenerProtocolsEnum{
	"ANY": ListenerProtocolsAny,
	"TCP": ListenerProtocolsTcp,
	"UDP": ListenerProtocolsUdp,
}

// GetListenerProtocolsEnumValues Enumerates the set of values for ListenerProtocolsEnum
func GetListenerProtocolsEnumValues() []ListenerProtocolsEnum {
	values := make([]ListenerProtocolsEnum, 0)
	for _, v := range mappingListenerProtocols {
		values = append(values, v)
	}
	return values
}
