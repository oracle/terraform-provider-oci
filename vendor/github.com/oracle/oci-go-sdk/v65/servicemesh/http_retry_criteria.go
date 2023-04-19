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

// HttpRetryCriteriaEnum Enum with underlying type: string
type HttpRetryCriteriaEnum string

// Set of constants representing the allowable values for HttpRetryCriteriaEnum
const (
	HttpRetryCriteriaReset          HttpRetryCriteriaEnum = "RESET"
	HttpRetryCriteriaConnectFailure HttpRetryCriteriaEnum = "CONNECT_FAILURE"
	HttpRetryCriteriaRefusedStream  HttpRetryCriteriaEnum = "REFUSED_STREAM"
)

var mappingHttpRetryCriteriaEnum = map[string]HttpRetryCriteriaEnum{
	"RESET":           HttpRetryCriteriaReset,
	"CONNECT_FAILURE": HttpRetryCriteriaConnectFailure,
	"REFUSED_STREAM":  HttpRetryCriteriaRefusedStream,
}

var mappingHttpRetryCriteriaEnumLowerCase = map[string]HttpRetryCriteriaEnum{
	"reset":           HttpRetryCriteriaReset,
	"connect_failure": HttpRetryCriteriaConnectFailure,
	"refused_stream":  HttpRetryCriteriaRefusedStream,
}

// GetHttpRetryCriteriaEnumValues Enumerates the set of values for HttpRetryCriteriaEnum
func GetHttpRetryCriteriaEnumValues() []HttpRetryCriteriaEnum {
	values := make([]HttpRetryCriteriaEnum, 0)
	for _, v := range mappingHttpRetryCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpRetryCriteriaEnumStringValues Enumerates the set of values in String for HttpRetryCriteriaEnum
func GetHttpRetryCriteriaEnumStringValues() []string {
	return []string{
		"RESET",
		"CONNECT_FAILURE",
		"REFUSED_STREAM",
	}
}

// GetMappingHttpRetryCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpRetryCriteriaEnum(val string) (HttpRetryCriteriaEnum, bool) {
	enum, ok := mappingHttpRetryCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
