// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Field Default field object representing fields specified in the queryString.
type Field struct {

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

	// Field denoting field data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m Field) GetDisplayName() *string {
	return m.DisplayName
}

// GetIsDeclared returns IsDeclared
func (m Field) GetIsDeclared() *bool {
	return m.IsDeclared
}

// GetOriginalDisplayNames returns OriginalDisplayNames
func (m Field) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

// GetInternalName returns InternalName
func (m Field) GetInternalName() *string {
	return m.InternalName
}

// GetValueType returns ValueType
func (m Field) GetValueType() ValueTypeEnum {
	return m.ValueType
}

// GetIsGroupable returns IsGroupable
func (m Field) GetIsGroupable() *bool {
	return m.IsGroupable
}

// GetIsDuration returns IsDuration
func (m Field) GetIsDuration() *bool {
	return m.IsDuration
}

// GetAlias returns Alias
func (m Field) GetAlias() *string {
	return m.Alias
}

// GetFilterQueryString returns FilterQueryString
func (m Field) GetFilterQueryString() *string {
	return m.FilterQueryString
}

// GetUnitType returns UnitType
func (m Field) GetUnitType() *string {
	return m.UnitType
}

func (m Field) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Field) ValidateEnumValue() (bool, error) {
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
func (m Field) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeField Field
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeField
	}{
		"FIELD",
		(MarshalTypeField)(m),
	}

	return json.Marshal(&s)
}
