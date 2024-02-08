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
		common.Logf("Recieved unsupported enum value for Condition: %s.", m.Kind)
		return *m, nil
	}
}

func (m condition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m condition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConditionKindEnum Enum with underlying type: string
type ConditionKindEnum string

// Set of constants representing the allowable values for ConditionKindEnum
const (
	ConditionKindComposite ConditionKindEnum = "COMPOSITE"
	ConditionKindSimple    ConditionKindEnum = "SIMPLE"
)

var mappingConditionKindEnum = map[string]ConditionKindEnum{
	"COMPOSITE": ConditionKindComposite,
	"SIMPLE":    ConditionKindSimple,
}

var mappingConditionKindEnumLowerCase = map[string]ConditionKindEnum{
	"composite": ConditionKindComposite,
	"simple":    ConditionKindSimple,
}

// GetConditionKindEnumValues Enumerates the set of values for ConditionKindEnum
func GetConditionKindEnumValues() []ConditionKindEnum {
	values := make([]ConditionKindEnum, 0)
	for _, v := range mappingConditionKindEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionKindEnumStringValues Enumerates the set of values in String for ConditionKindEnum
func GetConditionKindEnumStringValues() []string {
	return []string{
		"COMPOSITE",
		"SIMPLE",
	}
}

// GetMappingConditionKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionKindEnum(val string) (ConditionKindEnum, bool) {
	enum, ok := mappingConditionKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
