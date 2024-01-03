// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"strings"
)

// HttpProbeMethodEnum Enum with underlying type: string
type HttpProbeMethodEnum string

// Set of constants representing the allowable values for HttpProbeMethodEnum
const (
	HttpProbeMethodGet  HttpProbeMethodEnum = "GET"
	HttpProbeMethodHead HttpProbeMethodEnum = "HEAD"
)

var mappingHttpProbeMethodEnum = map[string]HttpProbeMethodEnum{
	"GET":  HttpProbeMethodGet,
	"HEAD": HttpProbeMethodHead,
}

var mappingHttpProbeMethodEnumLowerCase = map[string]HttpProbeMethodEnum{
	"get":  HttpProbeMethodGet,
	"head": HttpProbeMethodHead,
}

// GetHttpProbeMethodEnumValues Enumerates the set of values for HttpProbeMethodEnum
func GetHttpProbeMethodEnumValues() []HttpProbeMethodEnum {
	values := make([]HttpProbeMethodEnum, 0)
	for _, v := range mappingHttpProbeMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpProbeMethodEnumStringValues Enumerates the set of values in String for HttpProbeMethodEnum
func GetHttpProbeMethodEnumStringValues() []string {
	return []string{
		"GET",
		"HEAD",
	}
}

// GetMappingHttpProbeMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpProbeMethodEnum(val string) (HttpProbeMethodEnum, bool) {
	enum, ok := mappingHttpProbeMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
