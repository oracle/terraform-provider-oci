// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
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
		return *m, nil
	}
}

func (m rule) String() string {
	return common.PointerString(m)
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

var mappingRuleAction = map[string]RuleActionEnum{
	"ADD_HTTP_REQUEST_HEADER":           RuleActionAddHttpRequestHeader,
	"EXTEND_HTTP_REQUEST_HEADER_VALUE":  RuleActionExtendHttpRequestHeaderValue,
	"REMOVE_HTTP_REQUEST_HEADER":        RuleActionRemoveHttpRequestHeader,
	"ADD_HTTP_RESPONSE_HEADER":          RuleActionAddHttpResponseHeader,
	"EXTEND_HTTP_RESPONSE_HEADER_VALUE": RuleActionExtendHttpResponseHeaderValue,
	"REMOVE_HTTP_RESPONSE_HEADER":       RuleActionRemoveHttpResponseHeader,
	"ALLOW": RuleActionAllow,
	"CONTROL_ACCESS_USING_HTTP_METHODS": RuleActionControlAccessUsingHttpMethods,
	"REDIRECT":                          RuleActionRedirect,
	"HTTP_HEADER":                       RuleActionHttpHeader,
}

// GetRuleActionEnumValues Enumerates the set of values for RuleActionEnum
func GetRuleActionEnumValues() []RuleActionEnum {
	values := make([]RuleActionEnum, 0)
	for _, v := range mappingRuleAction {
		values = append(values, v)
	}
	return values
}
