// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestDetails Details of the api request.
type RequestDetails struct {

	// Endpoint to be invoked.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// Method of the api.
	Method RequestDetailsMethodEnum `mandatory:"true" json:"method"`

	// request body of the api.
	Payload *string `mandatory:"false" json:"payload"`

	// Key-value pairs of query parameters to be included.
	// Example: `{"bar-key": "value"}`
	QueryParams map[string]string `mandatory:"false" json:"queryParams"`

	// Key-value pairs of headers to be included.
	// Example: `{"bar-key": "value"}`
	Headers map[string]string `mandatory:"false" json:"headers"`

	AuthenticationDetails AuthenticationDetails `mandatory:"false" json:"authenticationDetails"`
}

func (m RequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRequestDetailsMethodEnum(string(m.Method)); !ok && m.Method != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Method: %s. Supported values are: %s.", m.Method, strings.Join(GetRequestDetailsMethodEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RequestDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Payload               *string                  `json:"payload"`
		QueryParams           map[string]string        `json:"queryParams"`
		Headers               map[string]string        `json:"headers"`
		AuthenticationDetails authenticationdetails    `json:"authenticationDetails"`
		Endpoint              *string                  `json:"endpoint"`
		Method                RequestDetailsMethodEnum `json:"method"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Payload = model.Payload

	m.QueryParams = model.QueryParams

	m.Headers = model.Headers

	nn, e = model.AuthenticationDetails.UnmarshalPolymorphicJSON(model.AuthenticationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthenticationDetails = nn.(AuthenticationDetails)
	} else {
		m.AuthenticationDetails = nil
	}

	m.Endpoint = model.Endpoint

	m.Method = model.Method

	return
}

// RequestDetailsMethodEnum Enum with underlying type: string
type RequestDetailsMethodEnum string

// Set of constants representing the allowable values for RequestDetailsMethodEnum
const (
	RequestDetailsMethodGet    RequestDetailsMethodEnum = "GET"
	RequestDetailsMethodPost   RequestDetailsMethodEnum = "POST"
	RequestDetailsMethodPut    RequestDetailsMethodEnum = "PUT"
	RequestDetailsMethodDelete RequestDetailsMethodEnum = "DELETE"
	RequestDetailsMethodPatch  RequestDetailsMethodEnum = "PATCH"
)

var mappingRequestDetailsMethodEnum = map[string]RequestDetailsMethodEnum{
	"GET":    RequestDetailsMethodGet,
	"POST":   RequestDetailsMethodPost,
	"PUT":    RequestDetailsMethodPut,
	"DELETE": RequestDetailsMethodDelete,
	"PATCH":  RequestDetailsMethodPatch,
}

var mappingRequestDetailsMethodEnumLowerCase = map[string]RequestDetailsMethodEnum{
	"get":    RequestDetailsMethodGet,
	"post":   RequestDetailsMethodPost,
	"put":    RequestDetailsMethodPut,
	"delete": RequestDetailsMethodDelete,
	"patch":  RequestDetailsMethodPatch,
}

// GetRequestDetailsMethodEnumValues Enumerates the set of values for RequestDetailsMethodEnum
func GetRequestDetailsMethodEnumValues() []RequestDetailsMethodEnum {
	values := make([]RequestDetailsMethodEnum, 0)
	for _, v := range mappingRequestDetailsMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestDetailsMethodEnumStringValues Enumerates the set of values in String for RequestDetailsMethodEnum
func GetRequestDetailsMethodEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}
}

// GetMappingRequestDetailsMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestDetailsMethodEnum(val string) (RequestDetailsMethodEnum, bool) {
	enum, ok := mappingRequestDetailsMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
