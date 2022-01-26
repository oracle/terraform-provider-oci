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

// ClassifyColumn Column containing query results produced by the query language classify command.
type ClassifyColumn struct {

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

	// A list of fields specified in the classify command in the query string.
	ClassifyFieldNames []string `mandatory:"false" json:"classifyFieldNames"`

	// Count of nulls found in each of the fields specified in the classify command in the query string.
	ClassifyFieldNullCount []int64 `mandatory:"false" json:"classifyFieldNullCount"`

	// Count of anomalies for each timeseries datapoint.
	ClassifyAnomalyIntervalCounts []int64 `mandatory:"false" json:"classifyAnomalyIntervalCounts"`

	// Column descriptors for the classify result.
	ClassifyColumns []AbstractColumn `mandatory:"false" json:"classifyColumns"`

	// Results of the classify command.
	ClassifyResult []map[string]interface{} `mandatory:"false" json:"classifyResult"`

	// Column descriptors of fields with strong correlation with the classify fields.
	ClassifyCorrelateColumns []AbstractColumn `mandatory:"false" json:"classifyCorrelateColumns"`

	// Correlation results of the classify command.
	ClassifyCorrelateResult []map[string]interface{} `mandatory:"false" json:"classifyCorrelateResult"`

	// Subsystem column belongs to.
	SubSystem SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

//GetDisplayName returns DisplayName
func (m ClassifyColumn) GetDisplayName() *string {
	return m.DisplayName
}

//GetSubSystem returns SubSystem
func (m ClassifyColumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

//GetValues returns Values
func (m ClassifyColumn) GetValues() []FieldValue {
	return m.Values
}

//GetIsListOfValues returns IsListOfValues
func (m ClassifyColumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

//GetIsMultiValued returns IsMultiValued
func (m ClassifyColumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

//GetIsCaseSensitive returns IsCaseSensitive
func (m ClassifyColumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

//GetIsGroupable returns IsGroupable
func (m ClassifyColumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsEvaluable returns IsEvaluable
func (m ClassifyColumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

//GetValueType returns ValueType
func (m ClassifyColumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetOriginalDisplayName returns OriginalDisplayName
func (m ClassifyColumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

//GetInternalName returns InternalName
func (m ClassifyColumn) GetInternalName() *string {
	return m.InternalName
}

func (m ClassifyColumn) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ClassifyColumn) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClassifyColumn ClassifyColumn
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeClassifyColumn
	}{
		"CLASSIFY_COLUMN",
		(MarshalTypeClassifyColumn)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ClassifyColumn) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                   *string                  `json:"displayName"`
		SubSystem                     SubSystemNameEnum        `json:"subSystem"`
		Values                        []FieldValue             `json:"values"`
		IsListOfValues                *bool                    `json:"isListOfValues"`
		IsMultiValued                 *bool                    `json:"isMultiValued"`
		IsCaseSensitive               *bool                    `json:"isCaseSensitive"`
		IsGroupable                   *bool                    `json:"isGroupable"`
		IsEvaluable                   *bool                    `json:"isEvaluable"`
		ValueType                     ValueTypeEnum            `json:"valueType"`
		OriginalDisplayName           *string                  `json:"originalDisplayName"`
		InternalName                  *string                  `json:"internalName"`
		ClassifyFieldNames            []string                 `json:"classifyFieldNames"`
		ClassifyFieldNullCount        []int64                  `json:"classifyFieldNullCount"`
		ClassifyAnomalyIntervalCounts []int64                  `json:"classifyAnomalyIntervalCounts"`
		ClassifyColumns               []abstractcolumn         `json:"classifyColumns"`
		ClassifyResult                []map[string]interface{} `json:"classifyResult"`
		ClassifyCorrelateColumns      []abstractcolumn         `json:"classifyCorrelateColumns"`
		ClassifyCorrelateResult       []map[string]interface{} `json:"classifyCorrelateResult"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.SubSystem = model.SubSystem

	m.Values = make([]FieldValue, len(model.Values))
	for i, n := range model.Values {
		m.Values[i] = n
	}

	m.IsListOfValues = model.IsListOfValues

	m.IsMultiValued = model.IsMultiValued

	m.IsCaseSensitive = model.IsCaseSensitive

	m.IsGroupable = model.IsGroupable

	m.IsEvaluable = model.IsEvaluable

	m.ValueType = model.ValueType

	m.OriginalDisplayName = model.OriginalDisplayName

	m.InternalName = model.InternalName

	m.ClassifyFieldNames = make([]string, len(model.ClassifyFieldNames))
	for i, n := range model.ClassifyFieldNames {
		m.ClassifyFieldNames[i] = n
	}

	m.ClassifyFieldNullCount = make([]int64, len(model.ClassifyFieldNullCount))
	for i, n := range model.ClassifyFieldNullCount {
		m.ClassifyFieldNullCount[i] = n
	}

	m.ClassifyAnomalyIntervalCounts = make([]int64, len(model.ClassifyAnomalyIntervalCounts))
	for i, n := range model.ClassifyAnomalyIntervalCounts {
		m.ClassifyAnomalyIntervalCounts[i] = n
	}

	m.ClassifyColumns = make([]AbstractColumn, len(model.ClassifyColumns))
	for i, n := range model.ClassifyColumns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ClassifyColumns[i] = nn.(AbstractColumn)
		} else {
			m.ClassifyColumns[i] = nil
		}
	}

	m.ClassifyResult = make([]map[string]interface{}, len(model.ClassifyResult))
	for i, n := range model.ClassifyResult {
		m.ClassifyResult[i] = n
	}

	m.ClassifyCorrelateColumns = make([]AbstractColumn, len(model.ClassifyCorrelateColumns))
	for i, n := range model.ClassifyCorrelateColumns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ClassifyCorrelateColumns[i] = nn.(AbstractColumn)
		} else {
			m.ClassifyCorrelateColumns[i] = nil
		}
	}

	m.ClassifyCorrelateResult = make([]map[string]interface{}, len(model.ClassifyCorrelateResult))
	for i, n := range model.ClassifyCorrelateResult {
		m.ClassifyCorrelateResult[i] = n
	}

	return
}
