// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RestMonitorConfiguration Request configuration details for the REST monitor type.
type RestMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	// If redirection enabled, then redirects will be allowed while accessing target URL.
	IsRedirectionEnabled *bool `mandatory:"false" json:"isRedirectionEnabled"`

	// If certificate validation enabled, then call will fail for certificate errors.
	IsCertificateValidationEnabled *bool `mandatory:"false" json:"isCertificateValidationEnabled"`

	ReqAuthenticationDetails *RequestAuthenticationDetails `mandatory:"false" json:"reqAuthenticationDetails"`

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

	// Request http authentication scheme.
	ReqAuthenticationScheme RequestAuthenticationSchemesEnum `mandatory:"false" json:"reqAuthenticationScheme,omitempty"`
}

//GetIsFailureRetried returns IsFailureRetried
func (m RestMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

func (m RestMonitorConfiguration) String() string {
	return common.PointerString(m)
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
