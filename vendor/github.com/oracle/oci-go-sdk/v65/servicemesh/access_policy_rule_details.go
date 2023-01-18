// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessPolicyRuleDetails Access policy rule.
type AccessPolicyRuleDetails struct {

	// Action for the traffic between the source and the destination.
	Action AccessPolicyRuleDetailsActionEnum `mandatory:"true" json:"action"`

	Source AccessPolicyTargetDetails `mandatory:"true" json:"source"`

	Destination AccessPolicyTargetDetails `mandatory:"true" json:"destination"`
}

func (m AccessPolicyRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessPolicyRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessPolicyRuleDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAccessPolicyRuleDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AccessPolicyRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Action      AccessPolicyRuleDetailsActionEnum `json:"action"`
		Source      accesspolicytargetdetails         `json:"source"`
		Destination accesspolicytargetdetails         `json:"destination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Action = model.Action

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(AccessPolicyTargetDetails)
	} else {
		m.Source = nil
	}

	nn, e = model.Destination.UnmarshalPolymorphicJSON(model.Destination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Destination = nn.(AccessPolicyTargetDetails)
	} else {
		m.Destination = nil
	}

	return
}

// AccessPolicyRuleDetailsActionEnum Enum with underlying type: string
type AccessPolicyRuleDetailsActionEnum string

// Set of constants representing the allowable values for AccessPolicyRuleDetailsActionEnum
const (
	AccessPolicyRuleDetailsActionAllow AccessPolicyRuleDetailsActionEnum = "ALLOW"
)

var mappingAccessPolicyRuleDetailsActionEnum = map[string]AccessPolicyRuleDetailsActionEnum{
	"ALLOW": AccessPolicyRuleDetailsActionAllow,
}

var mappingAccessPolicyRuleDetailsActionEnumLowerCase = map[string]AccessPolicyRuleDetailsActionEnum{
	"allow": AccessPolicyRuleDetailsActionAllow,
}

// GetAccessPolicyRuleDetailsActionEnumValues Enumerates the set of values for AccessPolicyRuleDetailsActionEnum
func GetAccessPolicyRuleDetailsActionEnumValues() []AccessPolicyRuleDetailsActionEnum {
	values := make([]AccessPolicyRuleDetailsActionEnum, 0)
	for _, v := range mappingAccessPolicyRuleDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessPolicyRuleDetailsActionEnumStringValues Enumerates the set of values in String for AccessPolicyRuleDetailsActionEnum
func GetAccessPolicyRuleDetailsActionEnumStringValues() []string {
	return []string{
		"ALLOW",
	}
}

// GetMappingAccessPolicyRuleDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessPolicyRuleDetailsActionEnum(val string) (AccessPolicyRuleDetailsActionEnum, bool) {
	enum, ok := mappingAccessPolicyRuleDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
