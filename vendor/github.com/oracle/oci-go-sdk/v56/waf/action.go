// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Action An object that represents action and its options.
// The action can be terminating, if it stops further execution of rules and modules.
// And non-terminating, if it does not interrupt execution flow.
type Action interface {

	// Action name. Can be used to reference the action.
	GetName() *string
}

type action struct {
	JsonData []byte
	Name     *string `mandatory:"true" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RETURN_HTTP_RESPONSE":
		mm := ReturnHttpResponseAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ALLOW":
		mm := AllowAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CHECK":
		mm := CheckAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m action) GetName() *string {
	return m.Name
}

func (m action) String() string {
	return common.PointerString(m)
}

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCheck              ActionTypeEnum = "CHECK"
	ActionTypeAllow              ActionTypeEnum = "ALLOW"
	ActionTypeReturnHttpResponse ActionTypeEnum = "RETURN_HTTP_RESPONSE"
)

var mappingActionType = map[string]ActionTypeEnum{
	"CHECK":                ActionTypeCheck,
	"ALLOW":                ActionTypeAllow,
	"RETURN_HTTP_RESPONSE": ActionTypeReturnHttpResponse,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionType {
		values = append(values, v)
	}
	return values
}
