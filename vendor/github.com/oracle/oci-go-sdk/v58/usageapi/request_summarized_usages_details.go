// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RequestSummarizedUsagesDetails Details for the '/usage' query.
type RequestSummarizedUsagesDetails struct {

	// Tenant ID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The usage granularity.
	// HOURLY - Hourly data aggregation.
	// DAILY - Daily data aggregation.
	// MONTHLY - Monthly data aggregation.
	// TOTAL - Not yet supported.
	Granularity RequestSummarizedUsagesDetailsGranularityEnum `mandatory:"true" json:"granularity"`

	// Whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
	IsAggregateByTime *bool `mandatory:"false" json:"isAggregateByTime"`

	Forecast *Forecast `mandatory:"false" json:"forecast"`

	// The query usage type. COST by default if it is missing.
	// Usage - Query the usage data.
	// Cost - Query the cost/billing data.
	// Credit - Query the credit adjustments data.
	// ExpiredCredit - Query the expired credits data.
	// AllCredit - Query the credit adjustments and expired credit.
	QueryType RequestSummarizedUsagesDetailsQueryTypeEnum `mandatory:"false" json:"queryType,omitempty"`

	// Aggregate the result by.
	// example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "tenantId", "tenantName"]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list.
	// For example:
	//   `[{"namespace":"oracle", "key":"createdBy"]`
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	// The compartment depth level.
	CompartmentDepth *float32 `mandatory:"false" json:"compartmentDepth"`

	Filter *Filter `mandatory:"false" json:"filter"`
}

