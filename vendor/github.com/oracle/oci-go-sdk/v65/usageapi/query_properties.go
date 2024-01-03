// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryProperties The query properties.
type QueryProperties struct {

	// The usage granularity. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation.
	// Allowed values are:
	//   DAILY
	//   MONTHLY
	Granularity QueryPropertiesGranularityEnum `mandatory:"true" json:"granularity"`

	DateRange DateRange `mandatory:"true" json:"dateRange"`

	// Aggregate the result by. For example: [ "tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName" ]
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: [ { "namespace": "oracle", "key": "createdBy" ]
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	Filter *Filter `mandatory:"false" json:"filter"`

	// The depth level of the compartment.
	CompartmentDepth *float32 `mandatory:"false" json:"compartmentDepth"`

	// The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data.
	// Allowed values are:
	//   USAGE
	//   COST
	//   USAGE_AND_COST
	QueryType QueryPropertiesQueryTypeEnum `mandatory:"false" json:"queryType,omitempty"`

	// Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
	IsAggregateByTime *bool `mandatory:"false" json:"isAggregateByTime"`
}

func (m QueryProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryPropertiesGranularityEnum(string(m.Granularity)); !ok && m.Granularity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Granularity: %s. Supported values are: %s.", m.Granularity, strings.Join(GetQueryPropertiesGranularityEnumStringValues(), ",")))
	}

	if _, ok := GetMappingQueryPropertiesQueryTypeEnum(string(m.QueryType)); !ok && m.QueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryType: %s. Supported values are: %s.", m.QueryType, strings.Join(GetQueryPropertiesQueryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *QueryProperties) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		GroupBy           []string                       `json:"groupBy"`
		GroupByTag        []Tag                          `json:"groupByTag"`
		Filter            *Filter                        `json:"filter"`
		CompartmentDepth  *float32                       `json:"compartmentDepth"`
		QueryType         QueryPropertiesQueryTypeEnum   `json:"queryType"`
		IsAggregateByTime *bool                          `json:"isAggregateByTime"`
		Granularity       QueryPropertiesGranularityEnum `json:"granularity"`
		DateRange         daterange                      `json:"dateRange"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.GroupBy = make([]string, len(model.GroupBy))
	copy(m.GroupBy, model.GroupBy)
	m.GroupByTag = make([]Tag, len(model.GroupByTag))
	copy(m.GroupByTag, model.GroupByTag)
	m.Filter = model.Filter

	m.CompartmentDepth = model.CompartmentDepth

	m.QueryType = model.QueryType

	m.IsAggregateByTime = model.IsAggregateByTime

	m.Granularity = model.Granularity

	nn, e = model.DateRange.UnmarshalPolymorphicJSON(model.DateRange.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DateRange = nn.(DateRange)
	} else {
		m.DateRange = nil
	}

	return
}

// QueryPropertiesGranularityEnum Enum with underlying type: string
type QueryPropertiesGranularityEnum string

// Set of constants representing the allowable values for QueryPropertiesGranularityEnum
const (
	QueryPropertiesGranularityDaily   QueryPropertiesGranularityEnum = "DAILY"
	QueryPropertiesGranularityMonthly QueryPropertiesGranularityEnum = "MONTHLY"
)

var mappingQueryPropertiesGranularityEnum = map[string]QueryPropertiesGranularityEnum{
	"DAILY":   QueryPropertiesGranularityDaily,
	"MONTHLY": QueryPropertiesGranularityMonthly,
}

var mappingQueryPropertiesGranularityEnumLowerCase = map[string]QueryPropertiesGranularityEnum{
	"daily":   QueryPropertiesGranularityDaily,
	"monthly": QueryPropertiesGranularityMonthly,
}

// GetQueryPropertiesGranularityEnumValues Enumerates the set of values for QueryPropertiesGranularityEnum
func GetQueryPropertiesGranularityEnumValues() []QueryPropertiesGranularityEnum {
	values := make([]QueryPropertiesGranularityEnum, 0)
	for _, v := range mappingQueryPropertiesGranularityEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryPropertiesGranularityEnumStringValues Enumerates the set of values in String for QueryPropertiesGranularityEnum
func GetQueryPropertiesGranularityEnumStringValues() []string {
	return []string{
		"DAILY",
		"MONTHLY",
	}
}

// GetMappingQueryPropertiesGranularityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryPropertiesGranularityEnum(val string) (QueryPropertiesGranularityEnum, bool) {
	enum, ok := mappingQueryPropertiesGranularityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QueryPropertiesQueryTypeEnum Enum with underlying type: string
type QueryPropertiesQueryTypeEnum string

// Set of constants representing the allowable values for QueryPropertiesQueryTypeEnum
const (
	QueryPropertiesQueryTypeUsage        QueryPropertiesQueryTypeEnum = "USAGE"
	QueryPropertiesQueryTypeCost         QueryPropertiesQueryTypeEnum = "COST"
	QueryPropertiesQueryTypeUsageAndCost QueryPropertiesQueryTypeEnum = "USAGE_AND_COST"
)

var mappingQueryPropertiesQueryTypeEnum = map[string]QueryPropertiesQueryTypeEnum{
	"USAGE":          QueryPropertiesQueryTypeUsage,
	"COST":           QueryPropertiesQueryTypeCost,
	"USAGE_AND_COST": QueryPropertiesQueryTypeUsageAndCost,
}

var mappingQueryPropertiesQueryTypeEnumLowerCase = map[string]QueryPropertiesQueryTypeEnum{
	"usage":          QueryPropertiesQueryTypeUsage,
	"cost":           QueryPropertiesQueryTypeCost,
	"usage_and_cost": QueryPropertiesQueryTypeUsageAndCost,
}

// GetQueryPropertiesQueryTypeEnumValues Enumerates the set of values for QueryPropertiesQueryTypeEnum
func GetQueryPropertiesQueryTypeEnumValues() []QueryPropertiesQueryTypeEnum {
	values := make([]QueryPropertiesQueryTypeEnum, 0)
	for _, v := range mappingQueryPropertiesQueryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryPropertiesQueryTypeEnumStringValues Enumerates the set of values in String for QueryPropertiesQueryTypeEnum
func GetQueryPropertiesQueryTypeEnumStringValues() []string {
	return []string{
		"USAGE",
		"COST",
		"USAGE_AND_COST",
	}
}

// GetMappingQueryPropertiesQueryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryPropertiesQueryTypeEnum(val string) (QueryPropertiesQueryTypeEnum, bool) {
	enum, ok := mappingQueryPropertiesQueryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
