// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Rule An object that represents an action to apply to a listener.
type Rule interface {
}

type rule struct {
	JsonData []byte
	Action   string `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *rule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrule rule
	s := struct {
		Model Unmarshalerrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *rule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "ADD_HTTP_REQUEST_HEADER":
		mm := AddHttpRequestHeaderRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REDIRECT":
		mm := RedirectRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE_HTTP_REQUEST_HEADER":
		mm := RemoveHttpRequestHeaderRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTEND_HTTP_REQUEST_HEADER_VALUE":
		mm := ExtendHttpRequestHeaderValueRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOVE_HTTP_RESPONSE_HEADER":
		mm := RemoveHttpResponseHeaderRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTROL_ACCESS_USING_HTTP_METHODS":
		mm := ControlAccessUsingHttpMethodsRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALLOW":
		mm := AllowRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_HEADER":
		mm := HttpHeaderRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_HTTP_RESPONSE_HEADER":
		mm := AddHttpResponseHeaderRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTEND_HTTP_RESPONSE_HEADER_VALUE":
		mm := ExtendHttpResponseHeaderValueRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Rule: %s.", m.Action)
		return *m, nil
	}
}

func (m rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleActionEnum Enum with underlying type: string
type RuleActionEnum string

// Set of constants representing the allowable values for RuleActionEnum
const (
	RuleActionAddHttpRequestHeader          RuleActionEnum = "ADD_HTTP_REQUEST_HEADER"
	RuleActionExtendHttpRequestHeaderValue  RuleActionEnum = "EXTEND_HTTP_REQUEST_HEADER_VALUE"
	RuleActionRemoveHttpRequestHeader       RuleActionEnum = "REMOVE_HTTP_REQUEST_HEADER"
	RuleActionAddHttpResponseHeader         RuleActionEnum = "ADD_HTTP_RESPONSE_HEADER"
	RuleActionExtendHttpResponseHeaderValue RuleActionEnum = "EXTEND_HTTP_RESPONSE_HEADER_VALUE"
	RuleActionRemoveHttpResponseHeader      RuleActionEnum = "REMOVE_HTTP_RESPONSE_HEADER"
	RuleActionAllow                         RuleActionEnum = "ALLOW"
	RuleActionControlAccessUsingHttpMethods RuleActionEnum = "CONTROL_ACCESS_USING_HTTP_METHODS"
	RuleActionRedirect                      RuleActionEnum = "REDIRECT"
	RuleActionHttpHeader                    RuleActionEnum = "HTTP_HEADER"
)

var mappingRuleActionEnum = map[string]RuleActionEnum{
	"ADD_HTTP_REQUEST_HEADER":           RuleActionAddHttpRequestHeader,
	"EXTEND_HTTP_REQUEST_HEADER_VALUE":  RuleActionExtendHttpRequestHeaderValue,
	"REMOVE_HTTP_REQUEST_HEADER":        RuleActionRemoveHttpRequestHeader,
	"ADD_HTTP_RESPONSE_HEADER":          RuleActionAddHttpResponseHeader,
	"EXTEND_HTTP_RESPONSE_HEADER_VALUE": RuleActionExtendHttpResponseHeaderValue,
	"REMOVE_HTTP_RESPONSE_HEADER":       RuleActionRemoveHttpResponseHeader,
	"ALLOW":                             RuleActionAllow,
	"CONTROL_ACCESS_USING_HTTP_METHODS": RuleActionControlAccessUsingHttpMethods,
	"REDIRECT":                          RuleActionRedirect,
	"HTTP_HEADER":                       RuleActionHttpHeader,
}

var mappingRuleActionEnumLowerCase = map[string]RuleActionEnum{
	"add_http_request_header":           RuleActionAddHttpRequestHeader,
	"extend_http_request_header_value":  RuleActionExtendHttpRequestHeaderValue,
	"remove_http_request_header":        RuleActionRemoveHttpRequestHeader,
	"add_http_response_header":          RuleActionAddHttpResponseHeader,
	"extend_http_response_header_value": RuleActionExtendHttpResponseHeaderValue,
	"remove_http_response_header":       RuleActionRemoveHttpResponseHeader,
	"allow":                             RuleActionAllow,
	"control_access_using_http_methods": RuleActionControlAccessUsingHttpMethods,
	"redirect":                          RuleActionRedirect,
	"http_header":                       RuleActionHttpHeader,
}

// GetRuleActionEnumValues Enumerates the set of values for RuleActionEnum
func GetRuleActionEnumValues() []RuleActionEnum {
	values := make([]RuleActionEnum, 0)
	for _, v := range mappingRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleActionEnumStringValues Enumerates the set of values in String for RuleActionEnum
func GetRuleActionEnumStringValues() []string {
	return []string{
		"ADD_HTTP_REQUEST_HEADER",
		"EXTEND_HTTP_REQUEST_HEADER_VALUE",
		"REMOVE_HTTP_REQUEST_HEADER",
		"ADD_HTTP_RESPONSE_HEADER",
		"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
		"REMOVE_HTTP_RESPONSE_HEADER",
		"ALLOW",
		"CONTROL_ACCESS_USING_HTTP_METHODS",
		"REDIRECT",
		"HTTP_HEADER",
	}
}

// GetMappingRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleActionEnum(val string) (RuleActionEnum, bool) {
	enum, ok := mappingRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
