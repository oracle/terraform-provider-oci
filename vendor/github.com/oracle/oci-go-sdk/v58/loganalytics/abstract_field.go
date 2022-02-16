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

// AbstractField Generic field defining all attributes common to all querylanguage fields.
type AbstractField interface {

	// Field display name - will be alias if field is renamed by queryStrng.
	GetDisplayName() *string

	// Field denoting if this is a declaration of the field in the queryString.
	GetIsDeclared() *bool

	// Same as displayName unless field renamed in which case this will hold the original display names for the field
	// across all renames.
	GetOriginalDisplayNames() []string

	// Internal identifier for the field.
	GetInternalName() *string

	// Field denoting field data type.
	GetValueType() ValueTypeEnum

	// Identifies if this field can be used as a grouping field in any grouping command.
	GetIsGroupable() *bool

	// Identifies if this field format is a duration.
	GetIsDuration() *bool

	// Alias of field if renamed by queryStrng.
	GetAlias() *string

	// Query used to derive this field if specified.
	GetFilterQueryString() *string

	// Field denoting field unit type.
	GetUnitType() *string
}

type abstractfield struct {
	JsonData             []byte
	DisplayName          *string       `mandatory:"false" json:"displayName"`
	IsDeclared           *bool         `mandatory:"false" json:"isDeclared"`
	OriginalDisplayNames []string      `mandatory:"false" json:"originalDisplayNames"`
	InternalName         *string       `mandatory:"false" json:"internalName"`
	ValueType            ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
	IsGroupable          *bool         `mandatory:"false" json:"isGroupable"`
	IsDuration           *bool         `mandatory:"false" json:"isDuration"`
	Alias                *string       `mandatory:"false" json:"alias"`
	FilterQueryString    *string       `mandatory:"false" json:"filterQueryString"`
	UnitType             *string       `mandatory:"false" json:"unitType"`
	Name                 string        `json:"name"`
}

// UnmarshalJSON unmarshals json
func (m *abstractfield) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractfield abstractfield
	s := struct {
		Model Unmarshalerabstractfield
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.IsDeclared = s.Model.IsDeclared
	m.OriginalDisplayNames = s.Model.OriginalDisplayNames
	m.InternalName = s.Model.InternalName
	m.ValueType = s.Model.ValueType
	m.IsGroupable = s.Model.IsGroupable
	m.IsDuration = s.Model.IsDuration
	m.Alias = s.Model.Alias
	m.FilterQueryString = s.Model.FilterQueryString
	m.UnitType = s.Model.UnitType
	m.Name = s.Model.Name

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractfield) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Name {
	case "FIELDS":
		mm := FieldsAddRemoveField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION":
		mm := FunctionField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD":
		mm := Field{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SORT":
		mm := SortField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m abstractfield) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsDeclared returns IsDeclared
func (m abstractfield) GetIsDeclared() *bool {
	return m.IsDeclared
}

//GetOriginalDisplayNames returns OriginalDisplayNames
func (m abstractfield) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

//GetInternalName returns InternalName
func (m abstractfield) GetInternalName() *string {
	return m.InternalName
}

//GetValueType returns ValueType
func (m abstractfield) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetIsGroupable returns IsGroupable
func (m abstractfield) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsDuration returns IsDuration
func (m abstractfield) GetIsDuration() *bool {
	return m.IsDuration
}

//GetAlias returns Alias
func (m abstractfield) GetAlias() *string {
	return m.Alias
}

//GetFilterQueryString returns FilterQueryString
func (m abstractfield) GetFilterQueryString() *string {
	return m.FilterQueryString
}

//GetUnitType returns UnitType
func (m abstractfield) GetUnitType() *string {
	return m.UnitType
}

func (m abstractfield) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractfield) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractFieldNameEnum Enum with underlying type: string
type AbstractFieldNameEnum string

// Set of constants representing the allowable values for AbstractFieldNameEnum
const (
	AbstractFieldNameField    AbstractFieldNameEnum = "FIELD"
	AbstractFieldNameFields   AbstractFieldNameEnum = "FIELDS"
	AbstractFieldNameFunction AbstractFieldNameEnum = "FUNCTION"
	AbstractFieldNameSort     AbstractFieldNameEnum = "SORT"
)

var mappingAbstractFieldNameEnum = map[string]AbstractFieldNameEnum{
	"FIELD":    AbstractFieldNameField,
	"FIELDS":   AbstractFieldNameFields,
	"FUNCTION": AbstractFieldNameFunction,
	"SORT":     AbstractFieldNameSort,
}

// GetAbstractFieldNameEnumValues Enumerates the set of values for AbstractFieldNameEnum
func GetAbstractFieldNameEnumValues() []AbstractFieldNameEnum {
	values := make([]AbstractFieldNameEnum, 0)
	for _, v := range mappingAbstractFieldNameEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractFieldNameEnumStringValues Enumerates the set of values in String for AbstractFieldNameEnum
func GetAbstractFieldNameEnumStringValues() []string {
	return []string{
		"FIELD",
		"FIELDS",
		"FUNCTION",
		"SORT",
	}
}

// GetMappingAbstractFieldNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractFieldNameEnum(val string) (AbstractFieldNameEnum, bool) {
	mappingAbstractFieldNameEnumIgnoreCase := make(map[string]AbstractFieldNameEnum)
	for k, v := range mappingAbstractFieldNameEnum {
		mappingAbstractFieldNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAbstractFieldNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
