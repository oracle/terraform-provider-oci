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

// HealthCheckProtocolsEnum Enum with underlying type: string
type HealthCheckProtocolsEnum string

// Set of constants representing the allowable values for HealthCheckProtocolsEnum
const (
	HealthCheckProtocolsHttp  HealthCheckProtocolsEnum = "HTTP"
	HealthCheckProtocolsHttps HealthCheckProtocolsEnum = "HTTPS"
	HealthCheckProtocolsTcp   HealthCheckProtocolsEnum = "TCP"
	HealthCheckProtocolsUdp   HealthCheckProtocolsEnum = "UDP"
	HealthCheckProtocolsDns   HealthCheckProtocolsEnum = "DNS"
)

var mappingHealthCheckProtocolsEnum = map[string]HealthCheckProtocolsEnum{
	"HTTP":  HealthCheckProtocolsHttp,
	"HTTPS": HealthCheckProtocolsHttps,
	"TCP":   HealthCheckProtocolsTcp,
	"UDP":   HealthCheckProtocolsUdp,
	"DNS":   HealthCheckProtocolsDns,
}

var mappingHealthCheckProtocolsEnumLowerCase = map[string]HealthCheckProtocolsEnum{
	"http":  HealthCheckProtocolsHttp,
	"https": HealthCheckProtocolsHttps,
	"tcp":   HealthCheckProtocolsTcp,
	"udp":   HealthCheckProtocolsUdp,
	"dns":   HealthCheckProtocolsDns,
}

// GetHealthCheckProtocolsEnumValues Enumerates the set of values for HealthCheckProtocolsEnum
func GetHealthCheckProtocolsEnumValues() []HealthCheckProtocolsEnum {
	values := make([]HealthCheckProtocolsEnum, 0)
	for _, v := range mappingHealthCheckProtocolsEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckProtocolsEnumStringValues Enumerates the set of values in String for HealthCheckProtocolsEnum
func GetHealthCheckProtocolsEnumStringValues() []string {
	return []string{
		"HTTP",
		"HTTPS",
		"TCP",
		"UDP",
		"DNS",
	}
}

// GetMappingHealthCheckProtocolsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckProtocolsEnum(val string) (HealthCheckProtocolsEnum, bool) {
	enum, ok := mappingHealthCheckProtocolsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
