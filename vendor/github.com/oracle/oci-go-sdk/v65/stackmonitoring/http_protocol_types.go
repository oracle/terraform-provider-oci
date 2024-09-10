// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// HttpProtocolTypesEnum Enum with underlying type: string
type HttpProtocolTypesEnum string

// Set of constants representing the allowable values for HttpProtocolTypesEnum
const (
	HttpProtocolTypesHttp  HttpProtocolTypesEnum = "HTTP"
	HttpProtocolTypesHttps HttpProtocolTypesEnum = "HTTPS"
)

var mappingHttpProtocolTypesEnum = map[string]HttpProtocolTypesEnum{
	"HTTP":  HttpProtocolTypesHttp,
	"HTTPS": HttpProtocolTypesHttps,
}

var mappingHttpProtocolTypesEnumLowerCase = map[string]HttpProtocolTypesEnum{
	"http":  HttpProtocolTypesHttp,
	"https": HttpProtocolTypesHttps,
}

// GetHttpProtocolTypesEnumValues Enumerates the set of values for HttpProtocolTypesEnum
func GetHttpProtocolTypesEnumValues() []HttpProtocolTypesEnum {
	values := make([]HttpProtocolTypesEnum, 0)
	for _, v := range mappingHttpProtocolTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpProtocolTypesEnumStringValues Enumerates the set of values in String for HttpProtocolTypesEnum
func GetHttpProtocolTypesEnumStringValues() []string {
	return []string{
		"HTTP",
		"HTTPS",
	}
}

// GetMappingHttpProtocolTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpProtocolTypesEnum(val string) (HttpProtocolTypesEnum, bool) {
	enum, ok := mappingHttpProtocolTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
