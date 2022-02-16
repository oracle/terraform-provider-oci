// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FunctionField Field outlining queryString aggregate function entries.
type FunctionField struct {

	// Field display name - will be alias if field is renamed by queryStrng.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Field denoting if this is a declaration of the field in the queryString.
	IsDeclared *bool `mandatory:"false" json:"isDeclared"`

	// Same as displayName unless field renamed in which case this will hold the original display names for the field
	// across all renames.
	OriginalDisplayNames []string `mandatory:"false" json:"originalDisplayNames"`

	// Internal identifier for the field.
	InternalName *string `mandatory:"false" json:"internalName"`

	// Identifies if this field can be used as a grouping field in any grouping command.
	IsGroupable *bool `mandatory:"false" json:"isGroupable"`

	// Identifies if this field format is a duration.
	IsDuration *bool `mandatory:"false" json:"isDuration"`

	// Alias of field if renamed by queryStrng.
	Alias *string `mandatory:"false" json:"alias"`

	// Query used to derive this field if specified.
	FilterQueryString *string `mandatory:"false" json:"filterQueryString"`

	// Field denoting field unit type.
	UnitType *string `mandatory:"false" json:"unitType"`

	// Name of the aggregate function.
	Function *string `mandatory:"false" json:"function"`

	// List of function arguments if specified.
	Arguments []Argument `mandatory:"false" json:"arguments"`

	// Field denoting field data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

//GetDisplayName returns DisplayName
func (m FunctionField) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsDeclared returns IsDeclared
func (m FunctionField) GetIsDeclared() *bool {
	return m.IsDeclared
}

//GetOriginalDisplayNames returns OriginalDisplayNames
func (m FunctionField) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

//GetInternalName returns InternalName
func (m FunctionField) GetInternalName() *string {
	return m.InternalName
}

//GetValueType returns ValueType
func (m FunctionField) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetIsGroupable returns IsGroupable
func (m FunctionField) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsDuration returns IsDuration
func (m FunctionField) GetIsDuration() *bool {
	return m.IsDuration
}

//GetAlias returns Alias
func (m FunctionField) GetAlias() *string {
	return m.Alias
}

//GetFilterQueryString returns FilterQueryString
func (m FunctionField) GetFilterQueryString() *string {
	return m.FilterQueryString
}

//GetUnitType returns UnitType
func (m FunctionField) GetUnitType() *string {
	return m.UnitType
}

func (m FunctionField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FunctionField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFunctionField FunctionField
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeFunctionField
	}{
		"FUNCTION",
		(MarshalTypeFunctionField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *FunctionField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string       `json:"displayName"`
		IsDeclared           *bool         `json:"isDeclared"`
		OriginalDisplayNames []string      `json:"originalDisplayNames"`
		InternalName         *string       `json:"internalName"`
		ValueType            ValueTypeEnum `json:"valueType"`
		IsGroupable          *bool         `json:"isGroupable"`
		IsDuration           *bool         `json:"isDuration"`
		Alias                *string       `json:"alias"`
		FilterQueryString    *string       `json:"filterQueryString"`
		UnitType             *string       `json:"unitType"`
		Function             *string       `json:"function"`
		Arguments            []argument    `json:"arguments"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.IsDeclared = model.IsDeclared

	m.OriginalDisplayNames = make([]string, len(model.OriginalDisplayNames))
	for i, n := range model.OriginalDisplayNames {
		m.OriginalDisplayNames[i] = n
	}

	m.InternalName = model.InternalName

	m.ValueType = model.ValueType

	m.IsGroupable = model.IsGroupable

	m.IsDuration = model.IsDuration

	m.Alias = model.Alias

	m.FilterQueryString = model.FilterQueryString

	m.UnitType = model.UnitType

	m.Function = model.Function

	m.Arguments = make([]Argument, len(model.Arguments))
	for i, n := range model.Arguments {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Arguments[i] = nn.(Argument)
		} else {
			m.Arguments[i] = nil
		}
	}

	return
}
