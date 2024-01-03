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

	// Identifies if this column should be hidden by default but can be displayed in the UI on demand.
	GetIsHidden() *bool

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
	IsHidden            *bool             `mandatory:"false" json:"isHidden"`
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
	m.IsHidden = s.Model.IsHidden
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
	case "TIME_STATS_COLUMN":
		mm := TimeStatsColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_CLUSTER_COLUMN":
		mm := TimeClusterColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COLUMN":
		mm := Column{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_CLUSTER_DATA_COLUMN":
		mm := TimeClusterDataColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_STATS_DATA_COLUMN":
		mm := TimeStatsDataColumn{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TABLE_COLUMN":
		mm := TableColumn{}
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
		common.Logf("Recieved unsupported enum value for AbstractColumn: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m abstractcolumn) GetDisplayName() *string {
	return m.DisplayName
}

// GetSubSystem returns SubSystem
func (m abstractcolumn) GetSubSystem() SubSystemNameEnum {
	return m.SubSystem
}

// GetValues returns Values
func (m abstractcolumn) GetValues() []FieldValue {
	return m.Values
}

// GetIsListOfValues returns IsListOfValues
func (m abstractcolumn) GetIsListOfValues() *bool {
	return m.IsListOfValues
}

// GetIsMultiValued returns IsMultiValued
func (m abstractcolumn) GetIsMultiValued() *bool {
	return m.IsMultiValued
}

// GetIsCaseSensitive returns IsCaseSensitive
func (m abstractcolumn) GetIsCaseSensitive() *bool {
	return m.IsCaseSensitive
}

// GetIsGroupable returns IsGroupable
func (m abstractcolumn) GetIsGroupable() *bool {
	return m.IsGroupable
}

// GetIsEvaluable returns IsEvaluable
func (m abstractcolumn) GetIsEvaluable() *bool {
	return m.IsEvaluable
}

// GetIsHidden returns IsHidden
func (m abstractcolumn) GetIsHidden() *bool {
	return m.IsHidden
}

// GetValueType returns ValueType
func (m abstractcolumn) GetValueType() ValueTypeEnum {
	return m.ValueType
}

// GetOriginalDisplayName returns OriginalDisplayName
func (m abstractcolumn) GetOriginalDisplayName() *string {
	return m.OriginalDisplayName
}

// GetInternalName returns InternalName
func (m abstractcolumn) GetInternalName() *string {
	return m.InternalName
}

func (m abstractcolumn) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractcolumn) ValidateEnumValue() (bool, error) {
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

// AbstractColumnTypeEnum Enum with underlying type: string
type AbstractColumnTypeEnum string

// Set of constants representing the allowable values for AbstractColumnTypeEnum
const (
	AbstractColumnTypeColumn                AbstractColumnTypeEnum = "COLUMN"
	AbstractColumnTypeChartColumn           AbstractColumnTypeEnum = "CHART_COLUMN"
	AbstractColumnTypeChartDataColumn       AbstractColumnTypeEnum = "CHART_DATA_COLUMN"
	AbstractColumnTypeTimeStatsColumn       AbstractColumnTypeEnum = "TIME_STATS_COLUMN"
	AbstractColumnTypeTimeStatsDataColumn   AbstractColumnTypeEnum = "TIME_STATS_DATA_COLUMN"
	AbstractColumnTypeTimeClusterColumn     AbstractColumnTypeEnum = "TIME_CLUSTER_COLUMN"
	AbstractColumnTypeTimeClusterDataColumn AbstractColumnTypeEnum = "TIME_CLUSTER_DATA_COLUMN"
	AbstractColumnTypeTableColumn           AbstractColumnTypeEnum = "TABLE_COLUMN"
	AbstractColumnTypeTimeColumn            AbstractColumnTypeEnum = "TIME_COLUMN"
	AbstractColumnTypeTrendColumn           AbstractColumnTypeEnum = "TREND_COLUMN"
	AbstractColumnTypeClassifyColumn        AbstractColumnTypeEnum = "CLASSIFY_COLUMN"
)

var mappingAbstractColumnTypeEnum = map[string]AbstractColumnTypeEnum{
	"COLUMN":                   AbstractColumnTypeColumn,
	"CHART_COLUMN":             AbstractColumnTypeChartColumn,
	"CHART_DATA_COLUMN":        AbstractColumnTypeChartDataColumn,
	"TIME_STATS_COLUMN":        AbstractColumnTypeTimeStatsColumn,
	"TIME_STATS_DATA_COLUMN":   AbstractColumnTypeTimeStatsDataColumn,
	"TIME_CLUSTER_COLUMN":      AbstractColumnTypeTimeClusterColumn,
	"TIME_CLUSTER_DATA_COLUMN": AbstractColumnTypeTimeClusterDataColumn,
	"TABLE_COLUMN":             AbstractColumnTypeTableColumn,
	"TIME_COLUMN":              AbstractColumnTypeTimeColumn,
	"TREND_COLUMN":             AbstractColumnTypeTrendColumn,
	"CLASSIFY_COLUMN":          AbstractColumnTypeClassifyColumn,
}

var mappingAbstractColumnTypeEnumLowerCase = map[string]AbstractColumnTypeEnum{
	"column":                   AbstractColumnTypeColumn,
	"chart_column":             AbstractColumnTypeChartColumn,
	"chart_data_column":        AbstractColumnTypeChartDataColumn,
	"time_stats_column":        AbstractColumnTypeTimeStatsColumn,
	"time_stats_data_column":   AbstractColumnTypeTimeStatsDataColumn,
	"time_cluster_column":      AbstractColumnTypeTimeClusterColumn,
	"time_cluster_data_column": AbstractColumnTypeTimeClusterDataColumn,
	"table_column":             AbstractColumnTypeTableColumn,
	"time_column":              AbstractColumnTypeTimeColumn,
	"trend_column":             AbstractColumnTypeTrendColumn,
	"classify_column":          AbstractColumnTypeClassifyColumn,
}

// GetAbstractColumnTypeEnumValues Enumerates the set of values for AbstractColumnTypeEnum
func GetAbstractColumnTypeEnumValues() []AbstractColumnTypeEnum {
	values := make([]AbstractColumnTypeEnum, 0)
	for _, v := range mappingAbstractColumnTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractColumnTypeEnumStringValues Enumerates the set of values in String for AbstractColumnTypeEnum
func GetAbstractColumnTypeEnumStringValues() []string {
	return []string{
		"COLUMN",
		"CHART_COLUMN",
		"CHART_DATA_COLUMN",
		"TIME_STATS_COLUMN",
		"TIME_STATS_DATA_COLUMN",
		"TIME_CLUSTER_COLUMN",
		"TIME_CLUSTER_DATA_COLUMN",
		"TABLE_COLUMN",
		"TIME_COLUMN",
		"TREND_COLUMN",
		"CLASSIFY_COLUMN",
	}
}

// GetMappingAbstractColumnTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractColumnTypeEnum(val string) (AbstractColumnTypeEnum, bool) {
	enum, ok := mappingAbstractColumnTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
