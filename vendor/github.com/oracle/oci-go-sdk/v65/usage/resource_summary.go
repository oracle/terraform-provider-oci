// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceSummary The details of a resource under a service.
type ResourceSummary struct {

	// Units to be used for daily aggregated data.
	DailyUnitDisplayName *string `mandatory:"false" json:"dailyUnitDisplayName"`

	// Units to be used for hourly aggregated data.
	HourlyUnitDisplayName *string `mandatory:"false" json:"hourlyUnitDisplayName"`

	// Default units to use when unspecified.
	RawUnitDisplayName *string `mandatory:"false" json:"rawUnitDisplayName"`

	// Usage data type of the resource.
	UsageDataType ResourceSummaryUsageDataTypeEnum `mandatory:"false" json:"usageDataType,omitempty"`

	// Name of the resource.
	Name *string `mandatory:"false" json:"name"`

	// Name of the service.
	Servicename *string `mandatory:"false" json:"servicename"`

	// Description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// Instance type for the resource.
	InstanceType *string `mandatory:"false" json:"instanceType"`

	// Indicates if the SKU was purchased
	IsPurchased *bool `mandatory:"false" json:"isPurchased"`

	// The details of any child resources.
	ChildResources []string `mandatory:"false" json:"childResources"`

	// The details of resource Skus.
	Skus []SkuProducts `mandatory:"false" json:"skus"`
}

func (m ResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceSummaryUsageDataTypeEnum(string(m.UsageDataType)); !ok && m.UsageDataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageDataType: %s. Supported values are: %s.", m.UsageDataType, strings.Join(GetResourceSummaryUsageDataTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceSummaryUsageDataTypeEnum Enum with underlying type: string
type ResourceSummaryUsageDataTypeEnum string

// Set of constants representing the allowable values for ResourceSummaryUsageDataTypeEnum
const (
	ResourceSummaryUsageDataTypeInterval  ResourceSummaryUsageDataTypeEnum = "INTERVAL"
	ResourceSummaryUsageDataTypePointData ResourceSummaryUsageDataTypeEnum = "POINT_DATA"
)

var mappingResourceSummaryUsageDataTypeEnum = map[string]ResourceSummaryUsageDataTypeEnum{
	"INTERVAL":   ResourceSummaryUsageDataTypeInterval,
	"POINT_DATA": ResourceSummaryUsageDataTypePointData,
}

var mappingResourceSummaryUsageDataTypeEnumLowerCase = map[string]ResourceSummaryUsageDataTypeEnum{
	"interval":   ResourceSummaryUsageDataTypeInterval,
	"point_data": ResourceSummaryUsageDataTypePointData,
}

// GetResourceSummaryUsageDataTypeEnumValues Enumerates the set of values for ResourceSummaryUsageDataTypeEnum
func GetResourceSummaryUsageDataTypeEnumValues() []ResourceSummaryUsageDataTypeEnum {
	values := make([]ResourceSummaryUsageDataTypeEnum, 0)
	for _, v := range mappingResourceSummaryUsageDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceSummaryUsageDataTypeEnumStringValues Enumerates the set of values in String for ResourceSummaryUsageDataTypeEnum
func GetResourceSummaryUsageDataTypeEnumStringValues() []string {
	return []string{
		"INTERVAL",
		"POINT_DATA",
	}
}

// GetMappingResourceSummaryUsageDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceSummaryUsageDataTypeEnum(val string) (ResourceSummaryUsageDataTypeEnum, bool) {
	enum, ok := mappingResourceSummaryUsageDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
