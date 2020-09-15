// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

// PingProbeProtocolEnum Enum with underlying type: string
type PingProbeProtocolEnum string

// Set of constants representing the allowable values for PingProbeProtocolEnum
const (
	PingProbeProtocolIcmp PingProbeProtocolEnum = "ICMP"
	PingProbeProtocolTcp  PingProbeProtocolEnum = "TCP"
)

var mappingPingProbeProtocol = map[string]PingProbeProtocolEnum{
	"ICMP": PingProbeProtocolIcmp,
	"TCP":  PingProbeProtocolTcp,
}

// GetPingProbeProtocolEnumValues Enumerates the set of values for PingProbeProtocolEnum
func GetPingProbeProtocolEnumValues() []PingProbeProtocolEnum {
	values := make([]PingProbeProtocolEnum, 0)
	for _, v := range mappingPingProbeProtocol {
		values = append(values, v)
	}
	return values
}
