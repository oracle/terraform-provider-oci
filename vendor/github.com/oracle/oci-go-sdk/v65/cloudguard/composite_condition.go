// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CompositeCondition Composite Condition object with nested Condition
type CompositeCondition struct {
	LeftOperand Condition `mandatory:"false" json:"leftOperand"`

	RightOperand Condition `mandatory:"false" json:"rightOperand"`

	CompositeOperator CompositeConditionCompositeOperatorEnum `mandatory:"false" json:"compositeOperator,omitempty"`
}

func (m CompositeCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompositeCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompositeConditionCompositeOperatorEnum(string(m.CompositeOperator)); !ok && m.CompositeOperator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompositeOperator: %s. Supported values are: %s.", m.CompositeOperator, strings.Join(GetCompositeConditionCompositeOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CompositeCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompositeCondition CompositeCondition
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeCompositeCondition
	}{
		"COMPOSITE",
		(MarshalTypeCompositeCondition)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CompositeCondition) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LeftOperand       condition                               `json:"leftOperand"`
		CompositeOperator CompositeConditionCompositeOperatorEnum `json:"compositeOperator"`
		RightOperand      condition                               `json:"rightOperand"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.LeftOperand.UnmarshalPolymorphicJSON(model.LeftOperand.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LeftOperand = nn.(Condition)
	} else {
		m.LeftOperand = nil
	}

	m.CompositeOperator = model.CompositeOperator

	nn, e = model.RightOperand.UnmarshalPolymorphicJSON(model.RightOperand.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RightOperand = nn.(Condition)
	} else {
		m.RightOperand = nil
	}

	return
}

// CompositeConditionCompositeOperatorEnum Enum with underlying type: string
type CompositeConditionCompositeOperatorEnum string

// Set of constants representing the allowable values for CompositeConditionCompositeOperatorEnum
const (
	CompositeConditionCompositeOperatorAnd CompositeConditionCompositeOperatorEnum = "AND"
	CompositeConditionCompositeOperatorOr  CompositeConditionCompositeOperatorEnum = "OR"
)

var mappingCompositeConditionCompositeOperatorEnum = map[string]CompositeConditionCompositeOperatorEnum{
	"AND": CompositeConditionCompositeOperatorAnd,
	"OR":  CompositeConditionCompositeOperatorOr,
}

var mappingCompositeConditionCompositeOperatorEnumLowerCase = map[string]CompositeConditionCompositeOperatorEnum{
	"and": CompositeConditionCompositeOperatorAnd,
	"or":  CompositeConditionCompositeOperatorOr,
}

// GetCompositeConditionCompositeOperatorEnumValues Enumerates the set of values for CompositeConditionCompositeOperatorEnum
func GetCompositeConditionCompositeOperatorEnumValues() []CompositeConditionCompositeOperatorEnum {
	values := make([]CompositeConditionCompositeOperatorEnum, 0)
	for _, v := range mappingCompositeConditionCompositeOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetCompositeConditionCompositeOperatorEnumStringValues Enumerates the set of values in String for CompositeConditionCompositeOperatorEnum
func GetCompositeConditionCompositeOperatorEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
	}
}

// GetMappingCompositeConditionCompositeOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompositeConditionCompositeOperatorEnum(val string) (CompositeConditionCompositeOperatorEnum, bool) {
	enum, ok := mappingCompositeConditionCompositeOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
