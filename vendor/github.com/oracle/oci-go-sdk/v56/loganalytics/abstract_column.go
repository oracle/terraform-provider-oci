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

// AbstractColumn Generic column defining all attributes common to all querylanguage columns.
type AbstractColumn interface {

	// Column display name - will be alias if column is renamed by queryStrng.
	GetDisplayName() *string

	// Subsystem column belongs to.
	GetSubSystem() SubSystemNameEnum

	// If the column is a 'List of Values' column, this array contains the field values that are applicable to query results or all if no filters applied.
	GetValues() []FieldValue

	// Identifies if all values in this column come from a pre-defined list of values.
	GetIsListOfValues() *bool

	// Identifies if this column allows multiple values to exist in a single row.
	GetIsMultiValued() *bool

	// A flag indicating whether or not the field is a case sensitive field.  Only applies to string fields.
	GetIsCaseSensitive() *bool

	// Identifies if this column can be used as a grouping field in any grouping command.
	GetIsGroupable() *bool

	// Identifies if this column can be used as an expression parameter in any command that accepts querylanguage expressions.
	GetIsEvaluable() *bool

	// Field denoting column data type.
	GetValueType() ValueTypeEnum

	// Same as displayName unless column renamed in which case this will hold the original display name for the column.
	GetOriginalDisplayName() *string

	// Internal identifier for the column.
	GetInternalName() *string
}

type abstractcolumn struct {
	JsonData            []byte
	DisplayName         *string           `mandatory:"false" json:"displayName"`
	SubSystem           SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`
	Values              []FieldValue      `mandatory:"false" json:"values"`
	IsListOfValues      *bool             `mandatory:"false" json:"isListOfValues"`
	IsMultiValued       *bool             `mandatory:"false" json:"isMultiValued"`
	IsCaseSensitive     *bool             `mandatory:"false" json:"isCaseSensitive"`
	IsGroupable         *bool             `mandatory:"false" json:"isGroupable"`
	IsEvaluable         *bool             `mandatory:"false" json:"isEvaluable"`
	ValueType           ValueTypeEnum     `mandatory:"false" json:"valueType,omitempty"`
	OriginalDisplayName *string           `mandatory:"false" json:"originalDisplayName"`
	InternalName        *string           `mandatory:"false" json:"internalName"`
	Type                string            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *abstractcolumn) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractcolumn abstractcolumn
	s := struct {
		Model Unmarshalerabstractcolumn
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.SubSystem = s.Model.SubSystem
	m.Values = s.Model.Values
	m.IsListOfValues = s.Model.IsListOfValues
	m.IsMultiValued = s.Model.IsMultiValued
	m.IsCaseSensitive = s.Model.IsCaseSensitive
	m.IsGroupable = s.Model.IsGroupable
	m.IsEvaluable = s.Model.IsEvaluable
	m.ValueType = s.Model.ValueType
	m.OriginalDisplayName = s.Model.OriginalDisplayName
	m.InternalName = s.Model.InternalName
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractcolumn) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TIME_COLUMN":
		mm := TimeColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLASSIFY_COLUMN":
		mm := ClassifyColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TREND_COLUMN":
		mm := TrendColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COLUMN":
		mm := Column{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CHART_COLUMN":
		mm := ChartColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CHART_DATA_COLUMN":
		mm := ChartDataColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m abstractcolumn) GetDisplayName() *string {
	return m.DisplayName
}

//GetSubSystem returns SubSystem
func (m abstractcolumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

//GetValues returns Values
func (m abstractcolumn) GetValues() []FieldValue {
	return m.Values
}

//GetIsListOfValues returns IsListOfValues
func (m abstractcolumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

//GetIsMultiValued returns IsMultiValued
func (m abstractcolumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

//GetIsCaseSensitive returns IsCaseSensitive
func (m abstractcolumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

//GetIsGroupable returns IsGroupable
func (m abstractcolumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsEvaluable returns IsEvaluable
func (m abstractcolumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

//GetValueType returns ValueType
func (m abstractcolumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetOriginalDisplayName returns OriginalDisplayName
func (m abstractcolumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

//GetInternalName returns InternalName
func (m abstractcolumn) GetInternalName() *string {
	return m.InternalName
}

func (m abstractcolumn) String() string {
	return common.PointerString(m)
}

// AbstractColumnTypeEnum Enum with underlying type: string
type AbstractColumnTypeEnum string

// Set of constants representing the allowable values for AbstractColumnTypeEnum
const (
	AbstractColumnTypeColumn          AbstractColumnTypeEnum = "COLUMN"
	AbstractColumnTypeChartColumn     AbstractColumnTypeEnum = "CHART_COLUMN"
	AbstractColumnTypeChartDataColumn AbstractColumnTypeEnum = "CHART_DATA_COLUMN"
	AbstractColumnTypeTimeColumn      AbstractColumnTypeEnum = "TIME_COLUMN"
	AbstractColumnTypeTrendColumn     AbstractColumnTypeEnum = "TREND_COLUMN"
	AbstractColumnTypeClassifyColumn  AbstractColumnTypeEnum = "CLASSIFY_COLUMN"
)

var mappingAbstractColumnType = map[string]AbstractColumnTypeEnum{
	"COLUMN":            AbstractColumnTypeColumn,
	"CHART_COLUMN":      AbstractColumnTypeChartColumn,
	"CHART_DATA_COLUMN": AbstractColumnTypeChartDataColumn,
	"TIME_COLUMN":       AbstractColumnTypeTimeColumn,
	"TREND_COLUMN":      AbstractColumnTypeTrendColumn,
	"CLASSIFY_COLUMN":   AbstractColumnTypeClassifyColumn,
}

// GetAbstractColumnTypeEnumValues Enumerates the set of values for AbstractColumnTypeEnum
func GetAbstractColumnTypeEnumValues() []AbstractColumnTypeEnum {
	values := make([]AbstractColumnTypeEnum, 0)
	for _, v := range mappingAbstractColumnType {
		values = append(values, v)
	}
	return values
}
