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

// FieldsAddRemoveField Field denoting a field specified in querylanguage FIELDS command.
type FieldsAddRemoveField struct {

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

	// Denotes if field entry in FIELDS command is to show / hide field in results.
	Operation FieldsAddRemoveFieldOperationEnum `mandatory:"false" json:"operation,omitempty"`

	// Field denoting field data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

//GetDisplayName returns DisplayName
func (m FieldsAddRemoveField) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsDeclared returns IsDeclared
func (m FieldsAddRemoveField) GetIsDeclared() *bool {
	return m.IsDeclared
}

//GetOriginalDisplayNames returns OriginalDisplayNames
func (m FieldsAddRemoveField) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

//GetInternalName returns InternalName
func (m FieldsAddRemoveField) GetInternalName() *string {
	return m.InternalName
}

//GetValueType returns ValueType
func (m FieldsAddRemoveField) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetIsGroupable returns IsGroupable
func (m FieldsAddRemoveField) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsDuration returns IsDuration
func (m FieldsAddRemoveField) GetIsDuration() *bool {
	return m.IsDuration
}

//GetAlias returns Alias
func (m FieldsAddRemoveField) GetAlias() *string {
	return m.Alias
}

//GetFilterQueryString returns FilterQueryString
func (m FieldsAddRemoveField) GetFilterQueryString() *string {
	return m.FilterQueryString
}

//GetUnitType returns UnitType
func (m FieldsAddRemoveField) GetUnitType() *string {
	return m.UnitType
}

func (m FieldsAddRemoveField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FieldsAddRemoveField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFieldsAddRemoveFieldOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetFieldsAddRemoveFieldOperationEnumStringValues(), ",")))
	}

	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FieldsAddRemoveField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFieldsAddRemoveField FieldsAddRemoveField
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeFieldsAddRemoveField
	}{
		"FIELDS",
		(MarshalTypeFieldsAddRemoveField)(m),
	}

	return json.Marshal(&s)
}

// FieldsAddRemoveFieldOperationEnum Enum with underlying type: string
type FieldsAddRemoveFieldOperationEnum string

// Set of constants representing the allowable values for FieldsAddRemoveFieldOperationEnum
const (
	FieldsAddRemoveFieldOperationAdd    FieldsAddRemoveFieldOperationEnum = "ADD"
	FieldsAddRemoveFieldOperationRemove FieldsAddRemoveFieldOperationEnum = "REMOVE"
)

var mappingFieldsAddRemoveFieldOperationEnum = map[string]FieldsAddRemoveFieldOperationEnum{
	"ADD":    FieldsAddRemoveFieldOperationAdd,
	"REMOVE": FieldsAddRemoveFieldOperationRemove,
}

// GetFieldsAddRemoveFieldOperationEnumValues Enumerates the set of values for FieldsAddRemoveFieldOperationEnum
func GetFieldsAddRemoveFieldOperationEnumValues() []FieldsAddRemoveFieldOperationEnum {
	values := make([]FieldsAddRemoveFieldOperationEnum, 0)
	for _, v := range mappingFieldsAddRemoveFieldOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetFieldsAddRemoveFieldOperationEnumStringValues Enumerates the set of values in String for FieldsAddRemoveFieldOperationEnum
func GetFieldsAddRemoveFieldOperationEnumStringValues() []string {
	return []string{
		"ADD",
		"REMOVE",
	}
}

// GetMappingFieldsAddRemoveFieldOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFieldsAddRemoveFieldOperationEnum(val string) (FieldsAddRemoveFieldOperationEnum, bool) {
	mappingFieldsAddRemoveFieldOperationEnumIgnoreCase := make(map[string]FieldsAddRemoveFieldOperationEnum)
	for k, v := range mappingFieldsAddRemoveFieldOperationEnum {
		mappingFieldsAddRemoveFieldOperationEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFieldsAddRemoveFieldOperationEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
