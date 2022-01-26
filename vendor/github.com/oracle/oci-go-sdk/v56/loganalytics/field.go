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
	"github.com/oracle/oci-go-sdk/v56/common"
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

//GetDisplayName returns DisplayName
func (m Field) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsDeclared returns IsDeclared
func (m Field) GetIsDeclared() *bool {
	return m.IsDeclared
}

//GetOriginalDisplayNames returns OriginalDisplayNames
func (m Field) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

//GetInternalName returns InternalName
func (m Field) GetInternalName() *string {
	return m.InternalName
}

//GetValueType returns ValueType
func (m Field) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetIsGroupable returns IsGroupable
func (m Field) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsDuration returns IsDuration
func (m Field) GetIsDuration() *bool {
	return m.IsDuration
}

//GetAlias returns Alias
func (m Field) GetAlias() *string {
	return m.Alias
}

//GetFilterQueryString returns FilterQueryString
func (m Field) GetFilterQueryString() *string {
	return m.FilterQueryString
}

//GetUnitType returns UnitType
func (m Field) GetUnitType() *string {
	return m.UnitType
}

func (m Field) String() string {
	return common.PointerString(m)
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
