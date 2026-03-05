// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs. For more information, see Overview of Network Firewall (https://docs.oracle.com/iaas/Content/network-firewall/overview.htm).
//

package networkfirewall

import (
	"strings"
)

// HealthStatusEnum Enum with underlying type: string
type HealthStatusEnum string

// Set of constants representing the allowable values for HealthStatusEnum
const (
	HealthStatusCritical HealthStatusEnum = "CRITICAL"
	HealthStatusWarning  HealthStatusEnum = "WARNING"
	HealthStatusOk       HealthStatusEnum = "OK"
	HealthStatusUnknown  HealthStatusEnum = "UNKNOWN"
)

var mappingHealthStatusEnum = map[string]HealthStatusEnum{
	"CRITICAL": HealthStatusCritical,
	"WARNING":  HealthStatusWarning,
	"OK":       HealthStatusOk,
	"UNKNOWN":  HealthStatusUnknown,
}

var mappingHealthStatusEnumLowerCase = map[string]HealthStatusEnum{
	"critical": HealthStatusCritical,
	"warning":  HealthStatusWarning,
	"ok":       HealthStatusOk,
	"unknown":  HealthStatusUnknown,
}

// GetHealthStatusEnumValues Enumerates the set of values for HealthStatusEnum
func GetHealthStatusEnumValues() []HealthStatusEnum {
	values := make([]HealthStatusEnum, 0)
	for _, v := range mappingHealthStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthStatusEnumStringValues Enumerates the set of values in String for HealthStatusEnum
func GetHealthStatusEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"WARNING",
		"OK",
		"UNKNOWN",
	}
}

// GetMappingHealthStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthStatusEnum(val string) (HealthStatusEnum, bool) {
	enum, ok := mappingHealthStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
