// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HttpProbeResultSummary The results returned by running an HTTP probe.  All times and durations are
// returned in milliseconds. All times are relative to the POSIX epoch
// (1970-01-01T00:00Z). Time properties conform to W3C Resource Timing.
// For more information, see
// PerformanceResourceTiming (https://w3c.github.io/resource-timing/#sec-resource-timing)
// interface.
type HttpProbeResultSummary struct {

	// A value identifying this specific probe result. The key is only unique within
	// the results of its probe configuration. The key may be reused after 90 days.
	Key *string `mandatory:"false" json:"key"`

	// The OCID of the monitor or on-demand probe responsible for creating this result.
	ProbeConfigurationId *string `mandatory:"false" json:"probeConfigurationId"`

	// The date and time the probe was executed, expressed in milliseconds since the
	// POSIX epoch. This field is defined by the PerformanceResourceTiming interface
	// of the W3C Resource Timing specification. For more information, see
	// Resource Timing (https://w3c.github.io/resource-timing/#sec-resource-timing).
	StartTime *float64 `mandatory:"false" json:"startTime"`

	// The target hostname or IP address of the probe.
	Target *string `mandatory:"false" json:"target"`

	// The name of the vantage point that executed the probe.
	VantagePointName *string `mandatory:"false" json:"vantagePointName"`

	// True if the probe did not complete before the configured `timeoutInSeconds` value.
	IsTimedOut *bool `mandatory:"false" json:"isTimedOut"`

	// True if the probe result is determined to be healthy based on probe
	// type-specific criteria.  For HTTP probes, a probe result is considered
	// healthy if the HTTP response code is greater than or equal to 200 and
	// less than 300.
	IsHealthy *bool `mandatory:"false" json:"isHealthy"`

	// The category of error if an error occurs executing the probe.
	// The `errorMessage` field provides a message with the error details.
	// * NONE - No error
	// * DNS - DNS errors
	// * TRANSPORT - Transport-related errors, for example a "TLS certificate expired" error.
	// * NETWORK - Network-related errors, for example a "network unreachable" error.
	// * SYSTEM - Internal system errors.
	ErrorCategory HttpProbeResultSummaryErrorCategoryEnum `mandatory:"false" json:"errorCategory,omitempty"`

	// The error information indicating why a probe execution failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	Protocol HttpProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`

	Connection *TcpConnection `mandatory:"false" json:"connection"`

	Dns *Dns `mandatory:"false" json:"dns"`

	// The HTTP response status code.
	StatusCode *int `mandatory:"false" json:"statusCode"`

	// The time immediately before the vantage point starts the domain name lookup for
	// the resource.
	DomainLookupStart *float64 `mandatory:"false" json:"domainLookupStart"`

	// The time immediately before the vantage point finishes the domain name lookup for
	// the resource.
	DomainLookupEnd *float64 `mandatory:"false" json:"domainLookupEnd"`

	// The time immediately before the vantage point starts establishing the connection
	// to the server to retrieve the resource.
	ConnectStart *float64 `mandatory:"false" json:"connectStart"`

	// The time immediately before the vantage point starts the handshake process to
	// secure the current connection.
	SecureConnectionStart *float64 `mandatory:"false" json:"secureConnectionStart"`

	// The time immediately after the vantage point finishes establishing the connection
	// to the server to retrieve the resource.
	ConnectEnd *float64 `mandatory:"false" json:"connectEnd"`

	// The time immediately before the vantage point starts to fetch the resource.
	FetchStart *float64 `mandatory:"false" json:"fetchStart"`

	// The time immediately before the vantage point starts requesting the resource from
	// the server.
	RequestStart *float64 `mandatory:"false" json:"requestStart"`

	// The time immediately after the vantage point's HTTP parser receives the first byte
	// of the response.
	ResponseStart *float64 `mandatory:"false" json:"responseStart"`

	// The time immediately after the vantage point receives the last byte of the response
	// or immediately before the transport connection is closed, whichever comes first.
	ResponseEnd *float64 `mandatory:"false" json:"responseEnd"`

	// The total duration from start of request until response is fully consumed or the
	// connection is closed.
	Duration *float64 `mandatory:"false" json:"duration"`

	// The size, in octets, of the payload body prior to removing any applied
	// content-codings.
	EncodedBodySize *int `mandatory:"false" json:"encodedBodySize"`
}

