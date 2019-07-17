// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

// HttpProbeMethodEnum Enum with underlying type: string
type HttpProbeMethodEnum string

// Set of constants representing the allowable values for HttpProbeMethodEnum
const (
	HttpProbeMethodGet  HttpProbeMethodEnum = "GET"
	HttpProbeMethodHead HttpProbeMethodEnum = "HEAD"
)

var mappingHttpProbeMethod = map[string]HttpProbeMethodEnum{
	"GET":  HttpProbeMethodGet,
	"HEAD": HttpProbeMethodHead,
}

// GetHttpProbeMethodEnumValues Enumerates the set of values for HttpProbeMethodEnum
func GetHttpProbeMethodEnumValues() []HttpProbeMethodEnum {
	values := make([]HttpProbeMethodEnum, 0)
	for _, v := range mappingHttpProbeMethod {
		values = append(values, v)
	}
	return values
}
