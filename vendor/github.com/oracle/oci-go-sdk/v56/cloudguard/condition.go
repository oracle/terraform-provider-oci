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

// Condition Base condition object
type Condition interface {
}

type condition struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *condition) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercondition condition
	s := struct {
		Model Unmarshalercondition
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *condition) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "SIMPLE":
		mm := SimpleCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE":
		mm := CompositeCondition{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m condition) String() string {
	return common.PointerString(m)
}

// ConditionKindEnum Enum with underlying type: string
type ConditionKindEnum string

// Set of constants representing the allowable values for ConditionKindEnum
const (
	ConditionKindComposite ConditionKindEnum = "COMPOSITE"
	ConditionKindSimple    ConditionKindEnum = "SIMPLE"
)

var mappingConditionKind = map[string]ConditionKindEnum{
	"COMPOSITE": ConditionKindComposite,
	"SIMPLE":    ConditionKindSimple,
}

// GetConditionKindEnumValues Enumerates the set of values for ConditionKindEnum
func GetConditionKindEnumValues() []ConditionKindEnum {
	values := make([]ConditionKindEnum, 0)
	for _, v := range mappingConditionKind {
		values = append(values, v)
	}
	return values
}
