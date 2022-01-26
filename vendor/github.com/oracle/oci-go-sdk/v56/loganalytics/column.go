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

// Column Default column object representing querylanguage result columns.
type Column struct {

	// Column display name - will be alias if column is renamed by queryStrng.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If the column is a 'List of Values' column, this array contains the field values that are applicable to query results or all if no filters applied.
	Values []FieldValue `mandatory:"false" json:"values"`

	// Identifies if all values in this column come from a pre-defined list of values.
	IsListOfValues *bool `mandatory:"false" json:"isListOfValues"`

	// Identifies if this column allows multiple values to exist in a single row.
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// A flag indicating whether or not the field is a case sensitive field.  Only applies to string fields.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// Identifies if this column can be used as a grouping field in any grouping command.
	IsGroupable *bool `mandatory:"false" json:"isGroupable"`

	// Identifies if this column can be used as an expression parameter in any command that accepts querylanguage expressions.
	IsEvaluable *bool `mandatory:"false" json:"isEvaluable"`

	// Same as displayName unless column renamed in which case this will hold the original display name for the column.
	OriginalDisplayName *string `mandatory:"false" json:"originalDisplayName"`

	// Internal identifier for the column.
	InternalName *string `mandatory:"false" json:"internalName"`

	// Subsystem column belongs to.
	SubSystem SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

//GetDisplayName returns DisplayName
func (m Column) GetDisplayName() *string {
	return m.DisplayName
}

//GetSubSystem returns SubSystem
func (m Column) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

//GetValues returns Values
func (m Column) GetValues() []FieldValue {
	return m.Values
}

//GetIsListOfValues returns IsListOfValues
func (m Column) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

//GetIsMultiValued returns IsMultiValued
func (m Column) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

//GetIsCaseSensitive returns IsCaseSensitive
func (m Column) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

//GetIsGroupable returns IsGroupable
func (m Column) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsEvaluable returns IsEvaluable
func (m Column) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

//GetValueType returns ValueType
func (m Column) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetOriginalDisplayName returns OriginalDisplayName
func (m Column) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

//GetInternalName returns InternalName
func (m Column) GetInternalName() *string {
	return m.InternalName
}

func (m Column) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Column) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeColumn Column
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeColumn
	}{
		"COLUMN",
		(MarshalTypeColumn)(m),
	}

	return json.Marshal(&s)
}
