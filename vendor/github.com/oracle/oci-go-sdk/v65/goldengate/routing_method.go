// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// RoutingMethodEnum Enum with underlying type: string
type RoutingMethodEnum string

// Set of constants representing the allowable values for RoutingMethodEnum
const (
	RoutingMethodSharedServiceEndpoint    RoutingMethodEnum = "SHARED_SERVICE_ENDPOINT"
	RoutingMethodSharedDeploymentEndpoint RoutingMethodEnum = "SHARED_DEPLOYMENT_ENDPOINT"
	RoutingMethodDedicatedEndpoint        RoutingMethodEnum = "DEDICATED_ENDPOINT"
)

var mappingRoutingMethodEnum = map[string]RoutingMethodEnum{
	"SHARED_SERVICE_ENDPOINT":    RoutingMethodSharedServiceEndpoint,
	"SHARED_DEPLOYMENT_ENDPOINT": RoutingMethodSharedDeploymentEndpoint,
	"DEDICATED_ENDPOINT":         RoutingMethodDedicatedEndpoint,
}

var mappingRoutingMethodEnumLowerCase = map[string]RoutingMethodEnum{
	"shared_service_endpoint":    RoutingMethodSharedServiceEndpoint,
	"shared_deployment_endpoint": RoutingMethodSharedDeploymentEndpoint,
	"dedicated_endpoint":         RoutingMethodDedicatedEndpoint,
}

// GetRoutingMethodEnumValues Enumerates the set of values for RoutingMethodEnum
func GetRoutingMethodEnumValues() []RoutingMethodEnum {
	values := make([]RoutingMethodEnum, 0)
	for _, v := range mappingRoutingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetRoutingMethodEnumStringValues Enumerates the set of values in String for RoutingMethodEnum
func GetRoutingMethodEnumStringValues() []string {
	return []string{
		"SHARED_SERVICE_ENDPOINT",
		"SHARED_DEPLOYMENT_ENDPOINT",
		"DEDICATED_ENDPOINT",
	}
}

// GetMappingRoutingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRoutingMethodEnum(val string) (RoutingMethodEnum, bool) {
	enum, ok := mappingRoutingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
