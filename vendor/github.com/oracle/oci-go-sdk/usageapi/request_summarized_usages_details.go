// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RequestSummarizedUsagesDetails details for the '/usage' query
type RequestSummarizedUsagesDetails struct {

	// tenant id
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The start time of the usage.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The end time of the usage.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The granularity of the usage.
	// HOURLY - Hourly aggregation of data
	// DAILY - Daily aggregation of data
	// MONTHLY - Monthly aggregation of data
	// TOTAL - Not Supported Yet
	Granularity RequestSummarizedUsagesDetailsGranularityEnum `mandatory:"true" json:"granularity"`

	// The type of query of the usage.
	// Usage - Query the usage data.
	// Cost - Query the cost / billing data.
	QueryType RequestSummarizedUsagesDetailsQueryTypeEnum `mandatory:"false" json:"queryType,omitempty"`

	// Aggregate the result by.
	// example:
	//   `["service"]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// The depth level of the compartment.
	CompartmentDepth *float32 `mandatory:"false" json:"compartmentDepth"`

	Filter *Filter `mandatory:"false" json:"filter"`
}

func (m RequestSummarizedUsagesDetails) String() string {
	return common.PointerString(m)
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

var mappingRequestSummarizedUsagesDetailsGranularity = map[string]RequestSummarizedUsagesDetailsGranularityEnum{
	"HOURLY":  RequestSummarizedUsagesDetailsGranularityHourly,
	"DAILY":   RequestSummarizedUsagesDetailsGranularityDaily,
	"MONTHLY": RequestSummarizedUsagesDetailsGranularityMonthly,
	"TOTAL":   RequestSummarizedUsagesDetailsGranularityTotal,
}

// GetRequestSummarizedUsagesDetailsGranularityEnumValues Enumerates the set of values for RequestSummarizedUsagesDetailsGranularityEnum
func GetRequestSummarizedUsagesDetailsGranularityEnumValues() []RequestSummarizedUsagesDetailsGranularityEnum {
	values := make([]RequestSummarizedUsagesDetailsGranularityEnum, 0)
	for _, v := range mappingRequestSummarizedUsagesDetailsGranularity {
		values = append(values, v)
	}
	return values
}

// RequestSummarizedUsagesDetailsQueryTypeEnum Enum with underlying type: string
type RequestSummarizedUsagesDetailsQueryTypeEnum string

// Set of constants representing the allowable values for RequestSummarizedUsagesDetailsQueryTypeEnum
const (
	RequestSummarizedUsagesDetailsQueryTypeUsage RequestSummarizedUsagesDetailsQueryTypeEnum = "USAGE"
	RequestSummarizedUsagesDetailsQueryTypeCost  RequestSummarizedUsagesDetailsQueryTypeEnum = "COST"
)

var mappingRequestSummarizedUsagesDetailsQueryType = map[string]RequestSummarizedUsagesDetailsQueryTypeEnum{
	"USAGE": RequestSummarizedUsagesDetailsQueryTypeUsage,
	"COST":  RequestSummarizedUsagesDetailsQueryTypeCost,
}

// GetRequestSummarizedUsagesDetailsQueryTypeEnumValues Enumerates the set of values for RequestSummarizedUsagesDetailsQueryTypeEnum
func GetRequestSummarizedUsagesDetailsQueryTypeEnumValues() []RequestSummarizedUsagesDetailsQueryTypeEnum {
	values := make([]RequestSummarizedUsagesDetailsQueryTypeEnum, 0)
	for _, v := range mappingRequestSummarizedUsagesDetailsQueryType {
		values = append(values, v)
	}
	return values
}