func (m HttpProbeResultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpProbeResultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHttpProbeResultSummaryErrorCategoryEnum(string(m.ErrorCategory)); !ok && m.ErrorCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ErrorCategory: %s. Supported values are: %s.", m.ErrorCategory, strings.Join(GetHttpProbeResultSummaryErrorCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHttpProbeProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetHttpProbeProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HttpProbeResultSummaryErrorCategoryEnum Enum with underlying type: string
type HttpProbeResultSummaryErrorCategoryEnum string

// Set of constants representing the allowable values for HttpProbeResultSummaryErrorCategoryEnum
const (
	HttpProbeResultSummaryErrorCategoryNone      HttpProbeResultSummaryErrorCategoryEnum = "NONE"
	HttpProbeResultSummaryErrorCategoryDns       HttpProbeResultSummaryErrorCategoryEnum = "DNS"
	HttpProbeResultSummaryErrorCategoryTransport HttpProbeResultSummaryErrorCategoryEnum = "TRANSPORT"
	HttpProbeResultSummaryErrorCategoryNetwork   HttpProbeResultSummaryErrorCategoryEnum = "NETWORK"
	HttpProbeResultSummaryErrorCategorySystem    HttpProbeResultSummaryErrorCategoryEnum = "SYSTEM"
)

var mappingHttpProbeResultSummaryErrorCategoryEnum = map[string]HttpProbeResultSummaryErrorCategoryEnum{
	"NONE":      HttpProbeResultSummaryErrorCategoryNone,
	"DNS":       HttpProbeResultSummaryErrorCategoryDns,
	"TRANSPORT": HttpProbeResultSummaryErrorCategoryTransport,
	"NETWORK":   HttpProbeResultSummaryErrorCategoryNetwork,
	"SYSTEM":    HttpProbeResultSummaryErrorCategorySystem,
}

var mappingHttpProbeResultSummaryErrorCategoryEnumLowerCase = map[string]HttpProbeResultSummaryErrorCategoryEnum{
	"none":      HttpProbeResultSummaryErrorCategoryNone,
	"dns":       HttpProbeResultSummaryErrorCategoryDns,
	"transport": HttpProbeResultSummaryErrorCategoryTransport,
	"network":   HttpProbeResultSummaryErrorCategoryNetwork,
	"system":    HttpProbeResultSummaryErrorCategorySystem,
}

// GetHttpProbeResultSummaryErrorCategoryEnumValues Enumerates the set of values for HttpProbeResultSummaryErrorCategoryEnum
func GetHttpProbeResultSummaryErrorCategoryEnumValues() []HttpProbeResultSummaryErrorCategoryEnum {
	values := make([]HttpProbeResultSummaryErrorCategoryEnum, 0)
	for _, v := range mappingHttpProbeResultSummaryErrorCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpProbeResultSummaryErrorCategoryEnumStringValues Enumerates the set of values in String for HttpProbeResultSummaryErrorCategoryEnum
func GetHttpProbeResultSummaryErrorCategoryEnumStringValues() []string {
	return []string{
		"NONE",
		"DNS",
		"TRANSPORT",
		"NETWORK",
		"SYSTEM",
	}
}

// GetMappingHttpProbeResultSummaryErrorCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpProbeResultSummaryErrorCategoryEnum(val string) (HttpProbeResultSummaryErrorCategoryEnum, bool) {
	enum, ok := mappingHttpProbeResultSummaryErrorCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HttpProbeResultSummaryProtocolEnum is an alias to type: HttpProbeProtocolEnum
// Consider using HttpProbeProtocolEnum instead
// Deprecated
type HttpProbeResultSummaryProtocolEnum = HttpProbeProtocolEnum

// Set of constants representing the allowable values for HttpProbeProtocolEnum
// Deprecated
const (
	HttpProbeResultSummaryProtocolHttp  HttpProbeProtocolEnum = "HTTP"
	HttpProbeResultSummaryProtocolHttps HttpProbeProtocolEnum = "HTTPS"
)

// GetHttpProbeResultSummaryProtocolEnumValues Enumerates the set of values for HttpProbeProtocolEnum
// Consider using GetHttpProbeProtocolEnumValue
// Deprecated
var GetHttpProbeResultSummaryProtocolEnumValues = GetHttpProbeProtocolEnumValues