func (m RequestSummarizedUsagesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedUsagesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRequestSummarizedUsagesDetailsGranularityEnum(string(m.Granularity)); !ok && m.Granularity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Granularity: %s. Supported values are: %s.", m.Granularity, strings.Join(GetRequestSummarizedUsagesDetailsGranularityEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRequestSummarizedUsagesDetailsQueryTypeEnum(string(m.QueryType)); !ok && m.QueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryType: %s. Supported values are: %s.", m.QueryType, strings.Join(GetRequestSummarizedUsagesDetailsQueryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedUsagesDetailsGranularityEnum Enum with underlying type: string
type RequestSummarizedUsagesDetailsGranularityEnum string

// Set of constants representing the allowable values for RequestSummarizedUsagesDetailsGranularityEnum
const (
	RequestSummarizedUsagesDetailsGranularityHourly  RequestSummarizedUsagesDetailsGranularityEnum = "HOURLY"
	RequestSummarizedUsagesDetailsGranularityDaily   RequestSummarizedUsagesDetailsGranularityEnum = "DAILY"
	RequestSummarizedUsagesDetailsGranularityMonthly RequestSummarizedUsagesDetailsGranularityEnum = "MONTHLY"
	RequestSummarizedUsagesDetailsGranularityTotal   RequestSummarizedUsagesDetailsGranularityEnum = "TOTAL"
)

var mappingRequestSummarizedUsagesDetailsGranularityEnum = map[string]RequestSummarizedUsagesDetailsGranularityEnum{
	"HOURLY":  RequestSummarizedUsagesDetailsGranularityHourly,
	"DAILY":   RequestSummarizedUsagesDetailsGranularityDaily,
	"MONTHLY": RequestSummarizedUsagesDetailsGranularityMonthly,
	"TOTAL":   RequestSummarizedUsagesDetailsGranularityTotal,
}

// GetRequestSummarizedUsagesDetailsGranularityEnumValues Enumerates the set of values for RequestSummarizedUsagesDetailsGranularityEnum
func GetRequestSummarizedUsagesDetailsGranularityEnumValues() []RequestSummarizedUsagesDetailsGranularityEnum {
	values := make([]RequestSummarizedUsagesDetailsGranularityEnum, 0)
	for _, v := range mappingRequestSummarizedUsagesDetailsGranularityEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedUsagesDetailsGranularityEnumStringValues Enumerates the set of values in String for RequestSummarizedUsagesDetailsGranularityEnum
func GetRequestSummarizedUsagesDetailsGranularityEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"MONTHLY",
		"TOTAL",
	}
}

// GetMappingRequestSummarizedUsagesDetailsGranularityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedUsagesDetailsGranularityEnum(val string) (RequestSummarizedUsagesDetailsGranularityEnum, bool) {
	mappingRequestSummarizedUsagesDetailsGranularityEnumIgnoreCase := make(map[string]RequestSummarizedUsagesDetailsGranularityEnum)
	for k, v := range mappingRequestSummarizedUsagesDetailsGranularityEnum {
		mappingRequestSummarizedUsagesDetailsGranularityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRequestSummarizedUsagesDetailsGranularityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedUsagesDetailsQueryTypeEnum Enum with underlying type: string
type RequestSummarizedUsagesDetailsQueryTypeEnum string

// Set of constants representing the allowable values for RequestSummarizedUsagesDetailsQueryTypeEnum
const (
	RequestSummarizedUsagesDetailsQueryTypeUsage         RequestSummarizedUsagesDetailsQueryTypeEnum = "USAGE"
	RequestSummarizedUsagesDetailsQueryTypeCost          RequestSummarizedUsagesDetailsQueryTypeEnum = "COST"
	RequestSummarizedUsagesDetailsQueryTypeCredit        RequestSummarizedUsagesDetailsQueryTypeEnum = "CREDIT"
	RequestSummarizedUsagesDetailsQueryTypeExpiredcredit RequestSummarizedUsagesDetailsQueryTypeEnum = "EXPIREDCREDIT"
	RequestSummarizedUsagesDetailsQueryTypeAllcredit     RequestSummarizedUsagesDetailsQueryTypeEnum = "ALLCREDIT"
)

var mappingRequestSummarizedUsagesDetailsQueryTypeEnum = map[string]RequestSummarizedUsagesDetailsQueryTypeEnum{
	"USAGE":         RequestSummarizedUsagesDetailsQueryTypeUsage,
	"COST":          RequestSummarizedUsagesDetailsQueryTypeCost,
	"CREDIT":        RequestSummarizedUsagesDetailsQueryTypeCredit,
	"EXPIREDCREDIT": RequestSummarizedUsagesDetailsQueryTypeExpiredcredit,
	"ALLCREDIT":     RequestSummarizedUsagesDetailsQueryTypeAllcredit,
}

// GetRequestSummarizedUsagesDetailsQueryTypeEnumValues Enumerates the set of values for RequestSummarizedUsagesDetailsQueryTypeEnum
func GetRequestSummarizedUsagesDetailsQueryTypeEnumValues() []RequestSummarizedUsagesDetailsQueryTypeEnum {
	values := make([]RequestSummarizedUsagesDetailsQueryTypeEnum, 0)
	for _, v := range mappingRequestSummarizedUsagesDetailsQueryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedUsagesDetailsQueryTypeEnumStringValues Enumerates the set of values in String for RequestSummarizedUsagesDetailsQueryTypeEnum
func GetRequestSummarizedUsagesDetailsQueryTypeEnumStringValues() []string {
	return []string{
		"USAGE",
		"COST",
		"CREDIT",
		"EXPIREDCREDIT",
		"ALLCREDIT",
	}
}

// GetMappingRequestSummarizedUsagesDetailsQueryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedUsagesDetailsQueryTypeEnum(val string) (RequestSummarizedUsagesDetailsQueryTypeEnum, bool) {
	mappingRequestSummarizedUsagesDetailsQueryTypeEnumIgnoreCase := make(map[string]RequestSummarizedUsagesDetailsQueryTypeEnum)
	for k, v := range mappingRequestSummarizedUsagesDetailsQueryTypeEnum {
		mappingRequestSummarizedUsagesDetailsQueryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRequestSummarizedUsagesDetailsQueryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
