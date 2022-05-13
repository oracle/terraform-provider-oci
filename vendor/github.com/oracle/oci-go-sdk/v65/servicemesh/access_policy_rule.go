// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AccessPolicyRule Access policy rule.
type AccessPolicyRule struct {

	// Action for the traffic between the source and the destination.
	Action AccessPolicyRuleActionEnum `mandatory:"true" json:"action"`

	Source AccessPolicyTarget `mandatory:"true" json:"source"`

	Destination AccessPolicyTarget `mandatory:"true" json:"destination"`
}

func (m AccessPolicyRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessPolicyRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessPolicyRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAccessPolicyRuleActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AccessPolicyRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Action      AccessPolicyRuleActionEnum `json:"action"`
		Source      accesspolicytarget         `json:"source"`
		Destination accesspolicytarget         `json:"destination"`
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
		m.Source = nn.(AccessPolicyTarget)
	} else {
		m.Source = nil
	}

	nn, e = model.Destination.UnmarshalPolymorphicJSON(model.Destination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Destination = nn.(AccessPolicyTarget)
	} else {
		m.Destination = nil
	}

	return
}

// AccessPolicyRuleActionEnum Enum with underlying type: string
type AccessPolicyRuleActionEnum string

// Set of constants representing the allowable values for AccessPolicyRuleActionEnum
const (
	AccessPolicyRuleActionAllow AccessPolicyRuleActionEnum = "ALLOW"
)

var mappingAccessPolicyRuleActionEnum = map[string]AccessPolicyRuleActionEnum{
	"ALLOW": AccessPolicyRuleActionAllow,
}

var mappingAccessPolicyRuleActionEnumLowerCase = map[string]AccessPolicyRuleActionEnum{
	"allow": AccessPolicyRuleActionAllow,
}

// GetAccessPolicyRuleActionEnumValues Enumerates the set of values for AccessPolicyRuleActionEnum
func GetAccessPolicyRuleActionEnumValues() []AccessPolicyRuleActionEnum {
	values := make([]AccessPolicyRuleActionEnum, 0)
	for _, v := range mappingAccessPolicyRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessPolicyRuleActionEnumStringValues Enumerates the set of values in String for AccessPolicyRuleActionEnum
func GetAccessPolicyRuleActionEnumStringValues() []string {
	return []string{
		"ALLOW",
	}
}

// GetMappingAccessPolicyRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessPolicyRuleActionEnum(val string) (AccessPolicyRuleActionEnum, bool) {
	enum, ok := mappingAccessPolicyRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
