// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestMonitorConfiguration Request configuration details for the REST monitor type.
type RestMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// If redirection is enabled, then redirects will be allowed while accessing target URL.
	IsRedirectionEnabled *bool `mandatory:"false" json:"isRedirectionEnabled"`

	// If certificate validation is enabled, then call will fail for certificate errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	ReqAuthenticationDetails *RequestAuthenticationDetails `mandatory:"false" json:"reqAuthenticationDetails"`

	ClientCertificateDetails *ClientCertificateDetails `mandatory:"false" json:"clientCertificateDetails"`

	// List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]`
	RequestHeaders []Header `mandatory:"false" json:"requestHeaders"`

	// List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]`
	RequestQueryParams []RequestQueryParam `mandatory:"false" json:"requestQueryParams"`

	// Request post body content.
	RequestPostBody *string `mandatory:"false" json:"requestPostBody"`

	// Verify response content against regular expression based string.
	// If response content does not match the verifyResponseContent value, then it will be considered a failure.
	VerifyResponseContent *string `mandatory:"false" json:"verifyResponseContent"`

	// Expected HTTP response codes. For status code range, set values such as 2xx, 3xx.
	VerifyResponseCodes []string `mandatory:"false" json:"verifyResponseCodes"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// Request HTTP method.
	RequestMethod RequestMethodsEnum `mandatory:"false" json:"requestMethod,omitempty"`

	// Request HTTP authentication scheme.
	ReqAuthenticationScheme RequestAuthenticationSchemesEnum `mandatory:"false" json:"reqAuthenticationScheme,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m RestMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m RestMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m RestMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequestMethodsEnum(string(m.RequestMethod)); !ok && m.RequestMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestMethod: %s. Supported values are: %s.", m.RequestMethod, strings.Join(GetRequestMethodsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestAuthenticationSchemesEnum(string(m.ReqAuthenticationScheme)); !ok && m.ReqAuthenticationScheme != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReqAuthenticationScheme: %s. Supported values are: %s.", m.ReqAuthenticationScheme, strings.Join(GetRequestAuthenticationSchemesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RestMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRestMonitorConfiguration RestMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeRestMonitorConfiguration
	}{
		"REST_CONFIG",
		(MarshalTypeRestMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
