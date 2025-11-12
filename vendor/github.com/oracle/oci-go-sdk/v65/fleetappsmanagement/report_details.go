// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReportDetails fetch reports for FAM based on filter.
type ReportDetails struct {

	// Name of report.
	Name *string `mandatory:"true" json:"name"`

	// Granularity.
	Granularity ReportDetailsGranularityEnum `mandatory:"false" json:"granularity,omitempty"`

	ReportTimeRange *ReportTimeRange `mandatory:"false" json:"reportTimeRange"`

	// Condition.
	Condition ReportDetailsConditionEnum `mandatory:"false" json:"condition,omitempty"`

	// Filters for reports.
	Filters []Filter `mandatory:"false" json:"filters"`

	// order clauses for reports.
	OrderClause []OrderClause `mandatory:"false" json:"orderClause"`
}

func (m ReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReportDetailsGranularityEnum(string(m.Granularity)); !ok && m.Granularity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Granularity: %s. Supported values are: %s.", m.Granularity, strings.Join(GetReportDetailsGranularityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReportDetailsConditionEnum(string(m.Condition)); !ok && m.Condition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Condition: %s. Supported values are: %s.", m.Condition, strings.Join(GetReportDetailsConditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ReportDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Granularity     ReportDetailsGranularityEnum `json:"granularity"`
		ReportTimeRange *ReportTimeRange             `json:"reportTimeRange"`
		Condition       ReportDetailsConditionEnum   `json:"condition"`
		Filters         []filter                     `json:"filters"`
		OrderClause     []OrderClause                `json:"orderClause"`
		Name            *string                      `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Granularity = model.Granularity

	m.ReportTimeRange = model.ReportTimeRange

	m.Condition = model.Condition

	m.Filters = make([]Filter, len(model.Filters))
	for i, n := range model.Filters {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Filters[i] = nn.(Filter)
		} else {
			m.Filters[i] = nil
		}
	}
	m.OrderClause = make([]OrderClause, len(model.OrderClause))
	copy(m.OrderClause, model.OrderClause)
	m.Name = model.Name

	return
}

// ReportDetailsGranularityEnum Enum with underlying type: string
type ReportDetailsGranularityEnum string

// Set of constants representing the allowable values for ReportDetailsGranularityEnum
const (
	ReportDetailsGranularityDay     ReportDetailsGranularityEnum = "DAY"
	ReportDetailsGranularityWeek    ReportDetailsGranularityEnum = "WEEK"
	ReportDetailsGranularityMonth   ReportDetailsGranularityEnum = "MONTH"
	ReportDetailsGranularityQuarter ReportDetailsGranularityEnum = "QUARTER"
	ReportDetailsGranularityYear    ReportDetailsGranularityEnum = "YEAR"
)

var mappingReportDetailsGranularityEnum = map[string]ReportDetailsGranularityEnum{
	"DAY":     ReportDetailsGranularityDay,
	"WEEK":    ReportDetailsGranularityWeek,
	"MONTH":   ReportDetailsGranularityMonth,
	"QUARTER": ReportDetailsGranularityQuarter,
	"YEAR":    ReportDetailsGranularityYear,
}

var mappingReportDetailsGranularityEnumLowerCase = map[string]ReportDetailsGranularityEnum{
	"day":     ReportDetailsGranularityDay,
	"week":    ReportDetailsGranularityWeek,
	"month":   ReportDetailsGranularityMonth,
	"quarter": ReportDetailsGranularityQuarter,
	"year":    ReportDetailsGranularityYear,
}

// GetReportDetailsGranularityEnumValues Enumerates the set of values for ReportDetailsGranularityEnum
func GetReportDetailsGranularityEnumValues() []ReportDetailsGranularityEnum {
	values := make([]ReportDetailsGranularityEnum, 0)
	for _, v := range mappingReportDetailsGranularityEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDetailsGranularityEnumStringValues Enumerates the set of values in String for ReportDetailsGranularityEnum
func GetReportDetailsGranularityEnumStringValues() []string {
	return []string{
		"DAY",
		"WEEK",
		"MONTH",
		"QUARTER",
		"YEAR",
	}
}

// GetMappingReportDetailsGranularityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDetailsGranularityEnum(val string) (ReportDetailsGranularityEnum, bool) {
	enum, ok := mappingReportDetailsGranularityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReportDetailsConditionEnum Enum with underlying type: string
type ReportDetailsConditionEnum string

// Set of constants representing the allowable values for ReportDetailsConditionEnum
const (
	ReportDetailsConditionAll ReportDetailsConditionEnum = "ALL"
	ReportDetailsConditionAny ReportDetailsConditionEnum = "ANY"
)

var mappingReportDetailsConditionEnum = map[string]ReportDetailsConditionEnum{
	"ALL": ReportDetailsConditionAll,
	"ANY": ReportDetailsConditionAny,
}

var mappingReportDetailsConditionEnumLowerCase = map[string]ReportDetailsConditionEnum{
	"all": ReportDetailsConditionAll,
	"any": ReportDetailsConditionAny,
}

// GetReportDetailsConditionEnumValues Enumerates the set of values for ReportDetailsConditionEnum
func GetReportDetailsConditionEnumValues() []ReportDetailsConditionEnum {
	values := make([]ReportDetailsConditionEnum, 0)
	for _, v := range mappingReportDetailsConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDetailsConditionEnumStringValues Enumerates the set of values in String for ReportDetailsConditionEnum
func GetReportDetailsConditionEnumStringValues() []string {
	return []string{
		"ALL",
		"ANY",
	}
}

// GetMappingReportDetailsConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDetailsConditionEnum(val string) (ReportDetailsConditionEnum, bool) {
	enum, ok := mappingReportDetailsConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
