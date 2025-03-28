// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"strings"
)

// BackendTypeEnum Enum with underlying type: string
type BackendTypeEnum string

// Set of constants representing the allowable values for BackendTypeEnum
const (
	BackendTypeLoadBalancer BackendTypeEnum = "LOAD_BALANCER"
)

var mappingBackendTypeEnum = map[string]BackendTypeEnum{
	"LOAD_BALANCER": BackendTypeLoadBalancer,
}

var mappingBackendTypeEnumLowerCase = map[string]BackendTypeEnum{
	"load_balancer": BackendTypeLoadBalancer,
}

// GetBackendTypeEnumValues Enumerates the set of values for BackendTypeEnum
func GetBackendTypeEnumValues() []BackendTypeEnum {
	values := make([]BackendTypeEnum, 0)
	for _, v := range mappingBackendTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBackendTypeEnumStringValues Enumerates the set of values in String for BackendTypeEnum
func GetBackendTypeEnumStringValues() []string {
	return []string{
		"LOAD_BALANCER",
	}
}

// GetMappingBackendTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackendTypeEnum(val string) (BackendTypeEnum, bool) {
	enum, ok := mappingBackendTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
