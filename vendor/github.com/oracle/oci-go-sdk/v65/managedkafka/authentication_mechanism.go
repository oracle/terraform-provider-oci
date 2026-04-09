// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"strings"
)

// AuthenticationMechanismEnum Enum with underlying type: string
type AuthenticationMechanismEnum string

// Set of constants representing the allowable values for AuthenticationMechanismEnum
const (
	AuthenticationMechanismSasl AuthenticationMechanismEnum = "SASL"
	AuthenticationMechanismMtls AuthenticationMechanismEnum = "MTLS"
)

var mappingAuthenticationMechanismEnum = map[string]AuthenticationMechanismEnum{
	"SASL": AuthenticationMechanismSasl,
	"MTLS": AuthenticationMechanismMtls,
}

var mappingAuthenticationMechanismEnumLowerCase = map[string]AuthenticationMechanismEnum{
	"sasl": AuthenticationMechanismSasl,
	"mtls": AuthenticationMechanismMtls,
}

// GetAuthenticationMechanismEnumValues Enumerates the set of values for AuthenticationMechanismEnum
func GetAuthenticationMechanismEnumValues() []AuthenticationMechanismEnum {
	values := make([]AuthenticationMechanismEnum, 0)
	for _, v := range mappingAuthenticationMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationMechanismEnumStringValues Enumerates the set of values in String for AuthenticationMechanismEnum
func GetAuthenticationMechanismEnumStringValues() []string {
	return []string{
		"SASL",
		"MTLS",
	}
}

// GetMappingAuthenticationMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationMechanismEnum(val string) (AuthenticationMechanismEnum, bool) {
	enum, ok := mappingAuthenticationMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
