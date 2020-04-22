// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// HealthCheck Health checks monitor the status of your origin servers and only route traffic to the origins that pass the health check. If the health check fails, origin is automatically removed from the load balancing.
// There is roughly one health check per EDGE POP per period. Any checks that pass will be reported as "healthy".
type HealthCheck struct {

	// Enables or disables the health checks.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// An HTTP verb (i.e. HEAD, GET, or POST) to use when performing the health check.
	Method HealthCheckMethodEnum `mandatory:"false" json:"method,omitempty"`

	// Path to visit on your origins when performing the health check.
	Path *string `mandatory:"false" json:"path"`

	// HTTP header fields to include in health check requests, expressed as `"name": "value"` properties. Because HTTP header field names are case-insensitive, any use of names that are case-insensitive equal to other names will be rejected. If Host is not specified, requests will include a Host header field with value matching the policy's protected domain. If User-Agent is not specified, requests will include a User-Agent header field with value "waf health checks".
	// **Note:** The only currently-supported header fields are Host and User-Agent.
	Headers map[string]string `mandatory:"false" json:"headers"`

	// The HTTP response codes that signify a healthy state.
	// - **2XX:** Success response code group.
	// - **3XX:** Redirection response code group.
	// - **4XX:** Client errors response code group.
	// - **5XX:** Server errors response code group.
	ExpectedResponseCodeGroup []HealthCheckExpectedResponseCodeGroupEnum `mandatory:"false" json:"expectedResponseCodeGroup,omitempty"`

	// Enables or disables additional check for predefined text in addition to response code.
	IsResponseTextCheckEnabled *bool `mandatory:"false" json:"isResponseTextCheckEnabled"`

	// Health check will search for the given text in a case-sensitive manner within the response body and will fail if the text is not found.
	ExpectedResponseText *string `mandatory:"false" json:"expectedResponseText"`

	// Time between health checks of an individual origin server, in seconds.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Response timeout represents wait time until request is considered failed, in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Number of successful health checks after which the server is marked up.
	HealthyThreshold *int `mandatory:"false" json:"healthyThreshold"`

	// Number of failed health checks after which the server is marked down.
	UnhealthyThreshold *int `mandatory:"false" json:"unhealthyThreshold"`
}

func (m HealthCheck) String() string {
	return common.PointerString(m)
}

// HealthCheckMethodEnum Enum with underlying type: string
type HealthCheckMethodEnum string

// Set of constants representing the allowable values for HealthCheckMethodEnum
const (
	HealthCheckMethodGet  HealthCheckMethodEnum = "GET"
	HealthCheckMethodHead HealthCheckMethodEnum = "HEAD"
	HealthCheckMethodPost HealthCheckMethodEnum = "POST"
)

var mappingHealthCheckMethod = map[string]HealthCheckMethodEnum{
	"GET":  HealthCheckMethodGet,
	"HEAD": HealthCheckMethodHead,
	"POST": HealthCheckMethodPost,
}

// GetHealthCheckMethodEnumValues Enumerates the set of values for HealthCheckMethodEnum
func GetHealthCheckMethodEnumValues() []HealthCheckMethodEnum {
	values := make([]HealthCheckMethodEnum, 0)
	for _, v := range mappingHealthCheckMethod {
		values = append(values, v)
	}
	return values
}

// HealthCheckExpectedResponseCodeGroupEnum Enum with underlying type: string
type HealthCheckExpectedResponseCodeGroupEnum string

// Set of constants representing the allowable values for HealthCheckExpectedResponseCodeGroupEnum
const (
	HealthCheckExpectedResponseCodeGroup2xx HealthCheckExpectedResponseCodeGroupEnum = "2XX"
	HealthCheckExpectedResponseCodeGroup3xx HealthCheckExpectedResponseCodeGroupEnum = "3XX"
	HealthCheckExpectedResponseCodeGroup4xx HealthCheckExpectedResponseCodeGroupEnum = "4XX"
	HealthCheckExpectedResponseCodeGroup5xx HealthCheckExpectedResponseCodeGroupEnum = "5XX"
)

var mappingHealthCheckExpectedResponseCodeGroup = map[string]HealthCheckExpectedResponseCodeGroupEnum{
	"2XX": HealthCheckExpectedResponseCodeGroup2xx,
	"3XX": HealthCheckExpectedResponseCodeGroup3xx,
	"4XX": HealthCheckExpectedResponseCodeGroup4xx,
	"5XX": HealthCheckExpectedResponseCodeGroup5xx,
}

// GetHealthCheckExpectedResponseCodeGroupEnumValues Enumerates the set of values for HealthCheckExpectedResponseCodeGroupEnum
func GetHealthCheckExpectedResponseCodeGroupEnumValues() []HealthCheckExpectedResponseCodeGroupEnum {
	values := make([]HealthCheckExpectedResponseCodeGroupEnum, 0)
	for _, v := range mappingHealthCheckExpectedResponseCodeGroup {
		values = append(values, v)
	}
	return values
}
