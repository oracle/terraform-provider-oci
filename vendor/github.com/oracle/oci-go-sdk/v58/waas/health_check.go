// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHealthCheckMethodEnum(string(m.Method)); !ok && m.Method != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Method: %s. Supported values are: %s.", m.Method, strings.Join(GetHealthCheckMethodEnumStringValues(), ",")))
	}
	for _, val := range m.ExpectedResponseCodeGroup {
		if _, ok := GetMappingHealthCheckExpectedResponseCodeGroupEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedResponseCodeGroup: %s. Supported values are: %s.", val, strings.Join(GetHealthCheckExpectedResponseCodeGroupEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HealthCheckMethodEnum Enum with underlying type: string
type HealthCheckMethodEnum string

// Set of constants representing the allowable values for HealthCheckMethodEnum
const (
	HealthCheckMethodGet  HealthCheckMethodEnum = "GET"
	HealthCheckMethodHead HealthCheckMethodEnum = "HEAD"
	HealthCheckMethodPost HealthCheckMethodEnum = "POST"
)

var mappingHealthCheckMethodEnum = map[string]HealthCheckMethodEnum{
	"GET":  HealthCheckMethodGet,
	"HEAD": HealthCheckMethodHead,
	"POST": HealthCheckMethodPost,
}

// GetHealthCheckMethodEnumValues Enumerates the set of values for HealthCheckMethodEnum
func GetHealthCheckMethodEnumValues() []HealthCheckMethodEnum {
	values := make([]HealthCheckMethodEnum, 0)
	for _, v := range mappingHealthCheckMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckMethodEnumStringValues Enumerates the set of values in String for HealthCheckMethodEnum
func GetHealthCheckMethodEnumStringValues() []string {
	return []string{
		"GET",
		"HEAD",
		"POST",
	}
}

// GetMappingHealthCheckMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckMethodEnum(val string) (HealthCheckMethodEnum, bool) {
	mappingHealthCheckMethodEnumIgnoreCase := make(map[string]HealthCheckMethodEnum)
	for k, v := range mappingHealthCheckMethodEnum {
		mappingHealthCheckMethodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHealthCheckMethodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingHealthCheckExpectedResponseCodeGroupEnum = map[string]HealthCheckExpectedResponseCodeGroupEnum{
	"2XX": HealthCheckExpectedResponseCodeGroup2xx,
	"3XX": HealthCheckExpectedResponseCodeGroup3xx,
	"4XX": HealthCheckExpectedResponseCodeGroup4xx,
	"5XX": HealthCheckExpectedResponseCodeGroup5xx,
}

// GetHealthCheckExpectedResponseCodeGroupEnumValues Enumerates the set of values for HealthCheckExpectedResponseCodeGroupEnum
func GetHealthCheckExpectedResponseCodeGroupEnumValues() []HealthCheckExpectedResponseCodeGroupEnum {
	values := make([]HealthCheckExpectedResponseCodeGroupEnum, 0)
	for _, v := range mappingHealthCheckExpectedResponseCodeGroupEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthCheckExpectedResponseCodeGroupEnumStringValues Enumerates the set of values in String for HealthCheckExpectedResponseCodeGroupEnum
func GetHealthCheckExpectedResponseCodeGroupEnumStringValues() []string {
	return []string{
		"2XX",
		"3XX",
		"4XX",
		"5XX",
	}
}

// GetMappingHealthCheckExpectedResponseCodeGroupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthCheckExpectedResponseCodeGroupEnum(val string) (HealthCheckExpectedResponseCodeGroupEnum, bool) {
	mappingHealthCheckExpectedResponseCodeGroupEnumIgnoreCase := make(map[string]HealthCheckExpectedResponseCodeGroupEnum)
	for k, v := range mappingHealthCheckExpectedResponseCodeGroupEnum {
		mappingHealthCheckExpectedResponseCodeGroupEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingHealthCheckExpectedResponseCodeGroupEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
