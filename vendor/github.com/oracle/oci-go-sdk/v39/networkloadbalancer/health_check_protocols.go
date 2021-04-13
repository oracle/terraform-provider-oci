// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

// HealthCheckProtocolsEnum Enum with underlying type: string
type HealthCheckProtocolsEnum string

// Set of constants representing the allowable values for HealthCheckProtocolsEnum
const (
	HealthCheckProtocolsHttp  HealthCheckProtocolsEnum = "HTTP"
	HealthCheckProtocolsHttps HealthCheckProtocolsEnum = "HTTPS"
	HealthCheckProtocolsTcp   HealthCheckProtocolsEnum = "TCP"
	HealthCheckProtocolsUdp   HealthCheckProtocolsEnum = "UDP"
)

var mappingHealthCheckProtocols = map[string]HealthCheckProtocolsEnum{
	"HTTP":  HealthCheckProtocolsHttp,
	"HTTPS": HealthCheckProtocolsHttps,
	"TCP":   HealthCheckProtocolsTcp,
	"UDP":   HealthCheckProtocolsUdp,
}

// GetHealthCheckProtocolsEnumValues Enumerates the set of values for HealthCheckProtocolsEnum
func GetHealthCheckProtocolsEnumValues() []HealthCheckProtocolsEnum {
	values := make([]HealthCheckProtocolsEnum, 0)
	for _, v := range mappingHealthCheckProtocols {
		values = append(values, v)
	}
	return values
}
