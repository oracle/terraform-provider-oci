// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// HttpMethodEnum Enum with underlying type: string
type HttpMethodEnum string

// Set of constants representing the allowable values for HttpMethodEnum
const (
	HttpMethodGet    HttpMethodEnum = "GET"
	HttpMethodPut    HttpMethodEnum = "PUT"
	HttpMethodPost   HttpMethodEnum = "POST"
	HttpMethodDelete HttpMethodEnum = "DELETE"
	HttpMethodHead   HttpMethodEnum = "HEAD"
)

var mappingHttpMethodEnum = map[string]HttpMethodEnum{
	"GET":    HttpMethodGet,
	"PUT":    HttpMethodPut,
	"POST":   HttpMethodPost,
	"DELETE": HttpMethodDelete,
	"HEAD":   HttpMethodHead,
}

var mappingHttpMethodEnumLowerCase = map[string]HttpMethodEnum{
	"get":    HttpMethodGet,
	"put":    HttpMethodPut,
	"post":   HttpMethodPost,
	"delete": HttpMethodDelete,
	"head":   HttpMethodHead,
}

// GetHttpMethodEnumValues Enumerates the set of values for HttpMethodEnum
func GetHttpMethodEnumValues() []HttpMethodEnum {
	values := make([]HttpMethodEnum, 0)
	for _, v := range mappingHttpMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpMethodEnumStringValues Enumerates the set of values in String for HttpMethodEnum
func GetHttpMethodEnumStringValues() []string {
	return []string{
		"GET",
		"PUT",
		"POST",
		"DELETE",
		"HEAD",
	}
}

// GetMappingHttpMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpMethodEnum(val string) (HttpMethodEnum, bool) {
	enum, ok := mappingHttpMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
