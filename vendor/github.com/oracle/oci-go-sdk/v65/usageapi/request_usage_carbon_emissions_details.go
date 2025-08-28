// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by Cost Analysis (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm), Scheduled Reports (https://docs.oracle.com/iaas/Content/Billing/Concepts/scheduledreportoverview.htm), and Carbon Emissions Analysis (https://docs.oracle.com/iaas/Content/General/Concepts/emissions-management.htm) in the Console. Also see Using the Usage API (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestUsageCarbonEmissionsDetails Details for the '/usageCarbonEmissions' query.
type RequestUsageCarbonEmissionsDetails struct {

	// Tenant ID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// Specifies the method used for emission calculation, such as POWER_BASED or SPEND_BASED
	EmissionCalculationMethod RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum `mandatory:"false" json:"emissionCalculationMethod,omitempty"`

	// Specifies the type of emission, such as MARKET_BASED or LOCATION_BASED.
	EmissionType RequestUsageCarbonEmissionsDetailsEmissionTypeEnum `mandatory:"false" json:"emissionType,omitempty"`

	// The carbon emission granularity. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation.
	Granularity RequestUsageCarbonEmissionsDetailsGranularityEnum `mandatory:"false" json:"granularity,omitempty"`

	// Specifies whether aggregated by time. If isAggregateByTime is true, all carbon emissions usage over the query time period are summed.
	IsAggregateByTime *bool `mandatory:"false" json:"isAggregateByTime"`

	// Aggregate the result by.
	// For example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "resourceName", "tenantId", "tenantName", "subscriptionId"]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list.
	// For example:
	//   `[{"namespace":"oracle", "key":"createdBy"]`
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	// The compartment depth level.
	CompartmentDepth *int `mandatory:"false" json:"compartmentDepth"`

	Filter *Filter `mandatory:"false" json:"filter"`
}

func (m RequestUsageCarbonEmissionsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestUsageCarbonEmissionsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum(string(m.EmissionCalculationMethod)); !ok && m.EmissionCalculationMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EmissionCalculationMethod: %s. Supported values are: %s.", m.EmissionCalculationMethod, strings.Join(GetRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum(string(m.EmissionType)); !ok && m.EmissionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EmissionType: %s. Supported values are: %s.", m.EmissionType, strings.Join(GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestUsageCarbonEmissionsDetailsGranularityEnum(string(m.Granularity)); !ok && m.Granularity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Granularity: %s. Supported values are: %s.", m.Granularity, strings.Join(GetRequestUsageCarbonEmissionsDetailsGranularityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum Enum with underlying type: string
type RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum string

// Set of constants representing the allowable values for RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum
const (
	RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodSpendBased RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum = "SPEND_BASED"
	RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodPowerBased RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum = "POWER_BASED"
)

var mappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum = map[string]RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum{
	"SPEND_BASED": RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodSpendBased,
	"POWER_BASED": RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodPowerBased,
}

var mappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumLowerCase = map[string]RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum{
	"spend_based": RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodSpendBased,
	"power_based": RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodPowerBased,
}

// GetRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumValues Enumerates the set of values for RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum
func GetRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumValues() []RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum {
	values := make([]RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum, 0)
	for _, v := range mappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumStringValues Enumerates the set of values in String for RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum
func GetRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumStringValues() []string {
	return []string{
		"SPEND_BASED",
		"POWER_BASED",
	}
}

// GetMappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum(val string) (RequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnum, bool) {
	enum, ok := mappingRequestUsageCarbonEmissionsDetailsEmissionCalculationMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestUsageCarbonEmissionsDetailsEmissionTypeEnum Enum with underlying type: string
type RequestUsageCarbonEmissionsDetailsEmissionTypeEnum string

// Set of constants representing the allowable values for RequestUsageCarbonEmissionsDetailsEmissionTypeEnum
const (
	RequestUsageCarbonEmissionsDetailsEmissionTypeMarketBased   RequestUsageCarbonEmissionsDetailsEmissionTypeEnum = "MARKET_BASED"
	RequestUsageCarbonEmissionsDetailsEmissionTypeLocationBased RequestUsageCarbonEmissionsDetailsEmissionTypeEnum = "LOCATION_BASED"
)

var mappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum = map[string]RequestUsageCarbonEmissionsDetailsEmissionTypeEnum{
	"MARKET_BASED":   RequestUsageCarbonEmissionsDetailsEmissionTypeMarketBased,
	"LOCATION_BASED": RequestUsageCarbonEmissionsDetailsEmissionTypeLocationBased,
}

var mappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnumLowerCase = map[string]RequestUsageCarbonEmissionsDetailsEmissionTypeEnum{
	"market_based":   RequestUsageCarbonEmissionsDetailsEmissionTypeMarketBased,
	"location_based": RequestUsageCarbonEmissionsDetailsEmissionTypeLocationBased,
}

// GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumValues Enumerates the set of values for RequestUsageCarbonEmissionsDetailsEmissionTypeEnum
func GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumValues() []RequestUsageCarbonEmissionsDetailsEmissionTypeEnum {
	values := make([]RequestUsageCarbonEmissionsDetailsEmissionTypeEnum, 0)
	for _, v := range mappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumStringValues Enumerates the set of values in String for RequestUsageCarbonEmissionsDetailsEmissionTypeEnum
func GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumStringValues() []string {
	return []string{
		"MARKET_BASED",
		"LOCATION_BASED",
	}
}

// GetMappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum(val string) (RequestUsageCarbonEmissionsDetailsEmissionTypeEnum, bool) {
	enum, ok := mappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestUsageCarbonEmissionsDetailsGranularityEnum Enum with underlying type: string
type RequestUsageCarbonEmissionsDetailsGranularityEnum string

// Set of constants representing the allowable values for RequestUsageCarbonEmissionsDetailsGranularityEnum
const (
	RequestUsageCarbonEmissionsDetailsGranularityDaily   RequestUsageCarbonEmissionsDetailsGranularityEnum = "DAILY"
	RequestUsageCarbonEmissionsDetailsGranularityMonthly RequestUsageCarbonEmissionsDetailsGranularityEnum = "MONTHLY"
)

var mappingRequestUsageCarbonEmissionsDetailsGranularityEnum = map[string]RequestUsageCarbonEmissionsDetailsGranularityEnum{
	"DAILY":   RequestUsageCarbonEmissionsDetailsGranularityDaily,
	"MONTHLY": RequestUsageCarbonEmissionsDetailsGranularityMonthly,
}

var mappingRequestUsageCarbonEmissionsDetailsGranularityEnumLowerCase = map[string]RequestUsageCarbonEmissionsDetailsGranularityEnum{
	"daily":   RequestUsageCarbonEmissionsDetailsGranularityDaily,
	"monthly": RequestUsageCarbonEmissionsDetailsGranularityMonthly,
}

// GetRequestUsageCarbonEmissionsDetailsGranularityEnumValues Enumerates the set of values for RequestUsageCarbonEmissionsDetailsGranularityEnum
func GetRequestUsageCarbonEmissionsDetailsGranularityEnumValues() []RequestUsageCarbonEmissionsDetailsGranularityEnum {
	values := make([]RequestUsageCarbonEmissionsDetailsGranularityEnum, 0)
	for _, v := range mappingRequestUsageCarbonEmissionsDetailsGranularityEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestUsageCarbonEmissionsDetailsGranularityEnumStringValues Enumerates the set of values in String for RequestUsageCarbonEmissionsDetailsGranularityEnum
func GetRequestUsageCarbonEmissionsDetailsGranularityEnumStringValues() []string {
	return []string{
		"DAILY",
		"MONTHLY",
	}
}

// GetMappingRequestUsageCarbonEmissionsDetailsGranularityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestUsageCarbonEmissionsDetailsGranularityEnum(val string) (RequestUsageCarbonEmissionsDetailsGranularityEnum, bool) {
	enum, ok := mappingRequestUsageCarbonEmissionsDetailsGranularityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
