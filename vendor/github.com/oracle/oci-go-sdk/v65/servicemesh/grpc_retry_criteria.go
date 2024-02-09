// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GrpcRetryCriteriaEnum Enum with underlying type: string
type GrpcRetryCriteriaEnum string

// Set of constants representing the allowable values for GrpcRetryCriteriaEnum
const (
	GrpcRetryCriteriaCancelled         GrpcRetryCriteriaEnum = "CANCELLED"
	GrpcRetryCriteriaDeadlineExceeded  GrpcRetryCriteriaEnum = "DEADLINE_EXCEEDED"
	GrpcRetryCriteriaInternal          GrpcRetryCriteriaEnum = "INTERNAL"
	GrpcRetryCriteriaResourceExhausted GrpcRetryCriteriaEnum = "RESOURCE_EXHAUSTED"
	GrpcRetryCriteriaUnavailable       GrpcRetryCriteriaEnum = "UNAVAILABLE"
)

var mappingGrpcRetryCriteriaEnum = map[string]GrpcRetryCriteriaEnum{
	"CANCELLED":          GrpcRetryCriteriaCancelled,
	"DEADLINE_EXCEEDED":  GrpcRetryCriteriaDeadlineExceeded,
	"INTERNAL":           GrpcRetryCriteriaInternal,
	"RESOURCE_EXHAUSTED": GrpcRetryCriteriaResourceExhausted,
	"UNAVAILABLE":        GrpcRetryCriteriaUnavailable,
}

var mappingGrpcRetryCriteriaEnumLowerCase = map[string]GrpcRetryCriteriaEnum{
	"cancelled":          GrpcRetryCriteriaCancelled,
	"deadline_exceeded":  GrpcRetryCriteriaDeadlineExceeded,
	"internal":           GrpcRetryCriteriaInternal,
	"resource_exhausted": GrpcRetryCriteriaResourceExhausted,
	"unavailable":        GrpcRetryCriteriaUnavailable,
}

// GetGrpcRetryCriteriaEnumValues Enumerates the set of values for GrpcRetryCriteriaEnum
func GetGrpcRetryCriteriaEnumValues() []GrpcRetryCriteriaEnum {
	values := make([]GrpcRetryCriteriaEnum, 0)
	for _, v := range mappingGrpcRetryCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetGrpcRetryCriteriaEnumStringValues Enumerates the set of values in String for GrpcRetryCriteriaEnum
func GetGrpcRetryCriteriaEnumStringValues() []string {
	return []string{
		"CANCELLED",
		"DEADLINE_EXCEEDED",
		"INTERNAL",
		"RESOURCE_EXHAUSTED",
		"UNAVAILABLE",
	}
}

// GetMappingGrpcRetryCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrpcRetryCriteriaEnum(val string) (GrpcRetryCriteriaEnum, bool) {
	enum, ok := mappingGrpcRetryCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
