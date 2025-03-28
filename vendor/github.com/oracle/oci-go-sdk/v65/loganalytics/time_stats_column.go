// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TimeStatsColumn Column returned by querylanguage TIMESTATS command.
type TimeStatsColumn struct {

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

	// Time span between each timestamp in the timeseries datapoints.
	IntervalGap *string `mandatory:"false" json:"intervalGap"`

	// List of timestamps making up the timeseries datapoints.
	Intervals []int64 `mandatory:"false" json:"intervals"`

	// List of series data sets for each statistical function specified in the command.
	Series []TimeStatsDataColumn `mandatory:"false" json:"series"`

	// Subsystem column belongs to.
	SubSystem SubSystemNameEnum `mandatory:"false" json:"subSystem,omitempty"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m TimeStatsColumn) GetDisplayName() *string {
	return m.DisplayName
}

// GetSubSystem returns SubSystem
func (m TimeStatsColumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

// GetValues returns Values
func (m TimeStatsColumn) GetValues() []FieldValue {
	return m.Values
}

// GetIsListOfValues returns IsListOfValues
func (m TimeStatsColumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

// GetIsMultiValued returns IsMultiValued
func (m TimeStatsColumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

// GetIsCaseSensitive returns IsCaseSensitive
func (m TimeStatsColumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

// GetIsGroupable returns IsGroupable
func (m TimeStatsColumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

// GetIsEvaluable returns IsEvaluable
func (m TimeStatsColumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

// GetIsHidden returns IsHidden
func (m TimeStatsColumn) GetIsHidden() *bool {
	return m.IsHidden
}

// GetValueType returns ValueType
func (m TimeStatsColumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

// GetOriginalDisplayName returns OriginalDisplayName
func (m TimeStatsColumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

// GetInternalName returns InternalName
func (m TimeStatsColumn) GetInternalName() *string {
	return m.InternalName
}

func (m TimeStatsColumn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TimeStatsColumn) ValidateEnumValue() (bool, error) {
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
func (m TimeStatsColumn) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTimeStatsColumn TimeStatsColumn
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTimeStatsColumn
	}{
		"TIME_STATS_COLUMN",
		(MarshalTypeTimeStatsColumn)(m),
	}

	return json.Marshal(&s)
}
