// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EndpointRequest An object containing details to make a REST request.
type EndpointRequest struct {

	// The request URL.
	Url *string `mandatory:"true" json:"url"`

	// The endpoint method - GET or POST.
	Method EndpointRequestMethodEnum `mandatory:"false" json:"method,omitempty"`

	// The request content type.
	ContentType *string `mandatory:"false" json:"contentType"`

	// The request payload, applicable for POST requests.
	Payload *string `mandatory:"false" json:"payload"`

	// The request headers represented as a list of name-value pairs.
	Headers []NameValuePair `mandatory:"false" json:"headers"`

	// The request form parameters represented as a list of name-value pairs.
	FormParameters []NameValuePair `mandatory:"false" json:"formParameters"`
}

func (m EndpointRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EndpointRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEndpointRequestMethodEnum(string(m.Method)); !ok && m.Method != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Method: %s. Supported values are: %s.", m.Method, strings.Join(GetEndpointRequestMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointRequestMethodEnum Enum with underlying type: string
type EndpointRequestMethodEnum string

// Set of constants representing the allowable values for EndpointRequestMethodEnum
const (
	EndpointRequestMethodGet  EndpointRequestMethodEnum = "GET"
	EndpointRequestMethodPost EndpointRequestMethodEnum = "POST"
)

var mappingEndpointRequestMethodEnum = map[string]EndpointRequestMethodEnum{
	"GET":  EndpointRequestMethodGet,
	"POST": EndpointRequestMethodPost,
}

var mappingEndpointRequestMethodEnumLowerCase = map[string]EndpointRequestMethodEnum{
	"get":  EndpointRequestMethodGet,
	"post": EndpointRequestMethodPost,
}

// GetEndpointRequestMethodEnumValues Enumerates the set of values for EndpointRequestMethodEnum
func GetEndpointRequestMethodEnumValues() []EndpointRequestMethodEnum {
	values := make([]EndpointRequestMethodEnum, 0)
	for _, v := range mappingEndpointRequestMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointRequestMethodEnumStringValues Enumerates the set of values in String for EndpointRequestMethodEnum
func GetEndpointRequestMethodEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
	}
}

// GetMappingEndpointRequestMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointRequestMethodEnum(val string) (EndpointRequestMethodEnum, bool) {
	enum, ok := mappingEndpointRequestMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
