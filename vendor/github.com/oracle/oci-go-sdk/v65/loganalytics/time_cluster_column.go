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

// TimeClusterColumn Column returned by querylanguage TIMECLUSTER command.
type TimeClusterColumn struct {

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

	// Time span between each timestamp in the timeseries datapoints.
	IntervalGap *string `mandatory:"false" json:"intervalGap"`

	// List of timestamps making up the timeseries datapoints.
	Intervals []int64 `mandatory:"false" json:"intervals"`

	// Group by columns specified in the command.
	GroupByColumns []AbstractColumn `mandatory:"false" json:"groupByColumns"`

	// Timeseries clusters identified by the command.
	Clusters map[string]TimeStatsCluster `mandatory:"false" json:"clusters"`

	// List of series data sets for each statistical function specified in the command.
	Series []TimeClusterDataColumn `mandatory:"false" json:"series"`

	// Subsystem column belongs to.
	SubSystem SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m TimeClusterColumn) GetDisplayName() *string {
	return m.DisplayName
}

// GetSubSystem returns SubSystem
func (m TimeClusterColumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

// GetValues returns Values
func (m TimeClusterColumn) GetValues() []FieldValue {
	return m.Values
}

// GetIsListOfValues returns IsListOfValues
func (m TimeClusterColumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

// GetIsMultiValued returns IsMultiValued
func (m TimeClusterColumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

// GetIsCaseSensitive returns IsCaseSensitive
func (m TimeClusterColumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

// GetIsGroupable returns IsGroupable
func (m TimeClusterColumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

// GetIsEvaluable returns IsEvaluable
func (m TimeClusterColumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

// GetValueType returns ValueType
func (m TimeClusterColumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

// GetOriginalDisplayName returns OriginalDisplayName
func (m TimeClusterColumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

// GetInternalName returns InternalName
func (m TimeClusterColumn) GetInternalName() *string {
	return m.InternalName
}

func (m TimeClusterColumn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TimeClusterColumn) ValidateEnumValue() (bool, error) {
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
func (m TimeClusterColumn) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTimeClusterColumn TimeClusterColumn
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTimeClusterColumn
	}{
		"TIME_CLUSTER_COLUMN",
		(MarshalTypeTimeClusterColumn)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TimeClusterColumn) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName         *string                     `json:"displayName"`
		SubSystem           SubSystemNameEnum           `json:"subSystem"`
		Values              []FieldValue                `json:"values"`
		IsListOfValues      *bool                       `json:"isListOfValues"`
		IsMultiValued       *bool                       `json:"isMultiValued"`
		IsCaseSensitive     *bool                       `json:"isCaseSensitive"`
		IsGroupable         *bool                       `json:"isGroupable"`
		IsEvaluable         *bool                       `json:"isEvaluable"`
		ValueType           ValueTypeEnum               `json:"valueType"`
		OriginalDisplayName *string                     `json:"originalDisplayName"`
		InternalName        *string                     `json:"internalName"`
		IntervalGap         *string                     `json:"intervalGap"`
		Intervals           []int64                     `json:"intervals"`
		GroupByColumns      []abstractcolumn            `json:"groupByColumns"`
		Clusters            map[string]TimeStatsCluster `json:"clusters"`
		Series              []TimeClusterDataColumn     `json:"series"`
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

	m.ValueType = model.ValueType

	m.OriginalDisplayName = model.OriginalDisplayName

	m.InternalName = model.InternalName

	m.IntervalGap = model.IntervalGap

	m.Intervals = make([]int64, len(model.Intervals))
	copy(m.Intervals, model.Intervals)
	m.GroupByColumns = make([]AbstractColumn, len(model.GroupByColumns))
	for i, n := range model.GroupByColumns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.GroupByColumns[i] = nn.(AbstractColumn)
		} else {
			m.GroupByColumns[i] = nil
		}
	}
	m.Clusters = model.Clusters

	m.Series = make([]TimeClusterDataColumn, len(model.Series))
	copy(m.Series, model.Series)
	return
}
