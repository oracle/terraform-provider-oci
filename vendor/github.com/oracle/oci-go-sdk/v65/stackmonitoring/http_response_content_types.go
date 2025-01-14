// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// HttpResponseContentTypesEnum Enum with underlying type: string
type HttpResponseContentTypesEnum string

// Set of constants representing the allowable values for HttpResponseContentTypesEnum
const (
	HttpResponseContentTypesTextPlain       HttpResponseContentTypesEnum = "TEXT_PLAIN"
	HttpResponseContentTypesTextHtml        HttpResponseContentTypesEnum = "TEXT_HTML"
	HttpResponseContentTypesApplicationJson HttpResponseContentTypesEnum = "APPLICATION_JSON"
	HttpResponseContentTypesApplicationXml  HttpResponseContentTypesEnum = "APPLICATION_XML"
)

var mappingHttpResponseContentTypesEnum = map[string]HttpResponseContentTypesEnum{
	"TEXT_PLAIN":       HttpResponseContentTypesTextPlain,
	"TEXT_HTML":        HttpResponseContentTypesTextHtml,
	"APPLICATION_JSON": HttpResponseContentTypesApplicationJson,
	"APPLICATION_XML":  HttpResponseContentTypesApplicationXml,
}

var mappingHttpResponseContentTypesEnumLowerCase = map[string]HttpResponseContentTypesEnum{
	"text_plain":       HttpResponseContentTypesTextPlain,
	"text_html":        HttpResponseContentTypesTextHtml,
	"application_json": HttpResponseContentTypesApplicationJson,
	"application_xml":  HttpResponseContentTypesApplicationXml,
}

// GetHttpResponseContentTypesEnumValues Enumerates the set of values for HttpResponseContentTypesEnum
func GetHttpResponseContentTypesEnumValues() []HttpResponseContentTypesEnum {
	values := make([]HttpResponseContentTypesEnum, 0)
	for _, v := range mappingHttpResponseContentTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpResponseContentTypesEnumStringValues Enumerates the set of values in String for HttpResponseContentTypesEnum
func GetHttpResponseContentTypesEnumStringValues() []string {
	return []string{
		"TEXT_PLAIN",
		"TEXT_HTML",
		"APPLICATION_JSON",
		"APPLICATION_XML",
	}
}

// GetMappingHttpResponseContentTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpResponseContentTypesEnum(val string) (HttpResponseContentTypesEnum, bool) {
	enum, ok := mappingHttpResponseContentTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
