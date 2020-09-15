// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

// HttpProbeProtocolEnum Enum with underlying type: string
type HttpProbeProtocolEnum string

// Set of constants representing the allowable values for HttpProbeProtocolEnum
const (
	HttpProbeProtocolHttp  HttpProbeProtocolEnum = "HTTP"
	HttpProbeProtocolHttps HttpProbeProtocolEnum = "HTTPS"
)

var mappingHttpProbeProtocol = map[string]HttpProbeProtocolEnum{
	"HTTP":  HttpProbeProtocolHttp,
	"HTTPS": HttpProbeProtocolHttps,
}

// GetHttpProbeProtocolEnumValues Enumerates the set of values for HttpProbeProtocolEnum
func GetHttpProbeProtocolEnumValues() []HttpProbeProtocolEnum {
	values := make([]HttpProbeProtocolEnum, 0)
	for _, v := range mappingHttpProbeProtocol {
		values = append(values, v)
	}
	return values
}
