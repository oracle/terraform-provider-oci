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

// TableColumn Result column that contains a table within each row.
type TableColumn struct {

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

	// Identifies if this column should be hidden by default but can be displayed in the UI on demand.
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// Same as displayName unless column renamed in which case this will hold the original display name for the column.
	OriginalDisplayName *string `mandatory:"false" json:"originalDisplayName"`

	// Internal identifier for the column.
	InternalName *string `mandatory:"false" json:"internalName"`

	// Column descriptors for the table result.
	Columns []AbstractColumn `mandatory:"false" json:"columns"`

	// Results data of the table.
	Result []map[string]interface{} `mandatory:"false" json:"result"`

	// True if query did not complete processing all data.
	ArePartialResults *bool `mandatory:"false" json:"arePartialResults"`

	// Subsystem column belongs to.
	SubSystem SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m TableColumn) GetDisplayName() *string {
	return m.DisplayName
}

// GetSubSystem returns SubSystem
func (m TableColumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

// GetValues returns Values
func (m TableColumn) GetValues() []FieldValue {
	return m.Values
}

// GetIsListOfValues returns IsListOfValues
func (m TableColumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

// GetIsMultiValued returns IsMultiValued
func (m TableColumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

// GetIsCaseSensitive returns IsCaseSensitive
func (m TableColumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

// GetIsGroupable returns IsGroupable
func (m TableColumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

// GetIsEvaluable returns IsEvaluable
func (m TableColumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

// GetIsHidden returns IsHidden
func (m TableColumn) GetIsHidden() *bool {
	return m.IsHidden
}

// GetValueType returns ValueType
func (m TableColumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

// GetOriginalDisplayName returns OriginalDisplayName
func (m TableColumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

// GetInternalName returns InternalName
func (m TableColumn) GetInternalName() *string {
	return m.InternalName
}

func (m TableColumn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableColumn) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubSystemNameEnum(string(m.SubSystem)); !ok && m.SubSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubSystem: %s. Supported values are: %s.", m.SubSystem, strings.Join(GetSubSystemNameEnumStringValues(), ",")))
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
func (m TableColumn) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTableColumn TableColumn
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTableColumn
	}{
		"TABLE_COLUMN",
		(MarshalTypeTableColumn)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TableColumn) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                  `json:"displayName"`
		SubSystem           SubSystemNameEnum        `json:"subSystem"`
		Values              []FieldValue             `json:"values"`
		IsListOfValues      *bool                    `json:"isListOfValues"`
		IsMultiValued       *bool                    `json:"isMultiValued"`
		IsCaseSensitive     *bool                    `json:"isCaseSensitive"`
		IsGroupable         *bool                    `json:"isGroupable"`
		IsEvaluable         *bool                    `json:"isEvaluable"`
		IsHidden            *bool                    `json:"isHidden"`
		ValueType           ValueTypeEnum            `json:"valueType"`
		OriginalDisplayName *string                  `json:"originalDisplayName"`
		InternalName        *string                  `json:"internalName"`
		Columns             []abstractcolumn         `json:"columns"`
		Result              []map[string]interface{} `json:"result"`
		ArePartialResults   *bool                    `json:"arePartialResults"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.SubSystem = model.SubSystem

	m.Values = make([]FieldValue, len(model.Values))
	copy(m.Values, model.Values)
	m.IsListOfValues = model.IsListOfValues

	m.IsMultiValued = model.IsMultiValued

	m.IsCaseSensitive = model.IsCaseSensitive

	m.IsGroupable = model.IsGroupable

	m.IsEvaluable = model.IsEvaluable

	m.IsHidden = model.IsHidden

	m.ValueType = model.ValueType

	m.OriginalDisplayName = model.OriginalDisplayName

	m.InternalName = model.InternalName

	m.Columns = make([]AbstractColumn, len(model.Columns))
	for i, n := range model.Columns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Columns[i] = nn.(AbstractColumn)
		} else {
			m.Columns[i] = nil
		}
	}
	m.Result = make([]map[string]interface{}, len(model.Result))
	copy(m.Result, model.Result)
	m.ArePartialResults = model.ArePartialResults

	return
}
