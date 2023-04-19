// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"strings"
)

// GrpcRetryCritieriaEnum Enum with underlying type: string
type GrpcRetryCritieriaEnum string

// Set of constants representing the allowable values for GrpcRetryCritieriaEnum
const (
	GrpcRetryCritieriaCancelled         GrpcRetryCritieriaEnum = "CANCELLED"
	GrpcRetryCritieriaDeadlineExceeded  GrpcRetryCritieriaEnum = "DEADLINE_EXCEEDED"
	GrpcRetryCritieriaInternal          GrpcRetryCritieriaEnum = "INTERNAL"
	GrpcRetryCritieriaResourceExhausted GrpcRetryCritieriaEnum = "RESOURCE_EXHAUSTED"
	GrpcRetryCritieriaUnavailable       GrpcRetryCritieriaEnum = "UNAVAILABLE"
)

var mappingGrpcRetryCritieriaEnum = map[string]GrpcRetryCritieriaEnum{
	"CANCELLED":          GrpcRetryCritieriaCancelled,
	"DEADLINE_EXCEEDED":  GrpcRetryCritieriaDeadlineExceeded,
	"INTERNAL":           GrpcRetryCritieriaInternal,
	"RESOURCE_EXHAUSTED": GrpcRetryCritieriaResourceExhausted,
	"UNAVAILABLE":        GrpcRetryCritieriaUnavailable,
}

var mappingGrpcRetryCritieriaEnumLowerCase = map[string]GrpcRetryCritieriaEnum{
	"cancelled":          GrpcRetryCritieriaCancelled,
	"deadline_exceeded":  GrpcRetryCritieriaDeadlineExceeded,
	"internal":           GrpcRetryCritieriaInternal,
	"resource_exhausted": GrpcRetryCritieriaResourceExhausted,
	"unavailable":        GrpcRetryCritieriaUnavailable,
}

// GetGrpcRetryCritieriaEnumValues Enumerates the set of values for GrpcRetryCritieriaEnum
func GetGrpcRetryCritieriaEnumValues() []GrpcRetryCritieriaEnum {
	values := make([]GrpcRetryCritieriaEnum, 0)
	for _, v := range mappingGrpcRetryCritieriaEnum {
		values = append(values, v)
	}
	return values
}

// GetGrpcRetryCritieriaEnumStringValues Enumerates the set of values in String for GrpcRetryCritieriaEnum
func GetGrpcRetryCritieriaEnumStringValues() []string {
	return []string{
		"CANCELLED",
		"DEADLINE_EXCEEDED",
		"INTERNAL",
		"RESOURCE_EXHAUSTED",
		"UNAVAILABLE",
	}
}

// GetMappingGrpcRetryCritieriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrpcRetryCritieriaEnum(val string) (GrpcRetryCritieriaEnum, bool) {
	enum, ok := mappingGrpcRetryCritieriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
