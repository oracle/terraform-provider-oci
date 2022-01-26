// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingCompositeConditionCompositeOperator = map[string]CompositeConditionCompositeOperatorEnum{
	"AND": CompositeConditionCompositeOperatorAnd,
	"OR":  CompositeConditionCompositeOperatorOr,
}

// GetCompositeConditionCompositeOperatorEnumValues Enumerates the set of values for CompositeConditionCompositeOperatorEnum
func GetCompositeConditionCompositeOperatorEnumValues() []CompositeConditionCompositeOperatorEnum {
	values := make([]CompositeConditionCompositeOperatorEnum, 0)
	for _, v := range mappingCompositeConditionCompositeOperator {
		values = append(values, v)
	}
	return values
}
